package main

import "fmt"

func main() {
	//var serviceName string = "payment-api"
	//var port int = 8080
	//var timeoutSeconds = 5
	//debugEnabled := false
	//
	//var retryCount int
	//var metadata map[string]string
	//
	//fmt.Println("service:", serviceName)
	//fmt.Println("port:", port)
	//fmt.Println("timeout:", timeoutSeconds)
	//fmt.Println("debug:", debugEnabled)
	//fmt.Println("retry zero value:", retryCount)
	//fmt.Println("metadata is nil:", metadata == nil)

	var userID int = 1
	message := fmt.Sprintf("user_id=%d not found", userID)
	fmt.Println(message)
}
