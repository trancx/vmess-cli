package main

import (
	"github.com/chzyer/readline"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func usage(w io.Writer) {
	io.WriteString(w, "commands:\n")
	io.WriteString(w, completer.Tree("    "))
}

// Function constructor - constructs new function for listing given directory
func listFiles(path string) func(string) []string {
	return func(line string) []string {
		names := make([]string, 0)
		files, _ := ioutil.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("show"),
	readline.PcItem("subs"),
	readline.PcItem("start"),
	readline.PcItem("stop"),
	readline.PcItem("update"),
	readline.PcItem("mode"),
	readline.PcItem("bye"),
	readline.PcItem("?"),
	readline.PcItem("help"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/vmess-readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}

	defer func() {
		l.Close()
		dispose()
	}()
	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case line == "help" || line == "?":
			usage(l.Stderr())
		case line == "show":
			show()
		case strings.HasPrefix(line, "select "):
			idx, err := strconv.ParseInt(line[7:], 10, 0)
			if err == nil {
				selectServer(int(idx))
			} else {
				log.Println("index invalid!")
			}
		case strings.HasPrefix(line, "subs "):
			setSubscribeURL(line[5:])
		case line == "update":
			updateServers()
		case line == "start":
			startV2Ray()
		case line == "stop":
			exitV2Ray()
		case line == "bye":
			goto exit
		case line == "":
		default:
			log.Println("unknown command:", strconv.Quote(line))
		}
	}
exit:
}

