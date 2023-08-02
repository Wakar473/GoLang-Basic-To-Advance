package main
import"fmt"
func main(){
	//fallthrough= with the help of fallthrough next case will be executed or print true.
	var a interface{}
	a = "some data"
	switch t:= a.(type){
	case int64:
		fmt.Println("type an integer:", t)
	case float64:
		fmt.Println("type is a flot:",t)
	case string:
		fmt.Println("type is a string:",t)
	case nil:
		fmt.Println("type is a nil:",t)
	case bool:
		fmt.Println("type is a bool:",t)
	default:
		fmt.Println("type is unknown!")

	}
}
