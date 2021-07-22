package main
import (
	"fmt"
	"strings"
	"time"
)

func CountChar(str string, results chan<- [26]int) {
	var cnt [26]int
	for i := 0; i < 26; i++ {
		cnt[i] += strings.Count(str, string('a'+i))
	}
	results <- cnt
}

func main() {

	Input := make([]string, 0)

	var n int =3

	//fmt.Scanf("%d",n)
	//for i:=0;i<n;i++{
	// var s string
	// fmt.Scanf("%s",s)
	// Input=append(Input,s)
	//}

	Results := make(chan [26]int, n)

	Input = append(Input, "Manan")

	Input = append(Input, "Agarwal")

	Input = append(Input, "Razorpay")

	for i := 0; i < 3; i++ {
		go CountChar(Input[i], Results)
	}

	var Count [26]int

	time.Sleep(time.Second)

	close(Results)

	for result:= range Results {

		for i := 0; i < 26; i++ {
			Count[i] += result[i]
		}
	}
	fmt.Println("{")
	for i := 0; i < 26; i++ {
		fmt.Printf("	%c : %d\n", 'a'+i, Count[i])
	}
	fmt.Println("}")
}