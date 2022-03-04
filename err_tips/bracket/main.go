package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data = []byte(`{"status": 200}`)
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T\n", result["status"]) // float64
	var status = result["status"].(int)  // 类型断言错误
	fmt.Println("Status value: ", status)
}
