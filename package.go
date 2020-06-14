package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unsafe"
)

func bytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func getInput() string {
	fmt.Print("Are you OK? ")
	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	return bytesToString(line)
}

func isPositive(answer string) bool {
	switch answer {
	case
		"yes", "y", "aye", "yeah", "indeed":
		return true
	}
	return false

}

func message(answer string) string {
	saneAnswer := strings.ToLower(answer)
	if isPositive(saneAnswer) {
		return "Glad to hear it mate!"
	}
	return "I'm sorry about that!"
}

func main() {
	answer := getInput()
	fmt.Println(message(answer))
}
