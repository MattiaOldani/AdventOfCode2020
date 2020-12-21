package main 

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "unicode"

func main () {

    	N:=ReadText()
    	fmt.Println("Ho", N, "password valide")

}

func ReadText () (N int) {

    	var text string
    	scanner:=bufio.NewScanner(os.Stdin)
    	for scanner.Scan() {
        	if scanner.Text()=="" {
            		N+=Passport(text)
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

func Passport (text string) (int) {

    	var c int
    	var sep []string
    	sep=strings.Split(text, " ")
    	for _,field:=range sep {
        	switch {
            		case strings.Contains(field, "byr:"): c+=Valbyr(field)
            		case strings.Contains(field, "iyr:"): c+=Valiyr(field)
            		case strings.Contains(field, "eyr:"): c+=Valeyr(field)
            		case strings.Contains(field, "hgt:"): c+=Valhgt(field)
            		case strings.Contains(field, "hcl:"): c+=Valhcl(field)
            		case strings.Contains(field, "ecl:"): c+=Valecl(field)
            		case strings.Contains(field, "pid:"): c+=Valpid(field)
        	}
    	}
    	if c==7 {
        	return 1
   	}
    	return 0

}

func Valbyr (field string) (N int) {

    	ext:=field[4:]
    	year,_:=strconv.Atoi(ext)
    	if year>=1920 && year<=2002 {
        	N++
    	}
    	return
   
}

func Valiyr (field string) (N int) {

    	ext:=field[4:]
    	year,_:=strconv.Atoi(ext)
    	if year>=2010 && year<=2020 {
        	N++
    	}
    	return
   
}

func Valeyr (field string) (N int) {

    	ext:=field[4:]
    	year,_:=strconv.Atoi(ext)
    	if year>=2020 && year<=2030 {
        	N++
    	}
    	return
   
}

func Valhgt (field string) (N int) {

    	ext:=field[4:]
    	var aux string
    	var counter int
    	for _,c:=range ext {
        	if unicode.IsDigit(c) {
            		aux+=string(c)
            		counter++
        	} else {
            		break
        	}
    	}
        height,_:=strconv.Atoi(aux)
    	if strings.Contains(ext, "cm") {
        	if height>=150 && height<=193 {
            		N++
        	}
    	} else if strings.Contains(ext, "in") {
        	if height>=59 && height<=76 {
            		N++
        	}
    	}
    	return

}

func Valhcl (field string) (N int) {

    	ext:=field[5:]
    	var counter int
    	for _,c:=range ext {
        	if (c>='0' && c<='9') || (c>='a' && c<='f') {
            		counter++
        	}
    	}
    	if counter==6 {
        	N++
    	}
    	return

}

func Valecl (field string) (N int) {

    	ext:=field[4:]
    	switch ext {
        	case "amb": N++
        	case "blu": N++
        	case "brn": N++
        	case "gry": N++
        	case "grn": N++
        	case "hzl": N++
        	case "oth": N++
    	}
    	return

}

func Valpid (field string) (N int) {

    	ext:=field[4:]
    	var counter int
    	for _,c:=range ext {
        	if unicode.IsDigit(c) {
            		counter++
        	}
    	}
    	if counter==9 {
        	N++
    	}
    	return

}
