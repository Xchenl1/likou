package main

import "fmt"

func reverseStr(s string, k int) string {
	end := ""
	for len(s) >= 2*k {
		k1 := s[:2*k]
		k11 := []byte(k1)
		for i := 0; i < len(k11)/4; i++ {
			k11[i], k11[len(k1)/2-1-i] = k11[len(k1)/2-1-i], k11[i]
		}
		s = s[2*k:]
		end += string(k11)
	}
	if len(s) < 2*k && len(s) >= k {
		k11 := []byte(s)
		for i := 0; i < k/2; i++ {
			k11[i], k11[k-1-i] = k11[k-1-i], k11[i]
		}
		end += string(k11)
	} else if len(s) < k {
		k11 := []byte(s)
		for i := 0; i < len(s)/2; i++ {
			k11[i], k11[len(s)-1-i] = k11[len(s)-1-i], k11[i]
		}
		end += string(k11)
	}
	return end
}
func main() {
	s := "fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi"
	k := 39
	fmt.Printf(reverseStr(s, k))
}

//"fdcqkmxwholhytmhafpesaentdvxginrjlyqzyhehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqllgsqddebemjanqcqnfkjmi"
