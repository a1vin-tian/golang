package main

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

type WebServerConfig struct {
	Port int
}

type WebServer struct {
	server *http.Server
	Port   int
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func helloHand(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		return
	}
	for k, v := range r.Header {
		for _, s := range v {
			w.Header().Add(k, s)
		}
	}
	if env, b := os.LookupEnv("VERSION"); b {
		w.Header().Set("VERSION", env)
	}

}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if ip == "::1" {
			return "127.0.0.1"
		}
		return ip
	}

	return ""
}

func NewWebServer(p WebServerConfig) *WebServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHand)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/401", badrequest)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	s := WebServer{
		server: &http.Server{
			Addr: fmt.Sprintf(":%v", p.Port),
		}, Port: p.Port}
	s.server.Handler = Logger(mux)
	return &s
}

// Logs incoming requests.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		o := &responseObserver{ResponseWriter: w}
		o.status = 200
		h.ServeHTTP(o, r)
		glog.Infof("[URI:%s][ClientIP:%s][Status:%d]", r.URL.Path, ClientIP(r), o.status)
	})
}

type responseObserver struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (o *responseObserver) Write(p []byte) (n int, err error) {
	if !o.wroteHeader {
		o.WriteHeader(http.StatusOK)
	}
	n, err = o.ResponseWriter.Write(p)
	o.written += int64(n)
	return
}

func (o *responseObserver) WriteHeader(code int) {
	o.ResponseWriter.WriteHeader(code)
	if o.wroteHeader {
		return
	}
	o.wroteHeader = true
	o.status = code
}

func badrequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
}

func (s *WebServer) Run() {
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
