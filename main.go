package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("URL é um parâmetro obrigatório")
		return
	}

	var totalDuration time.Duration
	var successRequests int
	statusCodes := make(map[int]int)
	var mu sync.Mutex

	doRequest := func(wg *sync.WaitGroup) {
		defer wg.Done()

		start := time.Now()
		resp, err := http.Get(*url)
		duration := time.Since(start)

		mu.Lock()
		defer mu.Unlock()

		totalDuration += duration

		if err == nil {
			statusCodes[resp.StatusCode]++
			if resp.StatusCode == 200 {
				successRequests++
			}
			resp.Body.Close()
		} else {
			statusCodes[0]++
		}
	}

	var wg sync.WaitGroup

	requestsPerGoroutine := *requests / *concurrency

	extraRequests := *requests % *concurrency

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			numRequests := requestsPerGoroutine
			if i < extraRequests {
				numRequests++
			}
			for j := 0; j < numRequests; j++ {
				doRequest(&wg)
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("Tempo total gasto: %v\n", totalDuration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", *requests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")

	for code, count := range statusCodes {
		fmt.Printf("%d: %d\n", code, count)
	}
}
