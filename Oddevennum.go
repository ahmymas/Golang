package main

import "fmt"

func main(){

fmt.Print("Input Number:")

var a int

fmt.Scanln(&a)

if (a%2==0){

fmt.Println("Even number")

}else{

fmt.Println("Odd number")

}
}
