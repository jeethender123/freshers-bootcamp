package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	row,col int
	elements [2][3] int
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
	if m.row == n.row && m.col == n.col {
		for i := 0; i < m.row; i++ {
			for j := 0; j < m.col; j++ {
				final.elements[i][j] = m.elements[i][j] + n.elements[i][j]
			}
		}
		fmt.Println(final)
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
	var i,j,value int
	fmt.Println("enter values:\n")
	fmt.Scanf("%d",&i)
	fmt.Scanf("%d",&j)
	fmt.Scanf("%d",&value)
	a.set(i,j, value)
	add(a,b)

}

