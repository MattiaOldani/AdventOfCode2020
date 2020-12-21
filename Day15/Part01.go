package main

import "fmt"
import "bufio" 
import "os"
import "strings"
import "strconv" 

func main () {
    
    var text string
    text=ReadText()
    N:=Turn(text)
    fmt.Println(N)
    
}

func ReadText () (text string) {
    
    scanner:=bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        text=scanner.Text()
    }
    return
    
}

func Turn (text string) (N int) {
    
    var turn int
    var value []int
    a:=strings.Split(text, ",")
    for _,b:=range a {
        c,_:=strconv.Atoi(b)
        value=append(value, c)
        turn++
    }
    for turn=turn+1; turn<=2020; turn++ {
        var ok bool
        lastins:=value[len(value)-1]
        fmt.Println(lastins)
        aux:=value[:len(value)-1]
        for i,v:=range aux {
            if lastins==v {
                value[i]=-1
                d:=(turn-1)-(i+1) 
                value=append(value, d)
                ok=true
                break
            }
        }
        if !ok {
            value=append(value, 0)
        }
    }
    N=value[len(value)-1]
    return
    
}
