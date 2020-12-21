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
                for x:=i-1; x>0; x-- {
                    if cop[x][j]=='L' {
                        break
                    }
                    if cop[x][j]=='#' {
                        occ++
                        break
                    }
                }
                for x:=i+1; x<len(cop); x++ {
                    if cop[x][j]=='L' {
                        break
                    }
                    if cop[x][j]=='#' {
                        occ++
                        break
                    }
                }
                for y:=j+1; y<len(cop[i]); y++ {
                    if cop[i][y]=='L' {
                        break
                    }
                    if cop[i][y]=='#' {
                        occ++
                        break
                    }
                }
                for y:=j-1; y>0; y-- {
                    if cop[i][y]=='L' {
                        break
                    }
                    if cop[i][y]=='#' {
                        occ++
                        break
                    }
                }
                for x,y:=i+1,j+1; x<len(cop) && y<len(cop[i]); x,y=x+1,y+1 {
                    if cop[x][y]=='L' {
                        break
                    }
                    if cop[x][y]=='#' {
                        occ++
                        break
                    }
                }
                for x,y:=i+1,j-1; x<len(cop) && y>0; x,y=x+1,y-1{
                    if cop[x][y]=='L' {
                        break
                    }
                    if cop[x][y]=='#' {
                        occ++
                        break
                    }
                }
                for x,y:=i-1,j+1; x>0 && y<len(cop[i]); x,y=x-1,y+1{
                    if cop[x][y]=='L' {
                        break
                    }
                    if cop[x][y]=='#' {
                        occ++
                        break
                    }
                }
                for x,y:=i-1,j-1; x>0 && y>0; x,y=x-1,y-1{
                    if cop[x][y]=='L' {
                        break
                    }
                    if cop[x][y]=='#' {
                        occ++
                        break
                    }
                }
                if cop[i][j]=='L' && occ==0 {
                    field[i][j]='#'
                    counter++
                }
                if cop[i][j]=='#' && occ>=5 {
                    field[i][j]='L'
                    counter++
                }
            }
        }
        if counter==0 {
            break
        }
    }
    N=Count(field)
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
