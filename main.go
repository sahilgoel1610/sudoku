package main


import "fmt"
import "time"

type MatrixPointer struct {
	X int
	Y int
}


func main() {
	values := [9][9]int{
		{0,0,4,  0,9,6,  0,1,8},
		{3,8,6,  0,0,4,  9,5,0},
		{9,5,1,  0,3,0,  6,0,4},

		{0,4,3,  0,0,0,  5,0,9},
		{6,2,0,  0,0,0,  0,4,0},
		{7,9,0,  4,0,2,  1,6,0},
		
		{0,0,7,	 6,2,1,	 0,9,5},
		{0,0,0,  7,0,5,  0,3,0},
		{0,6,2,  3,4,0,  7,8,1}}

	// printMatrix(values)
	resultMatrix := make(chan [9][9]int)
	matrixPointer := findNextEmpty(values)
	finalResult := canFillThisPlace(copyArray(values), matrixPointer.X, matrixPointer.Y, resultMatrix)
	fmt.Print("FINAL RESULT IF THE SODOKU IS SOLVABLE : ")
	fmt.Println(finalResult)
	commitMatrix(<- resultMatrix)
	// fmt.Printf("final Matrix", <- resultMatrix)
	time.Sleep(time.Second * 13)

}


func canFillThisPlace(matrix [9][9]int, i int, j int, resultMatrix chan [9][9]int) bool {
	resultChan := make(chan bool, 9)
	for k := 1; k < 10; k++ {
		a := copyArray(matrix)
		go tryWithNumber(a, i, j, k, resultChan, resultMatrix)
	}
	
	
	for l := 0; l < 9; l++ {
		if <- resultChan {
			return true
		}
	}

	return false

}


func tryWithNumber(a [9][9]int, i int, j int , k int, resultChan chan bool, resultMatrix chan [9][9]int) {
	matrix := copyArray(a)
	matrix[i][j] = k
	if conditionsAtPositionValid(matrix, i, j) {
		matrixPointer := findNextEmpty(matrix)
		if matrixPointer.X == 0 && matrixPointer.Y == 0 {
			resultChan <- true
			resultMatrix <- matrix
		} else {
			resultChan <- canFillThisPlace(matrix, matrixPointer.X, matrixPointer.Y, resultMatrix)
		}
	} else {
		resultChan <- false
	}
}



func conditionsAtPositionValid(matrix [9][9]int, i int, j int) bool {
	for a := 0; a < 9; a++ {
		if a != i && matrix[a][j] == matrix[i][j] {
			return false;
		}
	}

	for b := 0; b < 9; b++ {
		if b != j && matrix[i][b] == matrix[i][j] {
			return false;
		}
	}
	
	return true
}

// Tested
func findNextEmpty(matrix [9][9]int) MatrixPointer {
	for i := 0; i < 9; i++ {
		for j:= 0; j < 9; j++ {
			if matrix[i][j] == 0 {
				return MatrixPointer{i ,j}
			}
		}
	}

	return MatrixPointer{0, 0}
}


func commitMatrix(matrix [9][9]int) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j:= 0; j < 9; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("|")
	}
}


func copyArray(matrix [9][9]int) [9][9]int {

	copiedArray := [9][9]int{}
	
	for i := 0; i < 9; i++ {
		for j:= 0; j < 9; j++ {
			copiedArray[i][j] = matrix[i][j]
		}
	}

	return copiedArray
}