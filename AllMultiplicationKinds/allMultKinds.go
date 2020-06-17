package AllMultiplicationKinds

import (
	"errors"
	"math"
	"runtime"
	"sync"
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

func AsyncMultiplicationWithTranspose(matrix1, matrix2 [][]float32)([][]float32, error) {
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2[0])
	matrix1Width := len(matrix2)

	matrix2Transposed := transpose(matrix2)
	resultMatrix := make([][]float32, resultHeight)

	goroutines := runtime.NumCPU()
	linesPerRoutine := int(math.Ceil(float64(resultHeight) / float64(goroutines)))
	var wg sync.WaitGroup

	for i2 :=0; i2 < resultHeight; i2 += linesPerRoutine {
		iMax2 := int(math.Min(float64(resultHeight),float64(i2+linesPerRoutine)))
		wg.Add(1)
		go func(i, iMax int){
			defer wg.Done()
			for ; i < iMax; i++ {
				resultMatrix[i] = make([]float32, resultWidth)
				for j := 0; j < resultWidth; j++ {
					for k := 0; k < matrix1Width; k++ {
						resultMatrix[i][j] += matrix1[i][k] * matrix2Transposed[j][k]
					}
				}
			}
		}(i2,iMax2)
	}

	wg.Wait()
	return resultMatrix,nil
}
