package main

import (
	"log"
)

var (
	words    = []string{"One ", "Two ", "Three ", "Four ", "Five ", "Six ", "Seven ", "Eight ", "Nine "}
	tinWords = []string{"Ten ", "Eleven ", "Twelve ", "Thirteen ", "Fourteen ", "Fifteen ", "Sixteen ", "Seventeen ", "Eighteen ", "Nineteen "}
	teyWords = []string{"Twenty ", "Thirty ", "Forty ", "Fifty ", "Sixty ", "Seventy ", "Eighty ", "Ninety "}
)

func numberToWords(num int) string {
	if num == 0 {return "zero"}
	res := ""
	if num >= 1000000000 {
		houMuchBillion := num / 1000000000
		num = num % 1000000000
		res = res + words[houMuchBillion-1] + "Billion "
	}
	if num >= 1000000 {
		houMuchMillion := num / 1000000
		num = num % 1000000
		res = res + nameForNumberSmallThanThousand(houMuchMillion) + "Million "
	}
	if num >= 1000 {
		houMuchThousand := num / 1000
		num = num % 1000
		res = res + nameForNumberSmallThanThousand(houMuchThousand) + "Thousand "
	}
	if num > 0 && num < 1000 {
		res = res + nameForNumberSmallThanThousand(num)
	}
	return string([]byte(res)[:len(res)-1])
}
func nameForNumberSmallThanThousand(num int) string {
	res := ""
	houMuchHundred := num / 100
	num = num % 100
	houMuchTen := num / 10
	hoMuch := num % 10
	if houMuchHundred > 0 {
		res = res + words[houMuchHundred-1] + "Hundred "
	}
	if houMuchTen > 0 {
		if houMuchTen == 1 {
			res = res + tinWords[hoMuch]
			return res
		} else {
			res = res + teyWords[houMuchTen-2]
		}
	}
	if hoMuch > 0 {
		res = res + words[hoMuch-1]
	}
	return res
}

func main() {
	log.Println(numberToWords(12345))
}
