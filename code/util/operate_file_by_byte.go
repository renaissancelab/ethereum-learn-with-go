package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "D:/gopath/src/golang_development_notes/example/log.txt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)

	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		fmt.Println(string(buf[:n]))
		break
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}


func main2() {

	file := "D:/gopath/src/golang_development_notes/example/log.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}


