package main

import "fmt"
import "os"
import "bufio"

func main () {

    	var forest []string
    	var N int
    	forest=ReadForest()
    	N=Walk(forest)
    	fmt.Println("Hai camminato sopra", N, "alberi")

}

func ReadForest () ([]string) {

    	var forest []string
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	forest=append(forest, scanner.Text())
        }
    	return forest

}

func Walk (forest []string) (N int) {

    	var j int
	var colums int=len(forest[0])
    	for i:=range forest {
        	j+=3
        	current:=forest[i+1]
        	if rune(current[j%colums])=='#' {
            		N++
        	}
    	}
    	return
   
}
