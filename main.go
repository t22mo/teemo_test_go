package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	pwd, _ := os.Getwd()

	fmt.Println("Current time : ", now)
	fmt.Println("Current directory : ", pwd)
}
