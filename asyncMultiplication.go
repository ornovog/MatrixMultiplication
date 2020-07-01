package main

import (
	"errors"
	"math"
	"runtime"
	"sync"
)

func AsyncMultiplicationWithTranspose(matrix1, matrix2 [][]float64)([][]float64, error) {
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("width of matrix1 is not equal to height og matrix2")
	}

	resultHeight := len(matrix1)
	resultWidth := len(matrix2[0])
	matrix1Width := len(matrix2)

	matrix2Transposed := Transpose(matrix2)
	resultMatrix := make([][]float64, resultHeight)

	goroutines := runtime.NumCPU()
	linesPerRoutine := int(math.Ceil(float64(resultHeight) / float64(goroutines)))
	var wg sync.WaitGroup

	for i2 :=0; i2 < resultHeight; i2 += linesPerRoutine {
		iMax2 := int(math.Min(float64(resultHeight),float64(i2+linesPerRoutine)))
		wg.Add(1)
		go func(i, iMax int){
			defer wg.Done()
			for ; i < iMax; i++ {
				resultMatrix[i] = make([]float64, resultWidth)
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

