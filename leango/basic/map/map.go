package main

//func main() {
//
//	m := map[string]string{
//		"name":    "ccmouse",
//		"course":  "golang",
//		"site":    "imooc",
//		"quality": "notbad",
//	}
//
//	m2 := make(map[string]int)
//
//	var m3 map[string]int
//
//	for k, v := range m3 {
//		fmt.Println(k, v)
//	}
//
//	for k := range m3 {
//		fmt.Println(k)
//	}
//
//	for _, v := range m3 {
//		fmt.Println(v)
//	}
//
//	courseName := m["course"]
//	fmt.Println(courseName)
//
//	if causeName, ok := m["cause"]; ok {
//		fmt.Println(causeName)
//	} else {
//		fmt.Println("key dose not exit")
//	}
//
//	delete(m, "name")
//
//}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
