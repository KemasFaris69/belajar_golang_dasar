package belajar_golang_dasar

import (
	"fmt"
	"testing"
	"time"
)

func RunHello() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHello()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}
