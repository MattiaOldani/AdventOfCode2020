package main

import "fmt"
import "os"
import "bufio"

func main() {

    	var text []string
    	var id []int
    	text=ReadText()
    	id=CheckID(text)
    	fmt.Println("Il massimo ID é", Max(id))
    
}

func ReadText () (text []string) {
    
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	text=append(text, scanner.Text())
    	}
    	return
    
}

func CheckID (text []string) (id []int) {
    
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
        	id=append(id, rmax*8+cmax)
    	}
   	  return
    
}

func Max (id []int) (max int) {
    
    	for i,n:=range id {
        	if i==0 {
            		max=n
        	} else {
            		if n>max {
                		max=n
            		}
        	}
    	}
    	return
    
}
