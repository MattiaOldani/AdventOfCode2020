package main

import "fmt"
import "bufio"
import "os"

func main () {
    
    var text [][]rune
    text=ReadText()
    N:=Seat(text)
    fmt.Println(N)
    
}

func ReadText () (text [][]rune) {
    
    scanner:=bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        x:=[]rune(scanner.Text())
        text=append(text, x)
    }
    return
    
}

func Seat (field [][]rune) (N int) {
    
    var cop [][]rune
    var occ, counter int
    for {
        cop=Copy(&field)
        counter=0
        for i:=0; i<len(cop); i++ {
            for j:=0; j<len(cop[i]); j++ {
                occ=0
                if i!=0 && cop[i-1][j]=='#' {
                    occ++
                }
                if i!=len(cop)-1 && cop[i+1][j]=='#' {
                    occ++
                }
                if j!=0 && cop[i][j-1]=='#'{
                    occ++
                }
                if j!=len(cop[i])-1 && cop[i][j+1]=='#' {
                    occ++
                }
                if i!=0 && j!=0 && cop[i-1][j-1]=='#' {
                    occ++
                }
                if i!=0 && j!=len(cop[i])-1 && cop[i-1][j+1]=='#' {
                    occ++
                }
                if i!=len(cop)-1 && j!=0 && cop[i+1][j-1]=='#' {
                    occ++
                }
                if i!=len(cop)-1 && j!=len(cop[i])-1 && cop[i+1][j+1]=='#' {
                    occ++
                }
                if cop[i][j]=='L' && occ==0 {
                    field[i][j]='#'
                    counter++
                }
                if cop[i][j]=='#' && occ>=4 {
                    field[i][j]='L'
                    counter++
                }
            }
        }
        if counter==0 {
            break
        }
    }
    N=Count(cop)
    return
    
}

func Count (field [][]rune) (N int) {

    for i,_:=range field {
        for j,_:=range field[i] {
            if field[i][j]=='#' {
                N++
            }
        }
    }
    return

}

func Copy (field *[][]rune) (cop [][]rune) {
	
	for i:=0; i<len(*field); i++ {
		cop=append(cop, make([]rune, len((*field)[i])))
		for k:=0; k<len((*field)[i]); k++ {
			cop[i][k]=(*field)[i][k]
		}
	}
	return
	
}
