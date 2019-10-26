package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	{
		//写入文件
		fo, err := os.Create("file.txt")
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
		for i := 0; i < 1024; i++ {
			value := fmt.Sprintf("index: %d, value: %d\n", i, i*2)
			if _, err := fo.Write([]byte(value)); err != nil {
				panic(err)
			}
		}
	}

	{
		//从文件中读取内容
		fi, err := os.Open("file.txt")
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := fi.Close(); err != nil {
				panic(err)
			}
		}()
		buf := make([]byte, 1024)
		for {
			n, err := fi.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			fmt.Println(string(buf[:n]))
		}
	}
}
