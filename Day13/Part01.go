package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main () {

    var text []int
    var timestamp int
    text,timestamp=ReadText()
    N:=Bus(text, timestamp)
    fmt.Println(N)

}

func ReadText () (text []int, timestamp int) {

    scanner:=bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        timestamp,_=strconv.Atoi(scanner.Text())
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

func Bus (bus []int, timestamp int) (N int) {
	
	var search map[int]int
	search=make(map[int]int, len(bus))
	for _,x:=range bus {
		if x==0 {
			continue
		}
		for t:=0; x*t<timestamp+x; t++ {
			search[x]=x*t
		}
	}
	min:=search[bus[0]]
	for i,x:=range search {
		if x<min {
			min=search[i]
			N=i
		}
	}
	N=(min-timestamp)*N
	return
	
}
