package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("== fmt and import ==")

	serviceName := "user-api"
	port := 8081
	message := fmt.Sprintf("service=%s port=%d", serviceName, port)
	fmt.Println(message)

	fmt.Println("\n== nil values ==")

	var headers map[string]string
	fmt.Println("headers is nil:", headers == nil)
	fmt.Println("read nil map:", headers["X-Request-ID"])

	headers = make(map[string]string)
	headers["X-Request-ID"] = "req-123"
	fmt.Println("read initialized map:", headers["X-Request-ID"])

	fmt.Println("\n== short variable declaration ==")

	timeoutSeconds := 5
	timeoutSeconds = 10
	fmt.Println("timeout:", timeoutSeconds)

	userID, err := findUserID("bong")
	if err != nil {
		fmt.Println("find user error:", err)
		return
	}
	fmt.Println("user id:", userID)

	fmt.Println("\n== numeric types ==")

	var orderID int64 = 10000000001
	var retryCount int = 3
	total := orderID + int64(retryCount)
	fmt.Println("total:", total)

	var ratio float64 = 10.0 / 3.0
	fmt.Println("ratio:", ratio)
	fmt.Println("max int64:", math.MaxInt64)
}

func findUserID(username string) (int64, error) {
	if username == "" {
		return 0, fmt.Errorf("username is empty")
	}

	return 1001, nil
}
