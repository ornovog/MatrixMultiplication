package AllMultiplicationKinds

import (
	"errors"
	"math"
	"unsafe"
)

func NaiveMultiplication(matrix1, matrix2 [][]float32)([][]float32, error) {

	if len(matrix1[0]) != len(matrix2){
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight :=len(matrix1)
	resultWidth :=len(matrix2[0])
	matrix1Width := len(matrix2)

	resultMatrix := make([][]float32,resultHeight)

	for i:=0; i<resultHeight; i++{
		resultMatrix[i] = make([]float32,resultWidth)
		for j:=0; j<resultWidth; j++{
			for k:=0; k<matrix1Width; k++{
				resultMatrix[i][j] += matrix1[i][k]*matrix2[k][j]
			}
		}
	}
	return  resultMatrix,nil
}

func BlocksMultiplication(matrix1, matrix2 [][]float32, blockHeight, blockWidth, kChunkSize int)([][]float32, error) {

	if len(matrix1[0]) != len(matrix2){
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight :=len(matrix1)
	resultWidth :=len(matrix2[0])
	matrix1Width := len(matrix2)

	resultMatrix := make([][]float32,resultHeight)
	for i:=0; i< resultHeight; i++{
		resultMatrix[i] = make([]float32, resultWidth)
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

func transpose(matrix [][]float32)[][]float32{
	resultHeight :=len(matrix[0])
	resultWeight :=len(matrix)

	resultMatrix := make([][]float32,resultHeight)

	for i:=0; i<resultHeight; i++{
		resultMatrix[i] =make([]float32,resultWeight)
		for j:=0; j<resultWeight; j++{
			resultMatrix[i][j] = matrix[j][i]
		}
	}

	return resultMatrix
}

func MultiplicationWithTranspose(matrix1, matrix2 [][]float32)([][]float32, error) {

	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2[0])
	matrix1Width := len(matrix2)

	matrix2Transposed := transpose(matrix2)
	resultMatrix := make([][]float32, resultHeight)

	for i := 0; i < resultHeight; i++ {
		resultMatrix[i] = make([]float32, resultWidth)
		for j := 0; j < resultWidth; j++ {
			for k := 0; k < matrix1Width; k++ {
				resultMatrix[i][j] += matrix1[i][k] * matrix2Transposed[j][k]
			}
		}
	}
	return resultMatrix,nil
}
var (
	sumFloat64 func([]float64) float64
)

func SumFloat64(v []float64) float64 {
	return sumFloat64(v)
}

func sum_float64_go(buf []float64) float64 {
	acc := float64(0)
	for i := range buf {
		acc += buf[i]
	}
	return acc
}

func sum_float64_go_unroll4(buf []float64) float64 {
	var (
		acc0, acc1, acc2, acc3 float64
	)

	for i := 0; i < len(buf); i += 4 {
		bb := (*[4]float64)(unsafe.Pointer(&buf[i]))
		acc0 += bb[0]
		acc1 += bb[1]
		acc2 += bb[2]
		acc3 += bb[3]
	}
	return acc0 + acc1 + acc2 + acc3
}

func sum_float64_go_unroll8(buf []float64) float64 {
	var (
		acc0, acc1, acc2, acc3 float64
		acc4, acc5, acc6, acc7 float64
	)
	for i := 0; i < len(buf); i += 8 {
		bb := (*[8]float64)(unsafe.Pointer(&buf[i]))
		acc0 += bb[0]
		acc1 += bb[1]
		acc2 += bb[2]
		acc3 += bb[3]
		acc4 += bb[4]
		acc5 += bb[5]
		acc6 += bb[6]
		acc7 += bb[7]
	}
	return acc0 + acc1 + acc2 + acc3 + acc4 + acc5 + acc6 + acc7
}