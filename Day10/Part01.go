package main

import "fmt"
import "bufio"
import "os"
import "sort"
import "strconv"

func main () {

    var text []int
    text=ReadText()
    N:=Jolts(text)
    fmt.Println("La risposta é", N)

}

func ReadText () (text []int) {

    scanner:=bufio.NewScanner(os.Stdin)
    text=append(text, 0)
    for scanner.Scan() {
        x,_:=strconv.Atoi(scanner.Text())
        text=append(text, x)
    }
    return

}

func Jolts (text []int) (N int) {

    var dif map[int]int
    dif=make(map[int]int)
    sort.Ints(text)
    text=append(text, text[len(text)-1]+3)
    fmt.Println(text)
    for i:=0; i<len(text)-1; i++{
        x:=text[i+1]-text[i]
        fmt.Println(x)
        switch x {
            case 1: dif[1]++
            case 3: dif[3]++
        }
        fmt.Println(dif)
    }
    N=dif[1]*dif[3]
    return

}
