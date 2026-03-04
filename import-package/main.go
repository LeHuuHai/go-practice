package main

import (
	"fmt"
	"test/utils"

	"github.com/google/uuid"
)

func main() {
	fmt.Println(utils.Pow(2, 6))

	id := uuid.New()
	fmt.Println(id)
}
