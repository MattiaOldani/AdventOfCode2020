package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func main () {

	var a,b []int
	var char []rune
	var password, validd []string
	var N int
	a,b,char,password=ReadPassword()
  	valid,N=Password(a,b,char,password)
  	fmt.Println("Il numero di password valide é", N)

}

func ReadPassword () (a []int, b []int, char []rune, password []string) {

	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line string=scanner.Text()
		var n1,n2 int
		n1,line=Extreme(line)
		n2,line=Extreme(line)
		a=append(a, n1)
		b=append(b, n2)
		var c rune
		c,line=Char(line)
		char=append(char, c)
		password=append(password, line)
	}
	return

}

func Extreme (line string) (ext int, rest string) {

	counter:=0
	for _,c:=range line {
		if counter==0 {
			counter++
		} else {
			if c==' ' || c=='-' {
				break
			} else {
				counter++
				break
			}
		}
	}
	ext,_=strconv.Atoi(line[0:counter])
	rest=line[counter+1:]
	return

}

func Char (line string) (char rune, rest string) {
	
	char=rune(line[0])
	rest=line[3:]
	return
	
}

func Password (a []int, b []int, char []rune, password []string) (valid []string, N int) {

    	for i,psw:=range password {
        	c:=0
        	for _,letter:=range psw {
            		if rune(letter)==rune(char[i]) {
                		c++
           		}
        	}
        	if c>=a[i] && c<=b[i] {
            		valid=append(valid, psw)
            		N++
        	}
    	}
    	return
    
}
