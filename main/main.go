package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"
	"strconv"
	"sync"
)

var ECEx = STBparam

func generateRandomHexString() (string, error) {
	// Определить случайный размер строки
	size, err := rand.Int(rand.Reader, big.NewInt(32))
	for size.Cmp(big.NewInt(0)) == 0 {
		size, err = rand.Int(rand.Reader, big.NewInt(32))
	}
	if err != nil {
		return "", err
	}

	// Создать слайс байтов для хранения случайных значений
	bytes := make([]byte, size.Int64())

	// Сгенерировать случайные байты
	_, err = rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Преобразовать байты в 16-ричную строку
	return hex.EncodeToString(bytes), nil
}

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	Q := point{
		x: rev("BD1A5650179D79E03FCEE49D4C2BD5DDF54CE46D0CF11E4FF87BF7A890857FD0", 64),
		y: rev("7AC6A60361E8C8173491686D461B2826190C2EDA5909054A9AB84D2AB9D99A90", 64),
	}

	for job := range jobs {
		X, _ := generateRandomHexString()
		eds := eds_gen(ECEx.p, ECEx.a, ECEx.q, ECEx.yG, X, "1F66B5B84B7339674533F0329C74F21834281FED0732429E0C79235FC273E269")
		if !eds_check(eds, ECEx.yG, Q, X, ECEx.q, ECEx.a, ECEx.p) {
			results <- "TEST #" + strconv.Itoa(job) + " ERROR   " + X + "       " + eds
			return
		} else {
			results <- "TEST #" + strconv.Itoa(job) + " OK      " + X + "       " + eds
		}
	}
}

func main() {
	const numTests = 10000
	const numWorkers = 50

	var wg sync.WaitGroup
	jobs := make(chan int, numTests)
	results := make(chan string, numTests)

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numTests; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		log.Println(res)
	}
}
