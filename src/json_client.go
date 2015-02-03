package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(res.Body)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}

		//まずinterfaceでうける
		//json化
		var result = interface {}
		fmt.Println(result)
		//		if err = json.Unmarshal(scanner.Bytes(), &result); err != nil {
		//			fmt.Println(err)
		//			continue
		//		}

	}
}
