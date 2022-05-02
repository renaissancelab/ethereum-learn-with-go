package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

//
import (
"os"
"io/ioutil"
"fmt"
)

func main() {
	file, err := os.Open("D:/gopath/src/golang_development_notes/example/log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

func main1() {
	filepath := "D:/gopath/src/golang_development_notes/example/log.txt"
	content ,err :=ioutil.ReadFile(filepath)
	if err !=nil {
		panic(err)
	}
	fmt.Println(string(content))
}
