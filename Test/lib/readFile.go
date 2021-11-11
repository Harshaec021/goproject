package lib

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(fileName string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
