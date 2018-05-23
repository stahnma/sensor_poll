package main

import (
	"fmt"
	"github.com/morus12/dht22"
	"log"
	"os"
	"time"
)

// grab humidity as well

func CtoF(t float32) float32 {
	return (t*9/5 + 32)
}

func check_temp() {
	GPIO_PIN := ""
	if value, ok := os.LookupEnv("GPIO_PIN"); ok {
		GPIO_PIN = "GPIO_" + value
	} else {
		GPIO_PIN = "GPIO_4"
	}

	sensor := dht22.New(GPIO_PIN)
	c1 := make(chan float32, 1)
	go func() {
		temperature, err := sensor.Temperature()
		if err != nil {
			log.SetOutput(os.Stderr)
			log.Println(err)
			return
		}
		c1 <- temperature
	}()

	log.SetOutput(os.Stdout)
	select {
	case res := <-c1:
		ctempstr := fmt.Sprintf("%.1f", res)
		log.Print(ctempstr + " C")
		resf := CtoF(res)
		ftempstr := fmt.Sprintf("%.1f", resf)
		log.Print(ftempstr + " F")
	case <-time.After(2 * time.Second):
		log.SetOutput(os.Stderr)
		log.Println("error - operation timed out")

	}
}

func main() {
	for true {
		check_temp()
		time.Sleep(5 * time.Second)
	}
}
