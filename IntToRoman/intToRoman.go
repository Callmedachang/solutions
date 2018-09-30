package IntToRoman

func IntToRoman(num int) string {
	romanNumber := ""
	//switch num {
	numsOfM := num / 1000
	num = num - 1000*numsOfM

	numsOfCM := num / 900
	num = num - 900*numsOfCM

	numsOfD := num / 500
	num = num - 500*numsOfD

	numsOfCD := num / 400
	num = num - 400*numsOfCD

	numsOfC := num / 100
	num = num - 100*numsOfC

	numsOfXC := num / 90
	num = num - 90*numsOfXC

	numsOfL := num / 50
	num = num - 50*numsOfL

	numsOfXL := num / 40
	num = num - 40*numsOfXL

	numsOfX := num / 10
	num = num - 10*numsOfX

	numsOfIX := num / 9
	num = num - 9*numsOfIX

	numsOfV := num / 5
	num = num - 5*numsOfV

	numsOfIV := num / 4
	num = num - 4*numsOfIV

	numsOfI := num

	for i := 0; i < numsOfM; i++ {
		romanNumber = romanNumber + "M"
	}
	for i := 0; i < numsOfCM; i++ {
		romanNumber = romanNumber + "CM"
	}
	for i := 0; i < numsOfD; i++ {
		romanNumber = romanNumber + "D"
	}
	for i := 0; i < numsOfCD; i++ {
		romanNumber = romanNumber + "CD"
	}
	for i := 0; i < numsOfC; i++ {
		romanNumber = romanNumber + "C"
	}
	for i := 0; i < numsOfXC; i++ {
		romanNumber = romanNumber + "XC"
	}
	for i := 0; i < numsOfL; i++ {
		romanNumber = romanNumber + "L"
	}
	for i := 0; i < numsOfXL; i++ {
		romanNumber = romanNumber + "XL"
	}
	for i := 0; i < numsOfX; i++ {
		romanNumber = romanNumber + "X"
	}
	for i := 0; i < numsOfIX; i++ {
		romanNumber = romanNumber + "IX"
	}
	for i := 0; i < numsOfV; i++ {
		romanNumber = romanNumber + "V"
	}
	for i := 0; i < numsOfIV; i++ {
		romanNumber = romanNumber + "IV"
	}
	for i := 0; i < numsOfI; i++ {
		romanNumber = romanNumber + "I"
	}
	return romanNumber
}
