package main

import "github.com/morus12/dht22"
import "fmt"
import "os"
import "time"
import "log"

// TODO CLI flags
// check data type
// grab humidity as well
// display in F and C
// timestamp it
// Handle errors
// Daemonize or make a loop

func CtoF(t float32) float32 {
	return (t*9/5 + 32)
}

func check_temp() {
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
			fmt.Fprintln(os.Stderr, "err is not nil")
			log.Fatal(err)
		}
		c1 <- temperature
	}()

	select {
	case res := <-c1:
		fmt.Print(res)
		fmt.Println(" C")
		resf := CtoF(res)
		fmt.Print(resf)
		fmt.Println(" F")
	case <-time.After(2 * time.Second):
		fmt.Println("Operation timed out")
	}
}

func main() {
	check_temp()
}
