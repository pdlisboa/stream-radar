package main

import (
	"stream-radar/internal/environment"
	"stream-radar/worker"
)

func init() {
	environment.InitEnv()
}

func main() {

	//api.InitApi()
	worker.TestTwitch("esacarry")
	worker.TestKick("esa")
	worker.TestYt("esacarry")
}
