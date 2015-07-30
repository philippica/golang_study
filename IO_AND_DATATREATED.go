package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Point struct {
	x          []float64
	y          []float64
	z          []float64
	Point_size int
}

func (instance *Point) addPoint(px float64, py float64, pz float64) {
	instance.x = append(instance.x, px)
	instance.y = append(instance.y, py)
	instance.z = append(instance.z, pz)
}

func main() {
	var P Point
	fmt.Print("hello world")
	f, _ := os.Open("a.txt")
	defer f.Close()
	for {
		r := bufio.NewReader(f)
		s, ok := r.ReadString('\n')
		if ok == io.EOF {
			break
		}
		var magic rune
		var x, y, z float64
		ss := strings.NewReader(s)
		fmt.Fscanf(ss, "%c", &magic)
		if magic == 'v' {
			fmt.Fscanf(ss, "%c%f%f%f", &magic, &x, &y, &z)
			P.addPoint(x, y, z)
		}
		if magic == 'f' {
			fmt.Fscanf(ss, "%c%f%f%f", &magic, &x, &y, &z)

		}
	}
}
