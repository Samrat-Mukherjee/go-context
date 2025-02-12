package main

import (
	"context"
	"math/rand"

	"fmt"
	"time"
)

func openConnection(done chan bool) {
	fmt.Println("Establishing Connection...")

	if rand.Intn(300) > 50 {
		fmt.Println("OOPS, Your connection hanged")
		time.Sleep(3 * time.Hour)

	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connection Established!")
	}

	done <- true
}

func connectionWithTimeOut() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan bool)

	go openConnection(done)

	select {
	case <-done:
		fmt.Println("Connection successful")
	case <-ctx.Done():
		fmt.Println("Connection timeout")
	}
}

func main() {

	connectionWithTimeOut()

	// req, err := http.NewRequestWithContext(ctx, "GET", "https://api.example.com/data", nil)

	// if err != nil {
	// 	fmt.Println("Error creating request", err)
	// 	return
	// }

	// client := http.DefaultClient
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error creating request", err)
	// 	return
	// }

	//Interfaces
	//Goruting
	//context

	//defer resp.Body.Close()
}
