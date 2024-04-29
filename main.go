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

	startTime := time.Now()

	doRequest := func() {
		resp, err := http.Get(*url)
		if err == nil {
			mu.Lock()
			statusCodes[resp.StatusCode]++
			if resp.StatusCode == 200 {
				successRequests++
			}
			mu.Unlock()
			resp.Body.Close()
		} else {
			mu.Lock()
			statusCodes[0]++
			mu.Unlock()
		}
	}

	var wg sync.WaitGroup
	wg.Add(*concurrency)

	for i := 0; i < *concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < *requests / *concurrency; j++ {
				doRequest()
			}
		}()
	}

	for i := 0; i < *requests%*concurrency; i++ {
		doRequest()
	}

	wg.Wait()

	totalDuration = time.Since(startTime)

	fmt.Printf("Tempo total gasto: %v\n", totalDuration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", *requests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")

	for code, count := range statusCodes {
		fmt.Printf("%d: %d\n", code, count)
	}
}
