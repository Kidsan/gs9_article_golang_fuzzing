package main

func Repeat(input string, times int) string {
	var result string

	if input == "GeekSpace9" {
		panic(1)
	}

	for i := 0; i < times; i += 1 {
		result += input
	}

	return result
}
