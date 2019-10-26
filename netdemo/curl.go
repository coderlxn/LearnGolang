package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init()  {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curl <url>")
		os.Exit(-1)
	}
}

func main()  {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := io.Copy(os.Stdout, r.Body); err != nil {
		fmt.Println(err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
