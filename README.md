## beego-requestid

### Example

```go
package main

import (
	"log"
	"time"

	"github.com/spf13/cast"
	"github.com/beego/beego"
	"github.com/beego/beego/context"
	beego_requestid "github.com/ibarryyan/beego-requestid"
)

func main() {
	example2()
}

func example1() {
	beego.InsertFilter("/*", beego.BeforeRouter, beego_requestid.NewFilter())

	beego.Get("/hello", func(c *context.Context) {
		reqId := c.Request.Header.Get("X-Request-Id")
		log.Printf("reqestid = %s", reqId)

		_, _ = c.ResponseWriter.Write([]byte("hello..."))
		return
	})

	beego.Run(":9900")
}

func example2() {
	beego.InsertFilter("/*", beego.BeforeRouter, beego_requestid.NewFilter(
		beego_requestid.WithGenRequestIdFunc(func() string {
			return cast.ToString(time.Now().Unix())
		}),
		beego_requestid.WithHeaderReqIdKey("my_header_reqid"),
		beego_requestid.WithCustomReqIdKey("my_reqid"),
	))

	beego.Get("/hello", func(c *context.Context) {
		reqId := c.Request.Header.Get("my_header_reqid")
		log.Printf("reqestid = %s", reqId)

		cReqId := c.Input.GetData("my_reqid")
		log.Printf("my reqestid = %s", cReqId)

		_, _ = c.ResponseWriter.Write([]byte("hello..."))
		return
	})

	beego.Run(":9900")
}

```


### Other   