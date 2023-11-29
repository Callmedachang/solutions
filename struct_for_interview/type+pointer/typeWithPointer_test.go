package type_pointer

import (
	"log"
	"testing"
)

func TestPersonal(t *testing.T) {
	p := personal{
		Name: "lwc",
	}
	pp := &p
	p.UpdateNameP("luzi")
	log.Println(p.Name)
	pp.UpdateNameP("luzi1")
	log.Println(p.Name)
}
