package main

import "fmt"
import "os"
import "bufio"
import "strconv" 

func main() {

    var assembly map[int]string
    var line int
    assembly,line=ReadMap()
    accumulator:=Substitution(assembly, line)
    fmt.Println(accumulator)
    
}

func ReadMap () (assembly map[int]string, line int) {
    
    assembly=make(map[int]string)
    scanner:=bufio.NewScanner(os.Stdin)
    for i:=0; scanner.Scan(); i++ {
        assembly[i]=scanner.Text()
        line++
    }
    return 
    
}

func Substitution (assembly map[int]string, line int) (accumulator int) {

    for i:=0; i<line; i++ {
        if (assembly[i])[:3]=="jmp" {
            var ok bool
            assembly[i]="nop"+(assembly[i])[3:]
            ok,accumulator=Loop(assembly, line)
            if !ok {
                return
            } else {
                assembly[i]="jmp"+(assembly[i])[3:]
            }
        } else if (assembly[i])[:3]=="nop" {
            var ok bool
            assembly[i]="jmp"+(assembly[i])[3:]
            ok,accumulator=Loop(assembly, line)
            if !ok {
                return
            } else {
                assembly[i]="nop"+(assembly[i])[3:]
            }
        }
    }
    return

}

func Loop (assembly map[int]string, line int) (bool, int) {
    
    var accumulator int
    var linevisited []string
    for i:=0; i<line; {
        var ok bool
        linevisited,ok=Visited(i, linevisited)
        if ok {
            return true, 0
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
    return false, accumulator
    
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
