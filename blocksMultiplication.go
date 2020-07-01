package main

import (
	"errors"
	"math"
)


func BlocksMultiplication(matrix1, matrix2 [][]float64, blockHeight, blockWidth, kChunkSize int)([][]float64, error) {

	if len(matrix1[0]) != len(matrix2){
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight :=len(matrix1)
	resultWidth :=len(matrix2[0])
	matrix1Width := len(matrix2)

	resultMatrix := make([][]float64,resultHeight)
	for i:=0; i< resultHeight; i++{
		resultMatrix[i] = make([]float64, resultWidth)
	}

	for i2:=0 ; i2<resultHeight; i2+=blockHeight {
		iMax := int(math.Min(float64(i2+blockWidth), float64(resultHeight)))
		for j2 := 0; j2 < resultHeight; j2 += blockWidth {
			jMax := int(math.Min(float64(j2+blockWidth), float64(resultWidth)))
			for k2 := 0; k2 < matrix1Width; k2 += kChunkSize {
				kMax := int(math.Min(float64(k2+kChunkSize), float64(matrix1Width)))
				for i := i2; i < iMax; i++ {
					for j := j2; j < jMax; j++ {
						for k := k2; k< kMax; k++ {
							resultMatrix[i][j] += matrix1[i][k] * matrix2[k][j]
						}
					}
				}
			}
		}
	}

	return  resultMatrix,nil
}
