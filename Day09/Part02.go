package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "sort"

func main () {

    var text []int
    text=ReadText()
    N:=Sum(text)
    x:=Piece(text, N)
    fmt.Println(x)

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

func Piece (text []int, x int) (N int) {

    for i:=0; i<len(text); i++ {
        min,max,ok:=MinMax(text[i:], x)
        if ok {
            N=min+max
            break
        }
    }
    return

}

func MinMax (text []int, target int) (int, int, bool) {

    var counter, max int
    for i:=0; i<len(text); i++ {
        counter+=text[i]
        if counter==target && i>1 {
            max=i
            break
        }
        if counter>target {
            return 0, 0, false
        }
    }
    text=text[:max+1]
    sort.Ints(text)
    return text[0], text[max], true

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
