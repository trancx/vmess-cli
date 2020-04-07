package main

type ConfigFile struct {
	Policy interface{} `json:"policy"`
	Log    struct {
		Access   string `json:"access"`
		Error    string `json:"error"`
		Loglevel string `json:"loglevel"`
	} `json:"log"`
	Inbounds []struct {
		Tag      string `json:"tag"`
		Port     int    `json:"port"`
		Listen   string `json:"listen"`
		Protocol string `json:"protocol"`
		Sniffing struct {
			Enabled      bool     `json:"enabled"`
			DestOverride []string `json:"destOverride"`
		} `json:"sniffing"`
		Settings struct {
			Auth    string      `json:"auth"`
			UDP     bool        `json:"udp"`
			IP      interface{} `json:"ip"`
			Address interface{} `json:"address"`
			Clients interface{} `json:"clients"`
		} `json:"settings"`
		StreamSettings interface{} `json:"streamSettings"`
	} `json:"inbounds"`
	Outbounds []struct {
		Tag      string `json:"tag"`
		Protocol string `json:"protocol"`
		Settings struct {
			Vnext []struct {
				Address string `json:"address"`
				Port    int    `json:"port"`
				Users   []struct {
					ID       string `json:"id"`
					AlterID  int    `json:"alterId"`
					Email    string `json:"email"`
					Security string `json:"security"`
				} `json:"users"`
			} `json:"vnext"`
			Servers  interface{} `json:"servers"`
			Response interface{} `json:"response"`
		} `json:"settings"`
		StreamSettings struct {
			Network      string      `json:"network"`
			Security     interface{} `json:"security"`
			TLSSettings  interface{} `json:"tlsSettings"`
			TCPSettings  interface{} `json:"tcpSettings"`
			KcpSettings  interface{} `json:"kcpSettings"`
			WsSettings   interface{} `json:"wsSettings"`
			HTTPSettings interface{} `json:"httpSettings"`
			QuicSettings interface{} `json:"quicSettings"`
		} `json:"streamSettings"`
		Mux struct {
			Enabled bool `json:"enabled"`
		} `json:"mux"`
	} `json:"outbounds"`
	Stats   interface{} `json:"stats"`
	API     interface{} `json:"api"`
	DNS     interface{} `json:"dns"`
	Routing struct {
		DomainStrategy string `json:"domainStrategy"`
		Rules          []struct {
			Type        string      `json:"type"`
			Port        interface{} `json:"port"`
			InboundTag  string      `json:"inboundTag"`
			OutboundTag string      `json:"outboundTag"`
			IP          interface{} `json:"ip"`
			Domain      interface{} `json:"domain"`
		} `json:"rules"`
	} `json:"routing"`
}

type VmessData struct {
	Version int    `json:"v"`
	Ps      string `json:"ps"`
	Address string `json:"add"`
	AlterID int    `json:"aid"`
	Port    int    `json:"port"`
	ID      string `json:"id"`
	Net     string `json:"net"`
	Type    string `json:"type"`
	Host    string `json:"host"`
	Path    string `json:"path"`
	TLS     string `json:"tls"`
}
