package main

import "fmt"

func main(){
	var a,b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	intercambia(&a,&b)
	fmt.Println(a,b)
}
func intercambia (x*int,y*int){
	aux:=*x
	*x=*y
	*y=aux
}
