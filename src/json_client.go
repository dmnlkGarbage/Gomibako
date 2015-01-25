package main

import (
	"net/http"
	"bufio"
	"fmt"
)

func main() {
	resp, _ := http.Get("http://localhost:8000")
	//defer resp.Body.Close()
	scan := bufio.NewScanner(resp.Body)
	for scan.Scan() {
		fmt.Println(string(scan.Bytes()))
	}
}

