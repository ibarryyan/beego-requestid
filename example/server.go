package main

import (
	"log"

	"github.com/beego/beego"
	"github.com/beego/beego/context"
	beego_requestid "github.com/ibarryyan/beego-requestid"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, beego_requestid.NewFilter())

	beego.Get("/hello", func(c *context.Context) {
		reqId := c.Request.Header.Get("X-Request-Id")
		log.Printf("reqestid = %s", reqId)

		_, _ = c.ResponseWriter.Write([]byte("hello..."))
		return
	})

	beego.Run(":9900")
}
