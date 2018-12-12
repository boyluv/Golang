package main

import (
	"fmt"
	"math/rand"
	"time"
)

type point struct {
	x,y int
}
func main() {
	//For more info
	//https://gobyexample.com/string-formatting


	p := point{1, 2}
	fmt.Printf("%v", p)
	fmt.Printf("\n\n")
	//Result
	//{1 2}

	fmt.Printf("%+v",p)
	fmt.Printf("\n\n")
	//Result
	//{x:1 y:2}


	fmt.Printf("%#v",p)
	fmt.Printf("\n\n")
	//Result
	//main.point{x:1, y:2}

	fmt.Printf(" %T",p)
	fmt.Printf("\n\n")
	//Result
	//main.point

	fmt.Printf(" %t",true)
	fmt.Printf("\n\n")
	//Result
	//true

	fmt.Printf(" %d",124/3)
	fmt.Printf("\n\n")
	//Result
	//41

	fmt.Printf(" %b",14)
	fmt.Printf("\n\n")
	//Result: Print binary
	//1110


	fmt.Printf(" %c",33)
	fmt.Printf("\n\n")
	//Result: Print corresponding interger
	//!


	fmt.Printf(" %x",456)
	fmt.Printf("\n\n")
	//Result: Print hex
	//1c8

	fmt.Printf("%f\n", 78.9)
	fmt.Printf("\n\n")
	//Result: Print hex
	//78.900000

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("\n\n")
	//Result: Print string
	//"string"

	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("\n\n")
	//Result: Print double qute string
	//"\"string\""

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("\n\n")
	//Result: Control the width and precision of the resulting figure
	//|    12|   345|

	fmt.Println("The time is", time.Now())
	fmt.Println("time is",rand.Intn(10))

	fmt.Println(add(8,5))

	fmt.Println(swap("World","Hello"))

	var i int
	fmt.Println(i,c,python,java)

	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func add(x,y int) int  {
	return (x+y)
}
var c,python,java bool
func swap(a,b string) (string,string)  {
	return b,a

}

