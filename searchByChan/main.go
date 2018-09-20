package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

//init 将在main前调用
func init()  {
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run("president")
}