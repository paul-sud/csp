/*
Problem: A square matrix A of order 3 is given. Three
streams are to be input, each stream representing a
column of an array IN. Three streams are to be output,
each representing a column of the product matrix IN ×
A. After an initial delay, the results are to be produced
at the same rate as the input is consumed. Consequently,
a high degree of parallelism is required. The solution
should take the form shown in Figure 2. Each of the nine
nonborder nodes inputs a vector component from the
west and a partial sum from the north. Each node outputs
the vector component to its east, and an updated partial
sum to the south. The input data is produced by the west
border nodes, and the desired results are consumed by
south border nodes. The north border is a constant
source of zeros and the east border is just a sink. No
provision need be made for termination nor for changing
the values of the array A.

Notes: this computes transpose(IN) * A, not the usual matrix vector product A * IN!
*/
package main

import (
	"fmt"
	"time"
)

const nDimensions int = 3

func main() {
	A := [nDimensions][nDimensions]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	x := [nDimensions]int{1, 2, 3}
	var westEast [nDimensions][nDimensions + 1]chan int
	var northSouth [nDimensions + 1][nDimensions]chan int
	for i := 0; i < nDimensions; i++ {
		for j := 0; j < (nDimensions + 1); j++ {
			westEast[i][j] = make(chan int)
		}
	}
	for i := 0; i < nDimensions+1; i++ {
		for j := 0; j < nDimensions; j++ {
			northSouth[i][j] = make(chan int)
		}
	}
	for i := 0; i < nDimensions; i++ {
		for j := 0; j < nDimensions; j++ {
			if i == 0 {
				go northBorderNode(northSouth[i][j])
			}
			if i == (nDimensions - 1) {
				go southBorderNode(northSouth[i+1][j])
			}
			if j == 0 {
				go westBorderNode(x[i], westEast[i][j])
			}
			if j == (nDimensions - 1) {
				go eastBorderNode(westEast[i][j+1])
			}
			go interiorNode(A[i][j], westEast[i][j], westEast[i][j+1], northSouth[i][j], northSouth[i+1][j])
		}
	}
	time.Sleep(100 * time.Millisecond)
}

func interiorNode(
	value int,
	vectorComponent <-chan int,
	eastNode chan<- int,
	prevSum <-chan int,
	nextSum chan<- int,
) {
	elem := <-vectorComponent
	sum := <-prevSum

	product := elem * value
	sum += product

	nextSum <- sum
	eastNode <- elem
}

func northBorderNode(initialSum chan<- int) {
	initialSum <- 0
}

func southBorderNode(result <-chan int) {
	sum := <-result
	fmt.Println("Output component is", sum)
}

func westBorderNode(inputVectorComponent int, output chan<- int) {
	output <- inputVectorComponent
}

func eastBorderNode(input <-chan int) {
	<-input
}
