package main

import (
	"math/rand"
	"testing"
)

const (
	N = 1024
	M = 512
)

func TestNaiveMultiplication(t *testing.T){
	matrix1 := make([][]float32,2)
	matrix1[0] =[]float32 {1,2,3}
	matrix1[1] =[]float32 {4,5,6}

	matrix2 := make([][]float32,3)
	matrix2[0] =[]float32 {1,2}
	matrix2[1] =[]float32 {3,4}
	matrix2[2] =[]float32 {5,6}

	result, err := NaiveMultiplication(matrix1,matrix2)
	if err != nil{
		t.Errorf(err.Error())
	}

	expected := make([][]float32,2)
	expected[0] =[]float32 {22,28}
	expected[1] =[]float32 {49,64}

	for i:=0 ;i<2; i++{
		for j:=0; j<2; j++{
			if result[i][j] != expected[i][j]{
				t.Errorf("result in index [%v,%v] is not equal to %v", i,j,expected[i][j])
			}
		}

	}
}

func TestBlocksMultiplication(t *testing.T) {
	matrix1 := make([][]float32,2)
	matrix1[0] =[]float32 {1,2,3}
	matrix1[1] =[]float32 {4,5,6}

	matrix2 := make([][]float32,3)
	matrix2[0] =[]float32 {1,2}
	matrix2[1] =[]float32 {3,4}
	matrix2[2] =[]float32 {5,6}

	result, err := BlocksMultiplication(matrix1,matrix2,1,1,1)
	if err != nil{
		t.Errorf(err.Error())
	}

	expected := make([][]float32,2)
	expected[0] =[]float32 {22,28}
	expected[1] =[]float32 {49,64}

	for i:=0 ;i<2; i++{
		for j:=0; j<2; j++{
			if result[i][j] != expected[i][j]{
				t.Errorf("result in index [%v,%v] is not equal to %v", i, j, expected[i][j])
			}
		}
	}

}

func TestMultiplicationWithTranspose(t *testing.T){
	matrix1 := make([][]float32,2)
	matrix1[0] =[]float32 {1,2,3}
	matrix1[1] =[]float32 {4,5,6}

	matrix2 := make([][]float32,3)
	matrix2[0] =[]float32 {1,2}
	matrix2[1] =[]float32 {3,4}
	matrix2[2] =[]float32 {5,6}

	result, err := MultiplicationWithTranspose(matrix1,matrix2)
	if err != nil{
		t.Errorf(err.Error())
	}

	expected := make([][]float32,2)
	expected[0] =[]float32 {22,28}
	expected[1] =[]float32 {49,64}

	for i:=0 ;i<2; i++{
		for j:=0; j<2; j++{
			if result[i][j] != expected[i][j]{
				t.Errorf("result in index [%v,%v] is not equal to %v", i, j, expected[i][j])
			}
		}

	}
}

func TestAsyncMultiplicationWithTranspose(t *testing.T) {
	matrix1 := make([][]float32,2)
	matrix1[0] =[]float32 {1,2,3}
	matrix1[1] =[]float32 {4,5,6}

	matrix2 := make([][]float32,3)
	matrix2[0] =[]float32 {1,2}
	matrix2[1] =[]float32 {3,4}
	matrix2[2] =[]float32 {5,6}

	result, err := AsyncMultiplicationWithTranspose(matrix1,matrix2)
	if err != nil{
		t.Errorf(err.Error())
	}

	expected := make([][]float32,2)
	expected[0] =[]float32 {22,28}
	expected[1] =[]float32 {49,64}

	for i:=0 ;i<2; i++{
		for j:=0; j<2; j++{
			if result[i][j] != expected[i][j]{
				t.Errorf("result in index [%v,%v] is not equal to %v", i, j, expected[i][j])
			}
		}
	}

}

func generateRandomMatrix(n,m int)[][]float32 {
	matrix := make([][]float32,n)

	for i:=0; i<n; i++{
		matrix[i] = make([]float32,m)
		for j:=0; j<m; j++{
			matrix[i][j] = rand.Float32()+(float32(rand.Int()))
		}
	}

	return matrix
}

func BenchmarkNaiveMultiplication(b *testing.B) {
	matrix1 := generateRandomMatrix(N, M)
	matrix2 := generateRandomMatrix(M, N)

	b.ResetTimer()
	_,_ = NaiveMultiplication(matrix1,matrix2)
}

func BenchmarkBlocksMultiplication(b *testing.B) {
	matrix1 := generateRandomMatrix(N, M)
	matrix2 := generateRandomMatrix(M, N)

	b.ResetTimer()
	_,_ = BlocksMultiplication(matrix1,matrix2,256,256,256)
}

func BenchmarkMultiplicationWithTranspose(b *testing.B) {
	matrix1 := generateRandomMatrix(N, M)
	matrix2 := generateRandomMatrix(M, N)

	b.ResetTimer()
	_,_ = MultiplicationWithTranspose(matrix1,matrix2)
}

func BenchmarkAsyncMultiplicationWithTranspose(b *testing.B) {
	matrix1 := generateRandomMatrix(N, M)
	matrix2 := generateRandomMatrix(M, N)

	b.ResetTimer()
	_,_ = AsyncMultiplicationWithTranspose(matrix1,matrix2)
}
