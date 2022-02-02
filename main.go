package main

import "mvc_for_gin/libs"

var Server = libs.Server

func main() {
	Server.Default().NoDataBase().Run()
}
