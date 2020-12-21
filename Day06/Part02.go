package main

import "fmt"
import "os"
import "bufio"
import "strings" 
import "sort"

func main() {

    	sum:=ReadText()
    	fmt.Println("La somma delle soluzioni é", sum)
    
}

func ReadText () (N int) {
    
    	var text string
    	var counter int
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	if scanner.Text()=="" {
            		N+=Answer(text, counter)
            		text=""
            		counter=0
        	} else {
            		if text=="" {
                		text+=scanner.Text()
            		} else {
                		text+=" "+scanner.Text()
            		}
            		counter++
        	}
    	}
    	return
    
}

func Answer (text string, counter int) (N int) {
    
    	var ans []string
    	var letter []string
    	ans=strings.Split(text, " ")
    	for _,x:=range ans {
        	for _,char:=range x {
            		letter=append(letter, string(char))
        	}
    	}
    	N=Count(letter, counter)
    	return
    
}

func Count (text []string, counter int) (N int) {

    	var nel int = 1
    	sort.Strings(text)
    	if counter==1 {
        	N=len(text)
        	return
    	}
    	for i:=0; i<len(text)-1; i++ {
        	if text[i]==text[i+1] {
            		nel++
        	} else {
            		text=text[i+1:]
            		i=-1
            		if nel==counter {
                		N++
            		}
            		nel=1
        	}
    	}
    	if nel==counter {
        	N++
    	}
    	return

}
