package MaxArea

import "testing"

func TestMaxArea(t *testing.T) {
	para:=[]int{1,8,6,2,5,4,8,3,7}
	if MaxArea(para)==49{
		t.Log("SUCCESS")
	}else{
		t.Fail()
	}
	para=[]int{}
	if MaxArea(para)==0{
		t.Log("SUCCESS")
	}else{
		t.Fail()
	}
}