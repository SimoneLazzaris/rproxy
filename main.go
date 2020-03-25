package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "fmt"
    "flag"
    "os"
)
 
type param struct {
	reverse_url	*url.URL
	listen_port	int
	tls_cert	string
	tls_key		string
	pidfile		string
}

var cfg param

func savePID() {

	file, err := os.Create(cfg.pidfile)
	if err != nil {
		panic(fmt.Sprintf("Unable to create pid file : %v\n", err))
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d",os.Getpid()))
	if err != nil {
		panic(fmt.Sprintf("Unable to create pid file : %v\n", err))
	}
	file.Sync() // flush to disk
}

func init() {
	var rUrl string
	var err  error
	flag.StringVar(&rUrl, "reverse_url", "http://127.0.0.1/","Reverse URL")
	flag.IntVar(&cfg.listen_port,"port",8080,"Listen port")
	flag.StringVar(&cfg.tls_cert,"tls_cert","","TLS Certificate file")
	flag.StringVar(&cfg.tls_key,"tls_key","","TLS Certificate key")
	flag.StringVar(&cfg.pidfile,"pidfile","/var/run/rproxy.pid","PID file")
	flag.Parse()
	cfg.reverse_url,err=url.Parse(rUrl)
	if err!=nil {
		panic(err)
	}
}
func main() { 
	// New functionality written in Go
	http.HandleFunc("/rproxy-info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "rproxy v1.0")
		})
	
	http.Handle("/", httputil.NewSingleHostReverseProxy(cfg.reverse_url))
	savePID()
	// Start the server
	var err error
	if cfg.tls_cert=="" {
		err=http.ListenAndServe(fmt.Sprintf(":%d",cfg.listen_port), nil)
	} else {
		err=http.ListenAndServeTLS(fmt.Sprintf(":%d",cfg.listen_port),cfg.tls_cert,cfg.tls_key, nil)
	}
	if err!=nil {
		panic(err)
	}
	
}
