package main

import (
	"strings"
	"fmt"
)

func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}
	if (s != "" && p == "") || (s == "" && p != "") {
		return false
	}
	arrayP := analysisP(p)
	arrayS := strings.Split(s, "")
	for len(arrayS) > 0 {
		if len(arrayP[0]) == 1 {
			if arrayP[0] == "." {
				arrayP = arrayP[1:]
				arrayS = arrayS[1:]
			}
			if len(arrayS) == 0 || len(arrayP) == 0 {
				break
			}
			if arrayP[0] != arrayS[0] {
				return false
			} else {
				arrayP = arrayP[1:]
				arrayS = arrayS[1:]
			}
			if len(arrayS) == 0 || len(arrayP) == 0 {
				break
			}
		} else {
			matStr := strings.Split(arrayP[0], "")[0]
			if matStr == "." {
				for len(arrayS) > 0 && len(arrayP) > 0 {
					sOuter := ""
					pOuter := ""
					arrayP = arrayP[1:]
					if len(arrayP) == 0 {
						break
					}
					for _, v := range arrayS {
						sOuter += v
					}
					for _, v := range arrayP {
						pOuter += v
					}
					if isMatch(sOuter, pOuter) {
						return true
					}
				}
			} else {
				for arrayS[0] == matStr&&len(arrayS) > 0 && len(arrayP) > 0 {
					sOuter := ""
					pOuter := ""
					if len(arrayS) == 0 {
						break
					}
					if len(arrayP) == 0 {
						break
					}
					if len(arrayP)==1 && len(arrayP[0])==2{
					}else{
						arrayP=arrayP[1:]
					}
					for _, v := range arrayS {
						sOuter += v
					}
					for _, v := range arrayP {
						pOuter += v
					}
					if isMatch(sOuter, pOuter) {
						return true
					}
					arrayS = arrayS[1:]
					if len(arrayS)==0{
						break
					}
				}
				if len(arrayP)!=0{
					arrayP = arrayP[1:]
				}
			}
		}
		if len(arrayS) == 0 || len(arrayP) == 0 {
			break
		}
	}
	if len(arrayP) > 0 || len(arrayS) > 0 {
		return false
	} else {
		return true
	}
}
func analysisP(p string) ([]string) {
	analysisPSlice := make([]string, 0)
	analysisPSliceOuter := make([]string, 0)
	pSli := strings.Split(p, "")
	for i := 0; i < len(pSli); i++ {
		if i != len(pSli)-1 {
			if pSli[i+1] == "*" {
				analysisPSlice = append(analysisPSlice, pSli[i]+"*")
				i++
			} else {
				analysisPSlice = append(analysisPSlice, pSli[i])
			}
		} else {
			analysisPSlice = append(analysisPSlice, pSli[i])
		}
	}
	if len(analysisPSlice) <= 1 {
		return analysisPSlice
	}
	analysisPSliceOuter = append(analysisPSliceOuter, analysisPSlice[0])
	for i, v := range analysisPSlice {
		if i > 0 {
			if len(analysisPSliceOuter[len(analysisPSliceOuter)-1]) == 1 && len(v) == 1 {
				analysisPSliceOuter = append(analysisPSliceOuter, v)
			}else{
				if strings.Trim(analysisPSliceOuter[len(analysisPSliceOuter)-1],"*")!=strings.Trim(v,"*"){
					analysisPSliceOuter = append(analysisPSliceOuter, v)
				}else{
					analysisPSliceOuter[len(analysisPSliceOuter)-1] = strings.Trim(v,"*")+"*"
				}
			}
		}
	}
	return analysisPSliceOuter
}
func main() {
	fmt.Println(isMatch("aaa","a*a*"))
}
