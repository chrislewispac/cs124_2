package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type MatrixFun struct {
	a   [][]int
	b   [][]int
	res [][]int
	a11 [][]int
	a12 [][]int
	a21 [][]int
	a22 [][]int
	b11 [][]int
	b12 [][]int
	b21 [][]int
	b22 [][]int
	c11 [][]int
	c12 [][]int
	c21 [][]int
	c22 [][]int
	p1  [][]int
	p2  [][]int
	p3  [][]int
	p4  [][]int
	p5  [][]int
	p6  [][]int
	p7  [][]int
}

func main() {
	custom, _ := strconv.Atoi(os.Args[1])
	dimensions, _ := strconv.Atoi(os.Args[2])
	inputfile := os.Args[3]

	ret := BuildMatricesFromInput(dimensions, inputfile)

	if custom == 1 { //run classic
		_ = ClassicMatrixMult(ret.a, ret.b)
	} else if custom == 2 { //run strassen
		_ = StrassenMatrixMult(&ret, 16)
	} else if custom == 0 { //TODO MAKE THIS RUN THE OPTIMIZED VERSION OF STRASSEN
		c := ClassicMatrixMult(ret.a, ret.b)
		ret.res = c
		PrintSpecOutput(&ret, dimensions)
	} else {
		_ = StrassenMatrixMult(&ret, custom)
	}

	os.Remove(inputfile)
}

