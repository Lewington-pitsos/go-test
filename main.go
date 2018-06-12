package main

import (
	"encoding/json"
	"fmt"
	"go/build"
	"io/ioutil"
)

func main() {
	var path = build.Default.GOPATH + "/src/test/data.json"
	fmt.Println(path)

	dat, err := ioutil.ReadFile(path)
	check(err)
	output := make([]string, 10000)
	json.Unmarshal(dat, &output)
	fmt.Print(output[8:10])

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
