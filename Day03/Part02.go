package main

import "fmt"
import "os"
import "bufio"

func main () {

    	var forest []string
    	var N, M, K, G, S int
    	forest=ReadForest()
    	N=Walk(forest, 1, 1)
    	M=Walk(forest, 3, 1)
    	K=Walk(forest, 5, 1)
    	G=Walk(forest, 7, 1)
    	S=Walk(forest, 1, 2)
    	fmt.Println("Hai camminato sopra", N*M*K*G*S, "alberi")

}

func ReadForest (N int) ([]string) {

    	var forest []string
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
            forest=append(forest, scanner.Text())
    	}
    	return forest

}

func Walk (forest []string, dx int, under int) (N int) {
    
    	var i, j int=0, 0
        var colums int=len(forest[0])
        for _,line:=range forest {
            j+=dx
            current:=forest[i+under]
            i+=under
            if rune(current[j%colums])=='#' {
                N++
            }
    	}
    	return
    
}
