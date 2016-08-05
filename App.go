package main

import (
	"fmt"
	"ohmyproxy-client-golang/ohmyproxyclient"
)

func main() {
	rs := ohmyproxyclient.GetProxiesFromWS()
	fmt.Println(rs)
}
