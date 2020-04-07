package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	//line "github.com/chzyer/readline"
)

var (
	vconfig      ConfigFile
	path         string
	clientConfig struct {
		URL    string      `json:"subscribe"`
		CurIdx int         `json:"idx"`
		Nodes  []*VmessData `json:"nodes"`
	}
	process *os.Process
)

func init() {
	var (
		err error
	)
	// 获取可执行文件相对于当前工作目录的相对路径,获取可执行文件的绝对路径
	path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	// 获取项目的路径，解析config
	data, err := ioutil.ReadFile(path + "/v2ray/config.json")
	if err != nil {
		log.Println("v2ray config is not exist!")
		panic(err)
	}
	if err = json.Unmarshal(data, &vconfig); err != nil {
		log.Println("v2ray config format error")
		panic(err)
	}

	// 获取订阅URL等等
	if data, err = ioutil.ReadFile(path + "/cli-config.json"); err != nil {
		return
	}
	if err = json.Unmarshal(data, &clientConfig); err != nil {
		clientConfig.CurIdx = 0
		clientConfig.Nodes = nil
		clientConfig.URL = "nil"
	}
}

func updateClientConfig() {
	data, err := JSONMarshal(&clientConfig)
	if err != nil {
		log.Println("json format error(%v)")
	}
	err = ioutil.WriteFile(path+"/cli-config.json", data, 0644)
	if err != nil {
		log.Println("update client config failed:%v")
	}
}

func updateV2RayConfig() {
	data, err := JSONMarshal(&vconfig)
	if err != nil {
		log.Println("json format error(%v)")
	}
	err = ioutil.WriteFile(path+"/v2ray/config.json", data, 0644)
	if err != nil {
		log.Println("update v2ray config failed:%v")
	}
}

func show() {
	// ID Name Address Port
	fmt.Printf("current index: %d\n", clientConfig.CurIdx)
	fmt.Printf("%-5s%-40s%-24s%-5s\n", "ID", "Name", "Address", "Port")
	for i, vd := range clientConfig.Nodes {
		fmt.Printf("%-5d%-40s%-24s%-5d\n", i, vd.Ps, vd.Address, vd.Port)
	}
}

func selectServer(idx int) {
	servers := clientConfig.Nodes
	if len(clientConfig.Nodes) <= idx || idx < 0 {
		return
	}
	vconfig.Outbounds[0].Settings.Vnext[0].Port = servers[idx].Port
	vconfig.Outbounds[0].Settings.Vnext[0].Address = servers[idx].Address
	vconfig.Outbounds[0].Settings.Vnext[0].Users[0].ID = servers[idx].ID
	vconfig.Outbounds[0].Settings.Vnext[0].Users[0].AlterID = servers[idx].AlterID
	clientConfig.CurIdx = idx
	updateClientConfig()
	updateV2RayConfig()
	reloadV2Ray()
}

func setSubscribeURL(URL string) {
	clientConfig.URL = URL
	updateClientConfig()
}

func updateServers() {
	var (
		tempServers      []*VmessData
		URL = clientConfig.URL
	)
	// FIXME: 判断URL是否合法
	if URL == "" {
		fmt.Println("plz set subscribe URL first")
		return
	}
	clent := http.DefaultClient
	req, err := http.NewRequest("GET", URL, nil)
	//req.Header.Set() anything u would like to change
	if err != nil {
		log.Println("update subscribe servers error: %v", err)
		return
	}

	reply, err := clent.Do(req)
	if err != nil {
		log.Println("update subscribe servers error: %v", err)
		return
	}
	defer reply.Body.Close()
	body, _ := ioutil.ReadAll(reply.Body)
	step1, err := base64.StdEncoding.DecodeString(string(body))
	step2 := strings.Split(string(step1), "vmess://")

	for i, temp := range step2 {
		if i == 0 {
			continue
		}
		tempVD := new(VmessData)
		step3, _ := base64.StdEncoding.DecodeString(temp)
		if err := json.Unmarshal(step3, tempVD); err != nil {
			log.Println("Decode error %v", err)
			continue
		}
		tempServers = append(tempServers, tempVD)
	}

	if len(tempServers) > 0 {
		clientConfig.Nodes = tempServers
		log.Println("subscribe update success!")
	}
	selectServer(0)
}

func startV2Ray() {
	var (
		err error
	)
	if process != nil {
		log.Println("V2Ray is already running")
		return
	}
	attr := &os.ProcAttr{
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	process, err = os.StartProcess(path+"/v2ray/v2ray", []string{"v2ray", "-config", path + "/v2ray/config.json"}, attr)
	if err != nil {
		panic(err)
	}
}

func exitV2Ray() {
	if process != nil {
		_ = process.Kill()
		if stat, _ := process.Wait(); stat != nil {
			process = nil
		}
		log.Println("v2Ray cannot exit, plz kill it by hand")

	} else {
		log.Println("V2Ray is not running")
	}
}

func reloadV2Ray() {
	exitV2Ray()
	startV2Ray()
}

func dispose() {
	if process != nil {
		exitV2Ray()
	}

}

func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}