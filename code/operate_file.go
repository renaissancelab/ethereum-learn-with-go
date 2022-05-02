package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func IntegerToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	filepath := "./test.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	count := 0
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
        binary := IntegerToBinary(count)
		whole := fmt.Sprintf("%011s", binary)
		fmt.Println(count,"\t", whole,"\t",line)
        count++
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}

}