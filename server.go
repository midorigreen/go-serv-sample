package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func server(test chan string) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	fmt.Println("Listen")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 1 request to 1 goroutine
		go func(chan string) {
			defer conn.Close()

			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			req, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				log.Println(err)
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(dump))
			test <- "test string"
		}(test)
	}
}
