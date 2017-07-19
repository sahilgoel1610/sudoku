package main


import "fmt"
import "time"

type MatrixPointer struct {
	X int
	Y int
}

func main() {
	values := [][]int{}
	row1 := []int{1,2,3,4}
	row2 := []int{0,3,4,1}
	row3 := []int{3,4,1,2}
	row4 := []int{4,1,0,3}

	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	matrixPointer := findNextEmpty(values)
	finalResult := canFillThisPlace(copyArray(values), matrixPointer.X, matrixPointer.Y)
	fmt.Print("FINAL RESULT IF THE SODOKU IS SOLVABLE : ")
	fmt.Println(finalResult)
	time.Sleep(time.Second * 3)

}


func canFillThisPlace(matrix [][]int, i int, j int) bool {
	resultChan := make(chan bool, 4)
	for k := 1; k < 5; k++ {
		a := copyArray(matrix)
		go tryWithNumber(a, i, j, k, resultChan)
	}
	
	return <- resultChan || <- resultChan || <- resultChan || <- resultChan

}


func tryWithNumber(a [][]int, i int, j int , k int, resultChan chan bool) {
	matrix := copyArray(a)
	matrix[i][j] = k
	if conditionsAtPositionValid(matrix, i, j) {
		matrixPointer := findNextEmpty(matrix)
		if matrixPointer.X == 0 && matrixPointer.Y == 0 {
			resultChan <- true
			commitMatrix(matrix)
		} else {
			resultChan <- canFillThisPlace(matrix, matrixPointer.X, matrixPointer.Y)
		}
	} else {
		resultChan <- false
	}
}



func conditionsAtPositionValid(matrix [][]int, i int, j int) bool {
	for a := 0; a < 4; a++ {
		if a != i && matrix[a][j] == matrix[i][j] {
			return false;
		}
	}

	for b := 0; b < 4; b++ {
		if b != j && matrix[i][b] == matrix[i][j] {
			return false;
		}
	}
	
	return true
}

// Tested
func findNextEmpty(matrix [][]int) MatrixPointer {
	for i := 0; i < 4; i++ {
		for j:= 0; j < 4; j++ {
			if matrix[i][j] == 0 {
				return MatrixPointer{i ,j}
			}
		}
	}

	return MatrixPointer{0, 0}
}


func commitMatrix(matrix [][]int) {
	for i := 0; i < 4; i++ {
		fmt.Print("|")
		for j:= 0; j < 4; j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println("|")
	}
}


func copyArray(matrix [][]int) [][]int {

	copiedArray := [][]int{}
	row1 := []int{0,0,0,0}
	row2 := []int{0,0,0,0}
	row3 := []int{0,0,0,0}
	row4 := []int{0,0,0,3}

	copiedArray = append(copiedArray, row1)
	copiedArray = append(copiedArray, row2)
	copiedArray = append(copiedArray, row3)
	copiedArray = append(copiedArray, row4)
	for i := 0; i < 4; i++ {
		for j:= 0; j < 4; j++ {
			copiedArray[i][j] = matrix[i][j]
		}
	}

	return copiedArray
}