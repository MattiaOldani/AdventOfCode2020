package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

type Bag struct {
    color string
    content []string
    quantity []int
}

var colorsVisited []string

func main () {

    var bags []Bag=ReadBag()
    fmt.Println(CountBags(bags, "shiny gold"))

}

func ReadBag () (bags []Bag) {

    scanner:=bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        var b Bag=ParseBag(scanner.Text())
        bags=append(bags, b)
    }
    return

}

func ParseBag (line string) (b Bag) {

    b.color=line[:ExtractColor(line)]
    line=line[ExtractColor(line)+14:]
    var containedBags []string=strings.Split(line, ", ")
    if line[:2]=="no" {
        return
    } 
    for _,x:=range containedBags {
        val,_:=strconv.Atoi(string(x[0]))
        b.quantity=append(b.quantity, val)
        x=x[2:]
        x=x[:ExtractColor(x)]
        b.content=append(b.content, x)
    }
    return

}

func ExtractColor (text string) int {

    var spaces int
    for i:=0; i<len(text); i++ {
        if text[i]==' ' {
            spaces++
        }
        if spaces==2 {
            return i
        }
    }
    return 0

}

func CountBags (bags []Bag, color string) (counter int) {

    for _,x:=range bags {
        if x.color==color {
            for i,c:=range x.content {
                counter+=x.quantity[i]*(CountBags(bags, c)+1)
            }
        }
    }
    return

}

func Contained (b Bag, color string) (bool) {

    for _,c:=range b.content {
        if c == color {
            return true
        }
    }
    return false

}

func Visited (b Bag) (bool) {

    for _,c:=range colorsVisited {
        if c==b.color {
            return true
        }
    }
    return false

}
