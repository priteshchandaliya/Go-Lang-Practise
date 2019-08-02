package main

import (
	"fmt"

	"github.com/golang-practise/helloworld"
	"github.com/golang-practise/isprime"
)

func main() {

	fmt.Println(helloworld.HelloWorld())

	fmt.Println("123 is a prime number:", isprime.IsPrime(123))
	fmt.Println("109 is a prime number:", isprime.IsPrime(109))
	fmt.Println("1 is a prime number:", isprime.IsPrime(1))
	fmt.Println("-10 is a prime number:", isprime.IsPrime(-10))

}
