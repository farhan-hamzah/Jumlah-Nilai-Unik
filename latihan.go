package main
import "fmt"

const NMAX int = 100

func main() {
	var A [NMAX]int
	var n, i, j, jum int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i = 0; i < n; i++ {
		var count int = 0
		for j = 0; j < n; j++ {
			if A[i] == A[j] {
				count++
			}
		}
		if count == 1 {
			jum += A[i]
		}
	}

	fmt.Print(jum)
}