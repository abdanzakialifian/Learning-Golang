package main

import (
	"learning-dependency-injection-golang/helper"
	"learning-dependency-injection-golang/inject"
)

func main() {
	server := inject.InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
