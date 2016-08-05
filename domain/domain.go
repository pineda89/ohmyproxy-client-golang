package domain

type Response struct {
	ProxyList []Proxy `json:"proxyList"`
}

type Proxy struct {
	Country       string `json:"country"`
	DiscoveryTime int         `json:"discoveryTime"`
	IP            string      `json:"ip"`
	LastCheckTime int         `json:"lastCheckTime"`
	Port          int         `json:"port"`
	RsTime        int         `json:"rsTime"`
	Type          string      `json:"type"`
	Website       string      `json:"website"`
}

