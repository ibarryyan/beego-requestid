## beego-requestid

### Example

```go

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

```


### Other   