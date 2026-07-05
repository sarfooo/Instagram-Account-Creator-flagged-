package main

import (
	"math/rand"
	"bufio"
	"time"
	"fmt"
	"os"
)

const Charset string = "ABCDEFGHIJKLMNOPQRSTUVWXYZsabcdefghijklmnopqrstuvwxyz0123456789"

func openFile(filename string) (slice []string) {
	file, _ := os.Open(filename)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		slice = append(slice, scan.Text())
	}
	file.Close()
	return slice
}

func writeFile(filename string, content string) {
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fmt.Fprintln(file, content)
	file.Close()
}

func randomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = Charset[rand.New(source).Intn(len(Charset))]
	}
	return string(bytes)
}