package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	i := 0
	for {
		<-time.After(500 * time.Millisecond)
		i++
		fmt.Println("Hit", i)
		_, err := http.Get("http://localhost:50000/panel")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
