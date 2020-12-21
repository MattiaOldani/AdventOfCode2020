package main

import "fmt"
import "os"
import "bufio"

func main() {

    	var text []string
    	var position [][]int
    	position=make([][]int, 128)
    	for i:=range position {
        	position[i]=make([]int, 8)
    	}
    	text=ReadText()
    	Seat(text, position)
    	id:=SeatSearch(position)
    	fmt.Println("Il mio posto ha ID", id)
    
}

func ReadText () (text []string) {
    
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	text=append(text, scanner.Text())
    	}
    	return
    
}

func Seat (text []string, seat [][]int) () {
    
    	for _,position:=range text {
        	var rmax, rmin, cmax, cmin int = 127, 0, 7, 0
        	for i:=0; i<7; i++ {
            		if rune(position[i])=='F' {
                		rmax=(rmax+rmin)/2
            		} else {
                		rmin=(rmax+rmin)/2+1
            		}
        	}
        	for i:=7; i<len(position); i++ {
           		if rune(position[i])=='L' {
                		cmax=(cmax+cmin)/2
            		} else {
                		cmin=(cmax+cmin)/2+1
            		}
        	}
        	seat[rmax][cmax]=1
    	}
    	for i:=range seat {
        	for j:=range seat[i] {
            		fmt.Printf("%d", seat[i][j])
        	}
        	fmt.Println()
    	}
    	return
    
}

func SeatSearch (position [][]int) (id int) {
    
    	var state bool = false
    	for i:=range position {
        	for j:=range position[i] {
            		if state==false {
                		if position[i][j]==1 {
                    			state=true
                		}
            		} else {
                		if position[i][j]==0 {
                    			id=i*8+j
                    			return
                		}
            		}
        	}
        }
    	return
    
}
