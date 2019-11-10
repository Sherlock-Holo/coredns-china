package main

import (
	"flag"
	"log"

	china_list "github.com/Sherlock-Holo/coredns-china/pkg/china-list"
)

func main() {
	output := flag.String("output", "china-list.txt", "output path")
	flag.Parse()

	if err := china_list.Run(*output); err != nil {
		log.Fatalf("generate china-list failed: %+v", err)
	}
}
