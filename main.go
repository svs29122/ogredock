package main

import(
	"fmt"
	//"ogredock/contmod"
	"ogredock/webmod"
)

func main(){
	fmt.Println("--OgreDock v.1.0--")

	//grab containers
	//contmod.ListContainers()

	//start webserver
	webmod.StartServer()
}
