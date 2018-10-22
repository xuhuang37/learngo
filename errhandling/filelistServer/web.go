package main

import (
	"net/http"
	"os"
	"learngo/errhandling/filelist"
			"log"
	_"net/http/pprof"
	"fmt"
)

type userError interface {
	error
	Message() string
}


type appHandler func(write http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r!=nil{
				fmt.Println(r)
				log.Printf("Panic:%v",r)
				http.Error(writer, http.StatusText(
					http.StatusInternalServerError),
					http.StatusInternalServerError)
			}

		}()
		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occurred handling request : %s",err.Error())
			if userError,ok := err.(userError);ok{
				http.Error(writer,userError.Message(),http.StatusBadRequest)
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
func main() {
	http.HandleFunc(
		"/list/",
		errWrapper(filelist.FileHandler))
	http.ListenAndServe(":8888", nil)
}
