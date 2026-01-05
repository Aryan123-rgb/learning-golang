package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	temperature := rand.Intn(15) + 20 // 20–34 °C range
	ch <- fmt.Sprintf("The temperature for %s is %d", city, temperature)
}

func main() {
	startTime := time.Now()

	cities := []string{"London", "Paris", "Tokyo", "Toronto"}
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(startTime))
}
