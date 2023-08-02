package main

//import "fmt" 
//fmt "formating" is golang inbuild package


// func main(){
// fmt.Println("hello go")
// }



//How to find data type
import(
"fmt"
"reflect"
)
func main(){
var a int
// var b = 7 (int)
var b = "9" //string
c:= 20  //  :=shorthand method

fmt.Printf("a is type of %v, b is type of %v and c is type of %v",
reflect.TypeOf(a),reflect.TypeOf(b),reflect.TypeOf(c))
// 
test()
}

func test() {
	fmt.Println(" hello world ")
}