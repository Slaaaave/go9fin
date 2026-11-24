package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = rand.Int()
	}
	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	maxData := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > maxData {
			maxData = data[i]
		}
	}
	return maxData
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}
	chunksSize := len(data) / CHUNKS
	maxValue := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		start := i * chunksSize
		end := start + chunksSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)

		go func(slice []int, index int) {
			defer wg.Done()
			maxValue[index] = maximum(slice)
		}(data[start:end], i)
	}
	wg.Wait()
	return maximum(maxValue)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел: ", SIZE)
	date := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startSingle := time.Now()
	maxSingle := maximum(date)
	elapsed := time.Since(startSingle).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxSingle, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	startMulti := time.Now()
	maxMulti := maxChunks(date)
	elapsed = time.Since(startMulti).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", maxMulti, elapsed)
}
