package main

import (
	"log"
	"sync"
)

func main() {
	// mu 代表以下流程要使用的互斥锁。
	var mu sync.Mutex

	mu.Lock()
	_, err := writer.Write([]byte(data))
	if err != nil {
		log.Printf("error: %s [%d]", err, id)
	}
	mu.Unlock()

}
