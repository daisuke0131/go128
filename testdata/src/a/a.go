package a

import "fmt"

func hoge(){
	fmt.Println("hoge")
}

func hogehoge(a int, b int, c string,d int,e int, f float32){
	fmt.Println("hogehoge")
}

func fugafuga(a int, b int, c string){
	fmt.Println("fugafuga")
}

func fugafuga2(a int, b int, c string, fuga... string){
	fmt.Println("fugafuga")
}

func fugafuga3(a int, b int, c string,d int, fuga... string){
	fmt.Println("fugafuga")
}