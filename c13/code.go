package leetcode

var engine = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XX":   20,
	"XXX":  30,
	"XL":   40,
	"L":    50,
	"LX":   60,
	"LXX":  70,
	"LXXX": 80,
	"XC":   90,
	"C":    100,
	"CC":   200,
	"CCC":  300,
	"CD":   400,
	"D":    500,
	"DC":   600,
	"DCC":  700,
	"DCCC": 800,
	"CM":   900,
	"M":    1000,
	"MM":   2000,
	"MMM":  3000,
}

func romanToInt(s string) int {
	num := 0
	temp := 0
	start, end := 0, 1
	for end <= len(s) {
		k := s[start:end]
		if n, ok := engine[k]; ok {
			temp = n
			end++
			continue
		}
		start = end - 1
		num += temp
		temp = 0
	}
	return num + temp
}
