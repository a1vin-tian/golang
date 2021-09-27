package main

import (
	"flag"
	"github.com/golang/glog"
	_ "net/http/pprof"
)

func main() {
	glog.Info("Starting http server...")
	port := flag.Int("port", 8080, "Server port")
	server := NewWebServer(WebServerConfig{
		Port: *port,
	})
	defer server.server.Close()
	server.Run()
}
