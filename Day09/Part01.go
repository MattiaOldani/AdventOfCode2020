package main

import "fmt"
import "bufio"
import "os"
import "strconv"

func main () {

    var text []int
    text=ReadText()
    N:=Sum(text)
    fmt.Println(N)

}

func ReadText () (text []int) {

    scanner:=bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        x,_:=strconv.Atoi(scanner.Text())
        text=append(text, x)
    }
    return

}

func Sum (text []int) (N int) {

    for i:=25; i<len(text); i++ {
        x:=Check(text[i-25:i], text[i])
        if !x {
            N=text[i]
            return
        }
    }
    return

}

func Check (first []int, ctrl int) (ok bool) {
	
	for i:=0; i<len(first); i++ {
		for j:=0; j<len(first); j++ {
			if first[i]+first[j]==ctrl {
				ok=true
				return
			}
		}
	}
	return
	
}
