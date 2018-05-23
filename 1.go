package main

import "github.com/morus12/dht22"
import "fmt"
import "os"
import "time"
import "log"

func main() {
	PIN := ""
	if value, ok := os.LookupEnv("PIN"); ok {
		PIN = "GPIO_" + value
	} else {
		PIN = "GPIO_4"
	}

	sensor := dht22.New(PIN)
	c1 := make(chan float32, 1)
	go func() {
		temperature, err := sensor.Temperature()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(temperature)
		c1 <- temperature
	}()

	select {
	case res := <-c1:
		fmt.Println(c1)
		fmt.Println(res)
	case <-time.After(5 * time.Second):
		fmt.Println("Operation timed out")
	}

}
