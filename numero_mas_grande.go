package main

import "fmt"
func mas_grande(args... int)int{

	mas:=0
	for _,v:=range args{
		if mas<v{
			mas =v
		}
	}
	return mas

}

func main(){
	a := []int{2,5,4,16}
	fmt.Println(mas_grande(2,5,6,15))
	fmt.Println(mas_grande(a...))
}

