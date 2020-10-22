package main
import "fmt"

func generarImpar() func() uint {
	i := uint(0) 
	return func() uint {
		var num = i
		if num%2==1{
			i += 2
			return num
		}else{
			i += 3
			return num
		}

	}
}
func main(){
	siguiente := generarImpar()
	fmt.Println(siguiente())
	fmt.Println(siguiente())
	fmt.Println(siguiente())
	fmt.Println(siguiente())
}