func BuildMatricesFromInput(dimensions int, inputfile string) MatrixFun {
	rand.Seed(time.Now().Unix())
	var _, err = os.Stat(inputfile)
	if os.IsNotExist(err) {
		var file, _ = os.Create(inputfile)
		defer file.Close()
	}

	f, err := os.OpenFile(inputfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	count := 0

	for count < (2 * (dimensions * dimensions)) {
		str := fmt.Sprintf("%d", Random(1, 10))
		if _, err = f.WriteString(str + "\n"); err != nil {
			panic(err)
		}
		count++
	}
	b, _ := ioutil.ReadFile(inputfile)
	tf := string(b)
	ints, _ := ReadInts(strings.NewReader(tf))
	ret := BuildMatrices(dimensions, ints)

	return ret
}

func PrintSpecOutput(matrix_set *MatrixFun, dimensions int) {
	for i := 0; i < dimensions; i++ {
		for j := 0; j < dimensions; j++ {
			if i == j {
				fmt.Println(matrix_set.res[i][j])
			}
		}
	}
}

func AddMatrices(a [][]int, b [][]int) [][]int {
	c := [][]int{}
	count := 0
	dimensions := len(a[0])
	for count < dimensions {
		y := make([]int, dimensions)
		c = append(c, y)
		count++
	}

	for i := 0; i < dimensions; i++ {
		for j := 0; j < dimensions; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c
}

func SubtractMatrices(a [][]int, b [][]int) [][]int {
	c := [][]int{}
	count := 0
	dimensions := len(a[0])
	for count < dimensions {
		y := make([]int, dimensions)
		c = append(c, y)
		count++
	}

	for i := 0; i < dimensions; i++ {
		for j := 0; j < dimensions; j++ {
			c[i][j] = a[i][j] - b[i][j]
		}
	}

	return c
}

func StrassenMatrixMult(matrix_set *MatrixFun, transition int) MatrixFun {

	dimensions := len(matrix_set.a[0])
	new_dimensions := dimensions / 2

	if dimensions <= transition {
		return ClassicMatrixMult_K_First(matrix_set)
	} else {
		for i := 0; i < dimensions; i++ {
			if i < new_dimensions {
				matrix_set.b11 = append(matrix_set.b11, matrix_set.b[i][:new_dimensions])
				matrix_set.b12 = append(matrix_set.b12, matrix_set.b[i][new_dimensions:])
				matrix_set.a11 = append(matrix_set.a11, matrix_set.a[i][:new_dimensions])
				matrix_set.a12 = append(matrix_set.a12, matrix_set.a[i][new_dimensions:])
			} else {
				matrix_set.b21 = append(matrix_set.b21, matrix_set.b[i][:new_dimensions])
				matrix_set.b22 = append(matrix_set.b22, matrix_set.b[i][new_dimensions:])
				matrix_set.a21 = append(matrix_set.a21, matrix_set.a[i][:new_dimensions])
				matrix_set.a22 = append(matrix_set.a22, matrix_set.a[i][new_dimensions:])
			}

		}

		matrix_set.p1 = StrassenMatrixMult(&MatrixFun{a: matrix_set.a11, b: SubtractMatrices(matrix_set.b12, matrix_set.b22)}, transition).res
		matrix_set.p2 = StrassenMatrixMult(&MatrixFun{a: AddMatrices(matrix_set.a11, matrix_set.a12), b: matrix_set.b22}, transition).res
		matrix_set.p3 = StrassenMatrixMult(&MatrixFun{a: AddMatrices(matrix_set.a21, matrix_set.a22), b: matrix_set.b11}, transition).res
		matrix_set.p4 = StrassenMatrixMult(&MatrixFun{a: matrix_set.a22, b: SubtractMatrices(matrix_set.b21, matrix_set.b11)}, transition).res
		matrix_set.p5 = StrassenMatrixMult(&MatrixFun{a: AddMatrices(matrix_set.a11, matrix_set.a22), b: AddMatrices(matrix_set.b11, matrix_set.b22)}, transition).res
		matrix_set.p6 = StrassenMatrixMult(&MatrixFun{a: SubtractMatrices(matrix_set.a12, matrix_set.a22), b: AddMatrices(matrix_set.b21, matrix_set.b22)}, transition).res
		matrix_set.p7 = StrassenMatrixMult(&MatrixFun{a: SubtractMatrices(matrix_set.a11, matrix_set.a21), b: AddMatrices(matrix_set.b11, matrix_set.b12)}, transition).res

		matrix_set.c11 = AddMatrices(SubtractMatrices(AddMatrices(matrix_set.p5, matrix_set.p4), matrix_set.p2), matrix_set.p6)
		matrix_set.c12 = AddMatrices(matrix_set.p1, matrix_set.p2)
		matrix_set.c21 = AddMatrices(matrix_set.p3, matrix_set.p4)
		matrix_set.c22 = SubtractMatrices(SubtractMatrices(AddMatrices(matrix_set.p5, matrix_set.p1), matrix_set.p3), matrix_set.p7)

		matrix_set.res = ZeroMatrix(dimensions)
		y := 0
		for x := 0; x < dimensions; x++ {
			if x < (dimensions / 2) {
				matrix_set.res[x] = append(matrix_set.c11[x], matrix_set.c12[x]...)
			} else {
				matrix_set.res[x] = append(matrix_set.c21[y], matrix_set.c22[y]...)
				y++
			}
		}

		return *matrix_set
	}
}

func ClassicMatrixMult_K_First(matrix_set *MatrixFun) MatrixFun {
	dimensions := len(matrix_set.a)
	matrix_set.res = ZeroMatrix(dimensions)
	for i := 0; i < dimensions; i++ {
		for k := 0; k < dimensions; k++ {
			for j := 0; j < dimensions; j++ {
				matrix_set.res[i][j] += (matrix_set.a[i][k] * matrix_set.b[k][j])
			}
		}
	}
	return *matrix_set
}

func ClassicMatrixMult(a [][]int, b [][]int) [][]int {
	c := [][]int{}
	count := 0
	dimensions := len(a[0])
	for count < dimensions {
		y := make([]int, dimensions)
		c = append(c, y)
		count++
	}
	for i := 0; i < dimensions; i++ {
		for j := 0; j < dimensions; j++ {
			for k := 0; k < dimensions; k++ {
				c[i][j] += (a[i][k] * b[k][j])
			}
		}
	}
	return c
}

func BuildMatrices(dimensions int, ints []int) MatrixFun {
	matrix_set := MatrixFun{}
	count := 0
	prev := 0
	next := dimensions

	for count < dimensions {
		y := ints[prev:next]
		matrix_set.a = append(matrix_set.a, y)
		prev = next
		next = next + dimensions
		count++
	}

	count = 0
	for count < dimensions {
		y := ints[prev:next]
		matrix_set.b = append(matrix_set.b, y)
		prev = next
		next = next + dimensions
		count++
	}

	count = 0
	for count < dimensions {
		y := make([]int, dimensions)
		matrix_set.res = append(matrix_set.res, y)
		prev = next
		next = next + dimensions
		count++
	}

	return matrix_set
}

func ZeroMatrix(dimensions int) [][]int {
	count := 0
	res := [][]int{}
	for count < dimensions {
		y := make([]int, dimensions)
		res = append(res, y)
		count++
	}
	return res
}

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}
