package main

import "mvc_for_gin/libs/server"

var http = new(server.HttpServer)

func main() {
	http.Default().Run()
}
