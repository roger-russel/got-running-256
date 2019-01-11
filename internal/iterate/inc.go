package iterate

import "fmt"

// Inc iterate over all numbers on 256
func Inc(spliter int, expo int) {

	switch true {
	case expo > 1:
		for count := 0; count < 255; count++ {
			Inc(spliter, expo-1)
		}
		break
	case expo == 1:
		ZeroInc(spliter, expo)
		break
	default:
		fmt.Println(expo, spliter, spliter)
	}

}

// ZeroInc increment when expo its 1
func ZeroInc(spliter int, expo int) {
	for count := 0; count < 255; count++ {
		fmt.Println(expo, spliter, count)
	}
}
