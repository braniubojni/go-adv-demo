package main

import (
	"fmt"
	"sync"
)

func fakeMain() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	routineNum := 4
	var wg sync.WaitGroup
	channel := make(chan int)
	for i := 0; i < len(arr); i += routineNum {
		prefix := i
		suffix := i + routineNum
		slice := arr[prefix:suffix]
		wg.Add(1)
		go func() {
			sum(slice, channel)

			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(channel)
	}()
	var result int
	for res := range channel {
		result += res
	}
	fmt.Println(result, "result")
}

func sum(nums []int, channel chan int) {
	var result int
	for _, num := range nums {
		result += num
	}
	channel <- result
}

// func main() {
// 	code := make(chan int)
// 	var wg sync.WaitGroup
// 	for range 10 {
// 		wg.Add(1)
// 		go func() {
// 			checkStatus(code)
// 			wg.Done()
// 		}()
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(code)
// 	}()
// 	for res := range code {
// 		fmt.Printf("code: %d\n", res)
// 	}
// }

// func checkStatus(codeCh chan int) {
// 	response, err := http.Get("http://google.com")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	if response.StatusCode != 200 {
// 		fmt.Printf("Status code is: %d\n", response.StatusCode)
// 	}
// 	defer func() {
// 		response.Body.Close()
// 		fmt.Printf("Status code from child: %d\n", response.StatusCode)
// 	}()
// 	codeCh <- response.StatusCode
// }

// func main() {
// 	t := time.Now()
// 	var wg sync.WaitGroup
// 	for range 10 {
// 		wg.Add(1)
// 		go checkStatus(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Printf("Time taken: %v\n", time.Since(t))
// }

// func checkStatus(wg *sync.WaitGroup) {
// 	response, err := http.Get("http://google.com")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	if response.StatusCode != 200 {
// 		fmt.Printf("Status code is: %d\n", response.StatusCode)
// 	}
// 	defer func() {
// 		response.Body.Close()
// 		fmt.Printf("Status code is: %d\n", response.StatusCode)
// 		wg.Done()
// 	}()
// }
