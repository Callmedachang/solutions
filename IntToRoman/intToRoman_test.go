package IntToRoman

import "testing"

func TestIntToRoman(t *testing.T) {
	if IntToRoman(3)=="III"{
		t.Log("SUCCESS1")
	}else{
		t.Fail()
	}
	if IntToRoman(4)=="IV"{
		t.Log("SUCCESS2")
	}else{
		t.Fail()
	}
	if IntToRoman(9)=="IX"{
		t.Log("SUCCESS3")
	}else{
		t.Fail()
	}
	if IntToRoman(58)=="LVIII"{
		t.Log("SUCCESS4")
	}else{
		t.Fail()
	}
	if IntToRoman(1994)=="MCMXCIV"{
		t.Log("SUCCESS5")
	}else{
		t.Fail()
	}
}