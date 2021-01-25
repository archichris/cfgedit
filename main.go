package main

import (
	"cfgedit/k8s"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		log.Panic(err)
	}
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"healtyz": "OK"})
	})
	k8s.Init(r)
	_ = r.Run()
}
