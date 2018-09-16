package main

import "fmt"

type Circle struct {
	radius float64
}

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main()  {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
	
	var Book1 Books
	
	Book1.title = "Golang"
	Book1.author = "Jax"
	Book1.subject = "Golang"
	Book1.book_id = 4568456

	printBook(Book1)
	printBookEx(&Book1)
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func printBook(book Books)  {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

func printBookEx(book *Books)  {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}