package main
	import"fmt"
	const(
		// first =iota      //iota means print from element first
		 first =iota+1
		second 
		//  second ='a'   //here 'a'denotes a "skycode" and tha value of 'a'skycode is 97
		third
		fourth
	)
	func main(){
		a:=10
		b:=32
		fmt.Println(test(a,b))
		fmt.Println(first,second, third,fourth)
	}
	func test(number1 int, number2 int) (int, error){
		return number1 + number2, nil
	}
	