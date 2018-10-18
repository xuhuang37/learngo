package main

import (
	"fmt"
	real2 "learngo/retriever/real"
	"learngo/retriever/mock"
	"time"
)

const url = "https://www.imooc.com"
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
	//Connect(host string)
}

func downloader(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name":"huangx",
		"alias":"fucker",
	})
}



func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)

}


func inspect(r Retriever)  {
	fmt.Println("Inspecting",r)
	fmt.Printf(">%T %v\n",r,r)
	fmt.Printf("> Type switch")

	switch v:=r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:",v.Contents)
	case *real2.Retriever:
		fmt.Println("UseAgent:",v.UserAgent)
	}
	fmt.Println()
}

//func ()  {
//
//}

func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake baidu"}
	r = &retriever
	fmt.Printf("%T %v\n", r, r)
	inspect(r)
	r = &real2.Retriever{"mozila 5.0", time.Minute}
	fmt.Printf("%T %v\n", r, r)
	inspect(r)

	//type assertion

	if mockRetriever,ok:= r.(*mock.Retriever);ok{
		fmt.Println(mockRetriever.Contents)
	}else {
		fmt.Println("not a mock retriever")
	}
	fmt.Printf(session(&retriever))
	retriever.String()
	//realRetriever := r.(*real2.Retriever)
	//fmt.Println(realRetriever)
	//fmt.Println(downloader(r))
}
