package proxy

import (
	"net/http/httputil"
	"net/url"
	"net/http"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type (
	Proxy struct {
		ListenHost  string
		ListenPort  int
		DefaultHost string
		Hosts       []Host
		Proxy       map[string]*httputil.ReverseProxy
	}
	Host struct {
		Host   string
		Target string
	}
)

var proxyCfg Proxy

func Run(fileConfig string) {
	
	proxyCfg.Proxy = make(map[string]*httputil.ReverseProxy)
	if _, err := toml.DecodeFile(fileConfig, &proxyCfg); err != nil {
		fmt.Println(err)
	}
	
	for _, v := range proxyCfg.Hosts {
		target, _ := url.Parse(v.Target)
		proxyCfg.Proxy[v.Host] = httputil.NewSingleHostReverseProxy(target)
	}
	
	http.HandleFunc("/", handleProxy)
	
	listenAddr := fmt.Sprintf("%s:%d", proxyCfg.ListenHost, proxyCfg.ListenPort)
	log.Printf("Start reverse proxy on %s...\n", listenAddr)
	log.Println(http.ListenAndServe(listenAddr, nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	
	dr, _ := httputil.DumpRequest(r, false)
	if env := os.Getenv("ENV"); env == "DEV" {
		log.Println(string(dr))
	}
	
	if rp, ok := proxyCfg.Proxy[r.Host]; ok == true {
		rp.ServeHTTP(w, r)
	} else {
		log.Printf("[WARNING] Default Host %s\n", proxyCfg.DefaultHost)
		proxyCfg.Proxy[proxyCfg.DefaultHost].ServeHTTP(w, r)
	}
	
}
