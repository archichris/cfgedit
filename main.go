package main

import (
	"cfgedit/k8s"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	port    = flag.String("port", "8080", "port of server ")
	tls     = flag.Bool("tls", true, "Whether to enable TLS")
	keyPath = flag.String("key", "./ca/ca.key", "Path of key")
	crtPath = flag.String("crt", "./ca/ca.crt", "Path of crt")
	base    = "api/v1"
)

func main() {
	err := flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		log.Panic(err)
	}
	r := gin.Default()
	g := r.Group(base)

	r.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	k8s.Init(g)
	if *tls {
		_ = r.RunTLS(":"+*port, *crtPath, *keyPath)
	} else {
		_ = r.Run(":" + *port)
	}
}

// r := gin.Default()
// server := http.Server{
// 	Addr:      addr,
// 	Handler:   r,
// 	TLSConfig: tlsConfig,
// }
// err = server.ListenAndServeTLS("", "")
// }
