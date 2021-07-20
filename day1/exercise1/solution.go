package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	row,col int
	elements [][] int
}

func (mat Matrix) initialize()  Matrix{
	mat.elements=make([][]int, mat.row)
	//fmt.Println(mat)

	rows := make([]int, mat.col)
	for index,_ := range rows{
		rows[index]=0
	}

	for index, _ := range mat.elements {
		mat.elements[index]= make([]int, len(rows))
		copy(mat.elements[index], rows)
	}
	//fmt.Println(mat)


	return mat
}

func (m Matrix) rows() int {
	return m.row
}
func (m Matrix) cols() int {
	return m.col
}
func (m *Matrix) set(i, j, element int)  {
	m.elements[i][j] = element
}
func add(m Matrix, n Matrix)  {
	var final = Matrix{}
	final.row = m.row
	final.col = m.col
	final=final.initialize()
	if m.row == n.row && m.col == n.col {
		for i := 0; i < m.row; i++ {
			for j := 0; j < m.col; j++ {
				final.elements[i][j] = m.elements[i][j] + n.elements[i][j]
			}
		}
		for i :=0;i< final.row;i++{
		   fmt.Println(final.elements[i])
		}

		fmt.Println(json.Marshal(final))

	} else {
		fmt.Println("Matrices cannot be added")
	}
}
func main() {
	var a = Matrix{}
	var b = Matrix{}
	//var c = Matrix{}
	a.row = 2
	a.col = 3
	b.row = 2
	b.col = 3
    fmt.Println("hello")
	a=a.initialize()
	b=b.initialize()
	var i,j,value int
	fmt.Println("enter values:\n")
	fmt.Scanf("%d",&i)
	fmt.Scanf("%d",&j)
	fmt.Scanf("%d",&value)
	a.set(i,j, value)
	add(a,b)

}

