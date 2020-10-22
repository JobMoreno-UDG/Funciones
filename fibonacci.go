package main
import "fmt"

func fibonacci(a uint)uint{
	if a<2{
		return 1
	}else{
		return (fibonacci(a-1)+fibonacci(a-2))
	}
}

func main(){
	fmt.Println(fibonacci(5))
}