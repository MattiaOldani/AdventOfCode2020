package main

import "fmt"
import "os"
import "bufio"
import "strconv" 

func main() {

    var assembly map[int]string
    assembly=ReadMap()
    accumulator:=Loop(assembly)
    fmt.Println(accumulator)
    
}

func ReadMap () (assembly map[int]string) {
    
    assembly=make(map[int]string)
    scanner:=bufio.NewScanner(os.Stdin)
    for i:=0; scanner.Scan(); i++ {
        assembly[i]=scanner.Text()
    }
    return 
    
}

func Loop (assembly map[int]string) (accumulator int) {
    
    var linevisited []string
    for i:=0; ; {
        var ok bool
        linevisited,ok=Visited(i, linevisited)
        if ok {
            break
        }
        instruction:=assembly[i]
        switch instruction[:3] {
            case "acc":
                sign:=string(instruction[4])
                number:=string(instruction[5:]) 
                if sign=="+" {
                    x,_:=strconv.Atoi(number)
                    accumulator+=x
                } else {
                    x,_:=strconv.Atoi(number)
                    accumulator-=x
                }
                i++
            case "jmp":
                sign:=string(instruction[4]) 
                number:=string(instruction[5:]) 
                if sign=="+" {
                    x,_:=strconv.Atoi(number)
                    i+=x
                } else {
                    x,_:=strconv.Atoi(number)
                    i-=x
                }
            default: i++
        }
    }
    return
    
}

func Visited (key int, line []string) (total []string, ok bool) {
    
    for _,x:=range line {
        if string(key)==x {
            ok=true
            return
        }
    }
    total=append(line, string(key))
    return
    
}
