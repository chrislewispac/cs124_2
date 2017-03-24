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
}

func main() {
	custom, _ := strconv.Atoi(os.Args[1])
	dimensions, _ := strconv.Atoi(os.Args[2])
	inputfile := os.Args[3]

	ret := BuildMatricesFromInput(dimensions, inputfile)

	if custom == 1 { //run classic
		_ = ClassicMatrixMult_K_First(&ret, dimensions)
	} else if custom == 2 { //run strassen
		_ = StrassenMatrixMult(ret.a, ret.b)
	} else if custom == 0 {
		r := ClassicMatrixMult_K_First(&ret, dimensions)
		PrintSpecOutput(&r, dimensions)
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

func StrassenMatrixMult(a [][]int, b [][]int) [][]int {
	dimensions := len(a[0])
	new_dimensions := dimensions / 2
	c := [][]int{}
	count := 0
	for count < dimensions {
		y := make([]int, dimensions)
		c = append(c, y)
		count++
	}

	if dimensions == 4 {
		b11 := [][]int{}
		b12 := [][]int{}
		b21 := [][]int{}
		b22 := [][]int{}
		a11 := [][]int{}
		a12 := [][]int{}
		a21 := [][]int{}
		a22 := [][]int{}

		for i := 0; i < dimensions; i++ {
			if i < new_dimensions {
				b11 = append(b11, b[i][:new_dimensions])
				b12 = append(b12, b[i][new_dimensions:])
				a11 = append(a11, a[i][:new_dimensions])
				a12 = append(a12, a[i][new_dimensions:])
			} else {
				b21 = append(b21, b[i][:new_dimensions])
				b22 = append(b22, b[i][new_dimensions:])
				a21 = append(a21, a[i][:new_dimensions])
				a22 = append(a22, a[i][new_dimensions:])
			}

		}

		c11 := AddMatrices(ClassicMatrixMult(a11, b11), ClassicMatrixMult(a12, b21))
		c12 := AddMatrices(ClassicMatrixMult(a11, b12), ClassicMatrixMult(a12, b22))
		c21 := AddMatrices(ClassicMatrixMult(a21, b11), ClassicMatrixMult(a22, b21))
		c22 := AddMatrices(ClassicMatrixMult(a21, b12), ClassicMatrixMult(a22, b22))

		y := 0
		for x := 0; x < dimensions; x++ {
			if x < (dimensions / 2) {
				c[x] = append(c11[x], c12[x]...)
			} else {
				c[x] = append(c21[y], c22[y]...)
				y++
			}
		}
		return c
	} else {

		//build submatrices
		b11 := [][]int{}
		b12 := [][]int{}
		b21 := [][]int{}
		b22 := [][]int{}
		a11 := [][]int{}
		a12 := [][]int{}
		a21 := [][]int{}
		a22 := [][]int{}
		for i := 0; i < dimensions; i++ {
			if i < new_dimensions {
				b11 = append(b11, b[i][:new_dimensions])
				b12 = append(b12, b[i][new_dimensions:])
				a11 = append(a11, a[i][:new_dimensions])
				a12 = append(a12, a[i][new_dimensions:])
			} else {
				b21 = append(b21, b[i][:new_dimensions])
				b22 = append(b22, b[i][new_dimensions:])
				a21 = append(a21, a[i][:new_dimensions])
				a22 = append(a22, a[i][new_dimensions:])
			}

		}

		s1 := SubtractMatrices(b12, b22)
		s2 := AddMatrices(a11, a12)
		s3 := AddMatrices(a21, a22)
		s4 := SubtractMatrices(b21, b11)
		s5 := AddMatrices(a11, a22)
		s6 := AddMatrices(b11, b22)
		s7 := SubtractMatrices(a12, a22)
		s8 := AddMatrices(b21, b22)
		s9 := SubtractMatrices(a11, a21)
		s10 := AddMatrices(b11, b12)

		p1 := StrassenMatrixMult(a11, s1)
		p2 := StrassenMatrixMult(s2, b22)
		p3 := StrassenMatrixMult(s3, b11)
		p4 := StrassenMatrixMult(a22, s4)
		p5 := StrassenMatrixMult(s5, s6)
		p6 := StrassenMatrixMult(s7, s8)
		p7 := StrassenMatrixMult(s9, s10)

		n1 := AddMatrices(p5, p4)
		n2 := SubtractMatrices(n1, p2)
		n3 := AddMatrices(n2, p6)

		n4 := AddMatrices(p5, p1)
		n5 := SubtractMatrices(n4, p3)
		n6 := SubtractMatrices(n5, p7)

		c11 := n3
		c12 := AddMatrices(p1, p2)
		c21 := AddMatrices(p3, p4)
		c22 := n6

		y := 0
		for x := 0; x < dimensions; x++ {
			if x < (dimensions / 2) {
				c[x] = append(c11[x], c12[x]...)
			} else {
				c[x] = append(c21[y], c22[y]...)
				y++
			}
		}

		return c
	}
}

func ClassicMatrixMult_K_First(matrix_set *MatrixFun, dimensions int) MatrixFun {
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
