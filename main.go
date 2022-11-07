package main

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	_ "go-test/statik"
	"io/ioutil"
	"log"
)

func main() {
	//statikFS为http.FileSystem类型
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	// 或者在gin框架下router.StaticFS("/public", statikFS)
	//获取文件内容
	r, err := statikFS.Open("/a.json")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(contents))
}
