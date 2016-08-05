# OhMyProxy-client-golang

OhMyProxy-client-golang is a library that provides free proxies

The operation of this library is by connecting via GET requests to the server , which contains different proxies, and will be updated every few minutes .

Perfect for:
  - Web scraping
  - Anonymous navigation
  - Check your services healthy from another site

### Usage

```sh
import (
	"ohmyproxy-client-golang/ohmyproxyclient"
)

func main() {
	rs := ohmyproxyclient.GetProxiesFromWS()
	fmt.Println(rs)
}
```

### Todos

 - Filter by country
 - More proxies!!!