package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main () {

    var text []int
    text=ReadText()
    N:=Bus(text, 100000000000000)
    fmt.Println(N)

}

func ReadText () (text []int) {

    scanner:=bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        _,_=strconv.Atoi(scanner.Text())
        if scanner.Scan() {
            aux:=scanner.Text()
            t:=strings.Split(aux, ",")
            for _,x:=range t {
                if x=="x" {
                    text=append(text, 0)
                } else {
                    a,_:=strconv.Atoi(x)
                    text=append(text, a)
                }
            }
        }
    }
    return

}

func Bus (bus []int, beat int) (int) {
	
	var mod int=1
	for i,x:=range bus {
		if x==0 {
			continue
		}
		for (beat+i)%x!=0 {
		    beat+=mod
		}
		mod*=x
	}
	return beat
	
}
