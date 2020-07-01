package main

import (
	"errors"
)

func Transpose(matrix [][]float64)[][]float64{
	resultHeight :=len(matrix[0])
	resultWeight :=len(matrix)

	resultMatrix := make([][]float64,resultHeight)

	for i:=0; i<resultHeight; i++{
		resultMatrix[i] =make([]float64,resultWeight)
		for j:=0; j<resultWeight; j++{
			resultMatrix[i][j] = matrix[j][i]
		}
	}

	return resultMatrix
}

func MultiplicationWithTranspose(matrix1, matrix2 [][]float64)([][]float64, error) {

	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2[0])
	matrix1Width := len(matrix2)

	matrix2Transposed := Transpose(matrix2)
	resultMatrix := make([][]float64, resultHeight)

	for i := 0; i < resultHeight; i++ {
		resultMatrix[i] = make([]float64, resultWidth)
		for j := 0; j < resultWidth; j++ {
			for k := 0; k < matrix1Width; k++ {
				resultMatrix[i][j] += matrix1[i][k] * matrix2Transposed[j][k]
			}
		}
	}
	return resultMatrix,nil
}
