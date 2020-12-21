package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "math" 

func main () {
    
    var text []string
    text=ReadText()
    N:=Distance(text)
    fmt.Println(N)
    
}

func ReadText () (text []string) {
    
    scanner:=bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        text=append(text, scanner.Text())
    }
    return
    
}

func Distance (text []string) (N int) {
    
    var face string="E"
    point:=[4]string{"E", "S", "W", "N"}
    var x, y, mod, i int=0, 0, 4, 0
    for _,line:=range text {
        k:=string(line[0])
        mv,_:=strconv.Atoi(line[1:])
        switch k {
            case "R":
                if mv==90 {
                    face=point[(i+1)%mod]
                    i++
                } else if mv==180 {
                    face=point[(i+2)%mod]
                    i+=2
                } else if mv==270 {
                    face=point[(i+3)%mod]
                    i+=3
                }
            case "L":
                if mv==90 {
                    face=point[(i+3)%mod]
                    i+=3
                } else if mv==180 {
                    face=point[(i+2)%mod]
                    i+=2
                } else if mv==270 {
                    face=point[(i+1)%mod]
                    i++
                }
            case "N": x+=mv
            case "S": x-=mv
            case "E": y+=mv
            case "W": y-=mv
            case "F":
                switch face {
                    case "N": x+=mv
                    case "S": x-=mv
                    case "E": y+=mv
                    case "W": y-=mv
                }
        }
    }
    N=int(math.Abs(float64(x))+math.Abs(float64(y)))
    return
    
}
