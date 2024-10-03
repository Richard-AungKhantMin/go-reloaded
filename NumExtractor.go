package main

func NumExtractor(s string) int {
	for _, each := range s {
		if each >= '0' && each <= '9' {
			return int(each - '0')
		}
	}
	return 0
}
