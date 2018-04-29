package main

import (
	"log"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func startHTTPServer() {

	router := routing.New()

	router.Get("/", XMLHandler)

	server := fasthttp.Server{
		Name:    "VK2RSS By RiftBit",
		Handler: router.HandleRequest,
	}

	log.Println("Started server on ", listenON)
	log.Fatal(server.ListenAndServe(listenON))
}

func XMLHandler(ctx *routing.Context) error {

	ctx.Response.Header.SetContentType("application/xml; charset=UTF-8")

	data := getDataFromVK()

	res, err := dataToRSS(data)
	if err != nil {
		return err
	}

	ctx.Response.SetBodyString(res)

	return nil
}
