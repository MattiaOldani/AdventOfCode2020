package main

import "fmt"
import "os"
import "bufio"
import "strings" 

func main() {

    	sum:=ReadText()
    	fmt.Println("La somma delle soluzioni é", sum)
    
}

func ReadText () (N int) {
    
    	var text string
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	if scanner.Text()=="" {
            		N+=Answer(text)
            		text=""
        	} else {
            		if text=="" {
                		text+=scanner.Text()
            		} else {
                		text+=" "+scanner.Text()
            		}
        	}
    	}
    	return
    
}

func Answer (text string) (N int) {
    
    	var ans []string
    	var letter []rune
    	ans=strings.Split(text, " ")
    	for _,x:=range ans {
        	for _,char:=range x {
            		var ok bool = true
            		for _,ris:=range letter {
                		if char==ris {
                    			ok=false
                		}
            		}
            		if ok {
               			letter=append(letter, char)
               			N++
            		}
        	}
    	}
    	return
    
}
