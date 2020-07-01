package main

import (
	"errors"
	"github.com/slimsag/rand/simd"
)


func PackForSIMD(matrix [][]float64)[][][4]float64{
	resultHeight :=len(matrix)
	resultWeight :=len(matrix[0])/4

	resultMatrix := make([][][4]float64,resultHeight)

	for i:=0; i<resultHeight; i++{
		resultMatrix[i] =make([][4]float64,resultWeight)
		for j:=0; j<resultWeight; j++{
			index := j*4
			resultMatrix[i][j] = simd.Vec64{matrix[i][index],matrix[i][index+1],matrix[i][index+2],matrix[i][index+3]}
		}
	}

	return resultMatrix
}

func SIMDMultiplication(matrix1, matrix2 [][]float64)([][]float64, error) {
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2[0])
	matrix1Width := len(matrix2)/4

	matrix1Simd :=PackForSIMD(matrix1)
	matrix2TransposedSimd := PackForSIMD(Transpose(matrix2))

	resultMatrix := make([][]float64, resultHeight)

	for i := 0; i < resultHeight; i++ {
		resultMatrix[i] = make([]float64, resultWidth)
		for j := 0; j < resultWidth; j++ {
			for k:=0; k < matrix1Width; k++{
				mulArr := simd.Vec64Mul(matrix1Simd[i][k],matrix2TransposedSimd[j][k])
				resultMatrix[i][j] += mulArr[0]+mulArr[1]+mulArr[2]+mulArr[3]
			}
		}
	}

	return resultMatrix,nil
}

func SIMDMultiplication2(matrix1, matrix2 [][][4]float64)([][]float64, error) {
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2)*4
	matrix1Width := len(matrix2[0])

	resultMatrix := make([][]float64, resultHeight)

	for i := 0; i < resultHeight; i++ {
		resultMatrix[i] = make([]float64, resultWidth)
		for j := 0; j < resultWidth; j++ {
			for k:=0; k < matrix1Width; k++{
				mulArr := simd.Vec64Mul(matrix1[i][k],matrix2[j][k])
				resultMatrix[i][j] += mulArr[0]+mulArr[1]+mulArr[2]+mulArr[3]
			}
		}
	}

	return resultMatrix,nil
}
