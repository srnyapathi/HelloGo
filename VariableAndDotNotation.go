package main

import "fmt"

func main(){
	var whatToSay string
	whatToSay = "Hello World"

	sayHelloWorld(whatToSay)

}

func sayHelloWorld(whatToSay string){
	fmt.Println(whatToSay)
}
