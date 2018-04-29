package main

const vkURL = "https://api.vk.com/method/"
const listenON = "0.0.0.0:8888"
const languageVK = "ru"
const apiVersion = "5.74"

func main() {
	go startHTTPServer()
	select {}
}
