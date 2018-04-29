package main

import (
	"log"

	"github.com/valyala/fasthttp"
)

func startHTTPServer() {

	server := fasthttp.Server{
		Name:    "VK2RSS By RiftBit",
		Handler: XMLHandler,
	}

	log.Println("Started server on ", listenON)
	log.Fatal(server.ListenAndServe(listenON))
}

func XMLHandler(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.SetContentType("application/xml; charset=UTF-8")

	data := getDataFromVK()

	var res string
	var err error

	if len(data.Response.Items) > 0 {
		res, err = dataToRSS(data)
		if err != nil {
			ctx.Response.SetBodyString(err.Error())
			ctx.SetStatusCode(500)
			return
		}
	} else {
		ctx.Response.SetBodyString("something going wrong")
		ctx.SetStatusCode(500)
		return
	}

	ctx.Response.SetBodyString(res)
}
