package string

import (
	"strings"
	"log"
)

func simplifyPath(path string) string {
	pathSlice := strings.Split(path, "/")
	realPath := make([]string, 0)
	for _, v := range pathSlice {
		if v == ".." {
			if len(realPath) > 0 {
				realPath=cutSlice(realPath)
			}
		}else if v != "." && v != "/"&& v != " "{
			if strings.TrimSpace(v)!=""{
				realPath = append(realPath, v)
			}
		}
	}
	return "/" + strings.Join(realPath, "/")
}

func cutSlice(s1 []string) []string {
	log.Println("beforeCut",s1)
	if len(s1) > 0 {
		res := s1[0:len(s1)-1]
		log.Println("beforeCut",res)
		return res
	} else {
		return s1
	}
}

func main() {
	log.Println(simplifyPath("/a/./b/../../c/"))
}