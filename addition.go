package main
import "fmt"
func main(){
	a:= 10
	b:= 12
	fmt.Println(test(a,b))

}

func test(number1 int, number2 int) (int, error){

	// return number1+ number2, nil
	// return number1- number2, nil
	return number1/ number2, nil
}