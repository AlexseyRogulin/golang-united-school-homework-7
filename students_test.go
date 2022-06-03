package coverage

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPersonSort(t *testing.T) {
	input := People{
		Person{birthDay: time.Unix(10, 0)},
		Person{birthDay: time.Unix(20, 0)},
		Person{firstName: "Alex", lastName: "Antonov"},
		Person{firstName: "Ben", lastName: "Antonov"},
		Person{firstName: "Ivan", lastName: "Sidorov"},
		Person{firstName: "Ivan", lastName: "Petrov"},
	}
	want := People{
		Person{birthDay: time.Unix(20, 0)},
		Person{birthDay: time.Unix(10, 0)},
		Person{firstName: "Alex", lastName: "Antonov"},
		Person{firstName: "Ben", lastName: "Antonov"},
		Person{firstName: "Ivan", lastName: "Petrov"},
		Person{firstName: "Ivan", lastName: "Sidorov"},
	}

	sort.Sort(input)
	for i := 0; i < len(input); i++ {
		if input[i] != want[i] {
			t.Errorf("want %v, got %v", want, input)
			return
		}
	}
}

func TestMatrixCreate(t *testing.T) {
	tests := map[string]struct {
		input      string
		wantMatrix *Matrix
		wantError  error
	}{
		"success res": {
			input:      "1 2 3\n4 5 6",
			wantMatrix: &Matrix{rows: 2, cols: 3, data: []int{1, 2, 3, 4, 5, 6}},
			wantError:  nil,
		},
		"different rows len": {
			input:      "1 2 3\n4 5 6 7",
			wantMatrix: nil,
			wantError:  fmt.Errorf("Rows need to be the same length"),
		},
		"letter in input": {
			input:      "1 2 3\n4 A 6",
			wantMatrix: nil,
			wantError:  strconv.ErrSyntax,
		},
	}

	for _, test := range tests {
		gotMatrix, gotError := New(test.input)
		if gotError != nil && test.wantError == nil || gotError == nil && test.wantError != nil {
			t.Errorf("want %v error got %v error", test.wantError, gotError)
			return
		}
		if gotError != nil && test.wantError != nil && test.wantError.Error() != gotError.Error() && !errors.Is(gotError, test.wantError) {
			t.Errorf("want %v error got %v error", test.wantError, gotError)
			return
		}
		if !equalMatrix(test.wantMatrix, gotMatrix) {
			t.Errorf("want %v matrix got %v matrix", test.wantMatrix, gotMatrix)
			return
		}
	}
}

func equalMatrix(a, b *Matrix) bool {
	if a == b {
		return true
	}
	if a.rows != b.rows || a.cols != b.cols || len(a.data) != len(b.data) {
		return false
	}
	for i := 0; i < len(a.data); i++ {
		if a.data[i] != b.data[i] {
			return false
		}
	}
	return true
}

func TestMatrixRows(t *testing.T) {
	input := Matrix{rows: 3, cols: 3, data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	want := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	got := input.Rows()
	if !equalsTwoDimensionalArray(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestMatrixCols(t *testing.T) {
	input := Matrix{rows: 3, cols: 3, data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	want := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	got := input.Cols()
	if !equalsTwoDimensionalArray(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func equalsTwoDimensionalArray(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestMatrixSet(t *testing.T) {
	tests := map[string]struct {
		inputMatrix                  Matrix
		inputRow, inputCol, inputVal int
		want                         bool
	}{
		"invalid input row": {
			inputMatrix: Matrix{rows: 2, cols: 2, data: []int{1, 2, 3, 4}},
			inputRow:    3, inputCol: 3, inputVal: 3,
			want: false,
		},
		"valid input": {
			inputMatrix: Matrix{rows: 2, cols: 2, data: []int{1, 2, 3, 4}},
			inputRow:    1, inputCol: 1, inputVal: 5,
			want: true,
		},
	}

	for _, test := range tests {
		gotRes := test.inputMatrix.Set(test.inputRow, test.inputCol, test.inputVal)
		if test.want && gotRes && test.inputVal != test.inputMatrix.Rows()[test.inputRow][test.inputCol] {
			t.Errorf("set %v in [%v,%v] does not affect matrix %v",
				test.inputVal, test.inputRow, test.inputCol, test.inputMatrix.Rows())
			return
		} else if test.want != gotRes {
			t.Errorf("want %v, got %v", test.want, gotRes)
			return
		}
	}
}
