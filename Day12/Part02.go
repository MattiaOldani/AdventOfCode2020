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
    
    var x, y int=0, 0
    var xwp, ywp int=10, 1
    for _,line:=range text {
        k:=string(line[0])
        mv,_:=strconv.Atoi(line[1:])
        switch k {
            case "R":
                if mv==90 {
                    xwp,ywp=ywp,-xwp
                } else if mv==180 {
                    xwp,ywp=-xwp,-ywp
                } else if mv==270 {
                    xwp,ywp=-ywp,xwp
                }
            case "L":
                if mv==90 {
                    xwp,ywp=-ywp,xwp
                } else if mv==180 {
                    xwp,ywp=-xwp,-ywp
                } else if mv==270 {
                    xwp,ywp=ywp,-xwp
                }
            case "N": ywp+=mv
            case "S": ywp-=mv
            case "E": xwp+=mv
            case "W": xwp-=mv
            case "F": x,y=x+mv*xwp,y+mv*ywp
        }
    }
    N=int(math.Abs(float64(x))+math.Abs(float64(y)))
    return
    
}
