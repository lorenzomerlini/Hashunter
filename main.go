package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func recursiveBrute(charset, current, targetHash string, maxLen int, found chan string, stop chan struct{}) {

	// verify if length is reached
	if len(current) == maxLen {
		return
	}
	// calculate hash and confront with target hash
	if HashPassword(current) == targetHash {
		found <- current
		close(stop)
		return
	}
	// generate new combinations
	for _, char := range charset {
		select {
		case <-stop:
			return //termination
		default:
			recursiveBrute(charset, current+string(char), targetHash, maxLen, found, stop)
		}
	}
}

func randomString(length int, charset string) string {

	rand.Seed(time.Now().UnixNano())
	var result string
	for i := 0; i < length; i++ {
		result += string(charset[rand.Intn(len(charset))])
	}
	return result
}

func Bruteforce(targetHash string, charset string, maxLen int, timeout time.Duration) (string, bool) {

	var wg sync.WaitGroup // waitgroup for goroutines
	found := make(chan string)
	stop := make(chan struct{})
	segmentSize := len(charset) / 4 // divide charset in 4 segments

	// start parallel searching
	brute := func(current string, startIdx, endIdx int) {

		defer wg.Done()
		// verify if password is found
		for i := startIdx; i < endIdx; i++ {
			select {
			case <-stop:
				return // termination
			default:
				recursiveBrute(charset, current+string(charset[i]), targetHash, maxLen, found, stop)
			}
		}
	}

	// loading
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("\nDone.")
				return
			default:
				randomMsg := randomString(12, charset)
				fmt.Printf("\r%s", randomMsg)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// set a timeout
	go func() {
		if timeout > 0 {
			time.Sleep(timeout)
			close(stop)
		}
	}()

	//start parallel search
	for i := 0; i < 4; i++ {
		wg.Add(1)
		startIdx := i * segmentSize
		endIdx := startIdx + segmentSize
		if i == 3 {
			endIdx = len(charset)
		}
		go brute("", startIdx, endIdx)
	}

	// wait the termination
	go func() {
		wg.Wait()
		close(found)
	}()

	// retrieve the result
	select {
	case result := <-found:
		return result, true
	case <-stop:
		return "", false
	}
}

func main() {

	var password string // target password
	fmt.Println("Insert password to be guessed: ")
	fmt.Scanln(&password)
	hash := HashPassword(password) // create target hash
	fmt.Printf("Target hash: %s\n", hash)

	var charset string // set of character to be used to bruteforce the password
	fmt.Println("Insert charset (default a-z 0-9): ")
	fmt.Scanln(&charset)
	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyz0123456789" // default characters
	}
	maxLen := 6 // password max lenght

	var timeoutSec int // timeout in seconds
	fmt.Println("Insert max timeout in seconds (0 for no timeout): ")
	fmt.Scanln(&timeoutSec)
	timeout := time.Duration(timeoutSec) * time.Second

	start := time.Now()                                                  // start counting time
	crackedPassword, found := Bruteforce(hash, charset, maxLen, timeout) // start bruteforce
	duration := time.Since(start)                                        // stop counting time

	if found {
		fmt.Printf("\r%s\n", "Password found!")
		fmt.Printf("Password: %s\n\n", crackedPassword)
	} else {
		fmt.Println("Failed to crack the password withing the given time.")
	}
	fmt.Printf("Time taken: %s\n", duration)
}
