package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
    url           string
    totalRequests int
    concurrency   int
    maxRetries    int
)

func init() {
    flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
    flag.IntVar(&totalRequests, "requests", 100, "Número total de requests")
    flag.IntVar(&concurrency, "concurrency", 10, "Número de chamadas simultâneas")
    flag.IntVar(&maxRetries, "retries", 3, "Número máximo de tentativas em caso de falha")
}

func main() {
    flag.Parse()

    if url == "" {
        fmt.Println("A URL é obrigatória")
        return
    }

    startTime := time.Now()

    var wg sync.WaitGroup
    requestCh := make(chan struct{}, concurrency)
    results := make(chan int, totalRequests)

    for i := 0; i < totalRequests; i++ {
        wg.Add(1)
        requestCh <- struct{}{}

        go func() {
            defer wg.Done()
            performRequest(results)
            <-requestCh
        }()
    }

    wg.Wait()
    close(results)

    totalDuration := time.Since(startTime)
    reportResults(totalDuration, results)
}

func performRequest(results chan<- int) {
    var resp *http.Response
    var err error
    for attempt := 0; attempt <= maxRetries; attempt++ {
        resp, err = http.Get(url)
        if err == nil {
            results <- resp.StatusCode
            resp.Body.Close()
            return
        }
        time.Sleep(500 * time.Millisecond) // espera antes de tentar novamente
    }
    fmt.Println("Erro ao fazer request após múltiplas tentativas:", err)
    results <- -1 // Use -1 para indicar um erro de requisição após retries
}

func reportResults(totalDuration time.Duration, results <-chan int) {
    totalRequests := 0
    statusCounts := make(map[int]int)

    for status := range results {
        totalRequests++
        statusCounts[status]++
    }

    fmt.Printf("Tempo total gasto: %v\n", totalDuration)
    fmt.Printf("Quantidade total de requests: %d\n", totalRequests)
    fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", statusCounts[200])
    for status, count := range statusCounts {
        if status != 200 {
            if status == -1 {
                fmt.Printf("Quantidade de requests que falharam após retries: %d\n", count)
            } else {
                fmt.Printf("Quantidade de requests com status HTTP %d: %d\n", status, count)
            }
        }
    }
}
