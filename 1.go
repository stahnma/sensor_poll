package main

import "github.com/morus12/dht22"
import "fmt"
import "os"

func main() {
	PIN := ""
	if value, ok := os.LookupEnv("PIN"); ok {
		PIN = "GPIO_" + value
	} else {
		PIN = "GPIO_17"
	}

	sensor := dht22.New(PIN)
	temperature, err := sensor.Temperature()

	fmt.Println(temperature)
	fmt.Println(err)

}
