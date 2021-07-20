package main
import (
	"fmt"
	"math/rand"
)
type Employee interface {
	salary() int
}
type Full_time struct {
	pay_d int
}

type Contactor struct {
	pay_d int
}

type Freelancer struct {
	pay_h int
}

func (e Full_time) salary() int {
	days := rand.Intn(28)
	return days * e.pay_d
}
func (e Contactor) salary() int {
	days := rand.Intn(28)
	return days * e.pay_d
}
func (e Freelancer) salary() int {
	hours := rand.Intn(24)
	days := rand.Intn(28)
	return days * hours * e.pay_h
}
func main(){
	ft := Full_time{pay_d:500}

	co := Contactor{pay_d: 100}

	frl := Freelancer{pay_h: 10}

	data := [3]Employee{ft, co, frl}

	for _, elem := range data {
			fmt.Println(elem.salary())
	}
}