package main 

import "fmt"
import "os"
import "bufio"
import "strings"

func main () {

    	N:=ReadText()
    	fmt.Println("Ho", N, "password valide")

}

func ReadText () (N int) {

    	var text string
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	if scanner.Text()=="" {
            		N+=Fields(text)
            		text=""
        	} else {
            		if text!="" {
                		text+=" "+scanner.Text()
            		} else {
                		text+=scanner.Text()
            		}
        	}
    	}
    	return

}

func Fields (text string) (int) {

    	var c int
    	var sep []string
    	sep=strings.Split(text, " ")
    	fmt.Println(sep)
    	for _,field:=range sep {
        	switch {
            		case strings.Contains(field, "byr:"): c++
            		case strings.Contains(field, "iyr:"): c++
            		case strings.Contains(field, "eyr:"): c++
            		case strings.Contains(field, "hgt:"): c++
            		case strings.Contains(field, "hcl:"): c++
            		case strings.Contains(field, "ecl:"): c++
            		case strings.Contains(field, "pid:"): c++
        	}
    	}
    	if c==7 {
        	return 1
    	}
    	return 0

}
