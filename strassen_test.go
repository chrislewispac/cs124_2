package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSpec(t *testing.T) {

	Convey("Given the test input file", t, func() {

		Convey("The first array should contain 2(INPUT*INPUT) numbers", func() {
			input := 4
			b, _ := ioutil.ReadFile("test.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			So(ints, ShouldHaveLength, 2*(input*input))
		})

		Convey("Each row should contain INPUT numbers", func() {
			input := 4
			b, _ := ioutil.ReadFile("test.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(4, ints)
			So(matrix.a, ShouldHaveLength, input)
			So(matrix.a[0], ShouldHaveLength, input)
			So(matrix.b, ShouldHaveLength, input)
			So(matrix.b[0], ShouldHaveLength, input)
		})

		Convey("CLASSIC Identity matrix times any matrix A should equal A", func() {
			input := 4
			b, _ := ioutil.ReadFile("test.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(4, ints)
			res := ClassicMatrixMult_K_First(&matrix)
			for i := 0; i < input; i++ {
				for j := 0; j < input; j++ {
					So(res.a[i][j], ShouldEqual, res.res[i][j])
				}
			}
		})

		Convey("STRASSEN Identity matrix times any matrix A should equal A", func() {
			input := 8
			b, _ := ioutil.ReadFile("test_8.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(input, ints)
			res := StrassenMatrixMult(&matrix, 4)
			for i := 0; i < input; i++ {
				for j := 0; j < input; j++ {
					So(res.res[i][j], ShouldEqual, matrix.a[i][j])
				}
			}
		})

		Convey("I can deconstruct and reconstruct a matrix", func() {
			d, _ := ioutil.ReadFile("test.txt")
			tf := string(d)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(4, ints)
			a := matrix.a
			b := matrix.b
			dimensions := len(a[0])
			new_dimensions := dimensions / 2
			c := [][]int{}
			count := 0
			for count < dimensions {
				y := make([]int, dimensions)
				c = append(c, y)
				count++
			}
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
			y := 0
			for x := 0; x < dimensions; x++ {
				if x < (dimensions / 2) {
					c[x] = append(a11[x], a12[x]...)
				} else {
					c[x] = append(a21[y], a22[y]...)
					y++
				}
			}

			for i := 0; i < dimensions; i++ {
				for j := 0; j < dimensions; j++ {
					So(c[i][j], ShouldEqual, matrix.a[i][j])
				}
			}

		})

		Convey("Adding two matrices together should work", func() {
			input := 4
			b, _ := ioutil.ReadFile("test_sum.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(4, ints)
			c := AddMatrices(matrix.a, matrix.b)
			for i := 0; i < input; i++ {
				for j := 0; j < input; j++ {
					So(c[i][j], ShouldEqual, 2)
				}
			}
		})

		Convey("Subtracting two matrices should work", func() {
			input := 4
			b, _ := ioutil.ReadFile("test_sum.txt")
			tf := string(b)
			ints, _ := ReadInts(strings.NewReader(tf))
			matrix := BuildMatrices(4, ints)
			c := SubtractMatrices(matrix.a, matrix.b)
			for i := 0; i < input; i++ {
				for j := 0; j < input; j++ {
					So(c[i][j], ShouldEqual, 0)
				}
			}
		})

	})

}
