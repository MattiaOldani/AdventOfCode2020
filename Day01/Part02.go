package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main () {

	var number []int
	var results int
	number=ReadText()
	for i:=range number {
		for j:=i+1; j<len(number); j++ {
			for k:=j+1; k<len(number); k++ {
				if number[i]+number[j]+number[k]==2020 {
					results=number[i]*number[j]*number[k]
					break
				}
			}
			if results > 0 {
				break
			}
		}
		if results > 0 {
			break
		}
	}
	fmt.Println(results)

}

func ReadText () (number []int) {

  	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		N,_ := strconv.Atoi(scanner.Text())
		number=append(number, N)
	}
  	return

}
