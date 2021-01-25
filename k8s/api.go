package k8s

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var isInCluster = flag.Bool("incluster", true, "running in k8s cluster")

func Init(r *gin.Engine) {
	if *isInCluster {
		initInCfg()
	} else {
		initOutCfg()
	}
	regRt(r)
}

func regRt(r *gin.Engine) {
	r.GET("/api/v1/cfgs", func(c *gin.Context) {
		ret := lstCfgs([]string{""})
		if ret != nil {
			c.JSON(200, ret)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
	r.GET("/api/v1/cfgs/:ns", func(c *gin.Context) {
		ns := c.Param("ns")
		ret := lstCfgs([]string{ns})
		if ret != nil {
			c.JSON(200, ret)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	})
	r.GET("/api/v1/cfgs/:ns/:name", func(c *gin.Context) {
		ns := c.Param("ns")
		name := c.Param("name")
		ret := getCfg(ns, name)
		if ret != nil {
			c.JSON(200, ret)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
	r.POST("/api/v1/cfgs/:ns/:name", func(c *gin.Context) {
		ns := c.Param("ns")
		name := c.Param("name")
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		log.Printf("Update:%s\n", string(body))
		b := map[string]string{}
		err = json.Unmarshal(body, &b)
		if err != nil {
			log.Println("error format")
			c.AbortWithStatus(http.StatusBadRequest)
		}
		err = updateKvs(ns, name, b)
		if err != nil {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})
}
