package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	for {
		randomenough := uuid.New().String()[0:6]
		fmt.Println(randomenough)
		time.Sleep(5 * time.Second)
	}
}
