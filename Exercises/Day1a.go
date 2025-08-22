package main

import (
	"encoding/json"
	"fmt"
)

// created a struct called matrix that has three components - no of rows ; no of cols; and the ele
type Matrix struct {
	Rows int
	Cols int
	Ele  [][]int
}

func newMatrix(rows, cols int) Matrix {
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols)
	}
	return Matrix{Rows: rows, Cols: cols, Ele: data}
}

func getRows(m Matrix) int {
	return m.Rows
}

func getCols(m Matrix) int {
	return m.Cols
}

func (m *Matrix) SetElement(i, j, ele int) {
	if i >= 0 && i < m.Rows && j >= 0 && j < m.Cols {
		m.Ele[i][j] = ele
	} else {
		fmt.Println("Index out of Range")
	}
}

func (m Matrix) addMatrix(other Matrix) Matrix {
	result := newMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Ele[i][j] = m.Ele[i][j] + other.Ele[i][j]
		}
	}
	return result
}

func (m Matrix) ToJSON() {
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(bytes))
}

func main() {
	m1 := newMatrix(2, 2)
	m2 := newMatrix(2, 2)

	m1.SetElement(0, 0, 1)
	m1.SetElement(0, 1, 2)
	m1.SetElement(1, 0, 3)
	m1.SetElement(1, 1, 4)

	m2.SetElement(0, 0, 5)
	m2.SetElement(0, 1, 6)
	m2.SetElement(1, 0, 7)
	m2.SetElement(1, 1, 8)

	sum := m1.addMatrix(m2)

	// Print as JSON
	fmt.Println("Matrix 1:")
	m1.ToJSON()

	fmt.Println("Matrix 2:")
	m2.ToJSON()

	fmt.Println("Sum of matrices:")
	sum.ToJSON()
}
