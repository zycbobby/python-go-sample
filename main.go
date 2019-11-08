package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	python "github.com/sbinet/go-python"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

var pyfile = flag.String("pyfile", "./py-service/main.py", "main entrance to the py file")

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	flag.Parse()

	fmt.Printf("going to execute  %s\n", *pyfile)

	err := python.PyRun_SimpleFile(*pyfile)
	if err != nil {
		fmt.Printf("run py, err=%v", err)
	}
}
