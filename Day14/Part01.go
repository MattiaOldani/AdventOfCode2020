package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main () {

    N:=ReadText()
    fmt.Println(N)

}

func ReadText () (N int) {

    var text []string
    var x string
    var i int
    var memory map[int]int
    memory=make(map[int]int)
    scanner:=bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        x=scanner.Text()
        if x[:4]=="mask" && i==0 {
            text=append(text, x)
            i=1
            continue
        }
        if x[:4]=="mask" {
            Mask(text, &memory)
            text=make([]string, 0)
        }
        text=append(text, x)
    }
    Mask(text, &memory)
    for _,val:=range memory {
        N+=val
    }
    return
    
}
                
func Mask (text []string, memory *map[int]int) () {

    x:=text[0]
    y:=strings.Split(x, " ")
    mask:=y[2]
    text=text[1:]
    for _,line:=range text {
        m,val:=Parse(line)
        val=ApplyMask(val, mask)
        (*memory)[m]=val
    }
    return

}

func ApplyMask (val int, mask string) (N int) {

    var bin, app string
    t:=Binary(uint(val))
    for _,c:=range t {
        bin+=string(c)
    }
    for i,x:=range mask {
        if string(x)=="1" {
            app+="1"
        } else if string(x)=="0" {
            app+="0"
        } else {
            app+=string(bin[i])
        }
    }
    N=int(Decimal(app))
    return

}

func Binary (val uint) ([]rune) {

    return []rune(fmt.Sprintf("%036b", val))
    
}

func Decimal (val string) (uint) {

    x, _:=strconv.ParseInt(val, 2, 37)
    return uint(x)

}

func Parse (line string) (m int, val int) {

    x:=strings.Split(line, "[")
    y:=x[1]
    z:=strings.Split(y, "]")
    m,_=strconv.Atoi(z[0])
    t:=z[1]
    k:=strings.Split(t, " ")
    val,_=strconv.Atoi(k[2])
    return

}
