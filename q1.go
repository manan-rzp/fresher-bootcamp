package main

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	row int
	col int
	matrix [][]int
}

func new_Matrix(r int, c int) Matrix{
	m := Matrix{
		row: r,
		col: c,
	}
	m.matrix = make([][]int, r)

	for i := 0; i < m.row; i++ {
		m.matrix[i] = make([]int, c)
		for j := 0; j < m.col; j++ {
			m.set_value(i,j,rand.Intn(10000))
		}
	}
	return m
}


func (m Matrix) print() bool {
	fmt.Println("{" )

	for i := 0; i < m.row; i++ {
		fmt.Printf("%s", "{")
		for j := 0; j < m.col; j++ {
			fmt.Printf("%d,",m.matrix[i][j])
		}
		fmt.Printf("},")
	}
	fmt.Printf("}")
	return true
}

func (m Matrix) get_row() int {
	return m.row
}

func (m Matrix) get_col() int {
	return m.col
}

func (m Matrix) set_value(i, j, v int) bool {
	if i >= m.row && j >= m.col {
		return false
	}
	m.matrix[i][j] = v
	return true
}

func (m Matrix) add(M Matrix) Matrix {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.col; j++ {
			m.matrix[i][j] += M.matrix[i][j]
		}
	}
	return m
}


func main(){
	m1 := new_Matrix(4, 4)
	m2 := new_Matrix(4, 4)

	fmt.Println(m1.get_row())

	fmt.Println(m1.get_col())

	m3 := m1.add(m2)
	m3.print()
	var c byte='b'
	fmt.Printf("%q",c)
}