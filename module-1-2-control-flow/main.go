package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("== if validation ==")
	if err := validateCreateOrderRequest("user-1001", 25000); err != nil {
		fmt.Println("invalid request:", err)
	} else {
		fmt.Println("request is valid")
	}

	fmt.Println("\n== for retry loop ==")
	if err := callPaymentGatewayWithRetry(3); err != nil {
		fmt.Println("payment failed:", err)
	} else {
		fmt.Println("payment success")
	}

	fmt.Println("\n== switch routing decision ==")
	fmt.Println("handler:", routeHandler("POST", "/orders"))
	fmt.Println("handler:", routeHandler("GET", "/health"))
	fmt.Println("handler:", routeHandler("DELETE", "/orders/1"))

	fmt.Println("\n== defer cleanup ==")
	processFileLikeResource("orders.csv")
}

func validateCreateOrderRequest(userID string, amount int64) error {
	if strings.TrimSpace(userID) == "" {
		return fmt.Errorf("user id is required")
	}

	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	return nil
}

func callPaymentGatewayWithRetry(maxAttempts int) error {
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Println("attempt:", attempt)

		if attempt == 2 {
			return nil
		}

		time.Sleep(100 * time.Millisecond)
	}

	return fmt.Errorf("all attempts failed")
}

func routeHandler(method string, path string) string {
	switch {
	case method == "POST" && path == "/orders":
		return "createOrder"
	case method == "GET" && path == "/health":
		return "healthCheck"
	default:
		return "notFound"
	}
}

func processFileLikeResource(name string) {
	fmt.Println("open:", name)
	defer fmt.Println("close:", name)

	fmt.Println("read:", name)
	fmt.Println("process:", name)
}

