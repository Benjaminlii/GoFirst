package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := "/Users/benjamin/code/GoFirst/hello/a.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error: Open failed")
		return
	}

	var data []byte
	buffer := make([]byte, 1000)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			fmt.Println("error: Read failed")
		}
		if n == 0 {
			break
		}
		data = append(data, buffer[0:n]...)
	}

	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")
	intLines := make([]int, 0)
	for _, v := range lines {
		intLine, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("info: Read failed")
		} else {
			intLines = append(intLines, intLine)
		}
	}

	sort.Ints(intLines)

	fmt.Println(intLines)

}
