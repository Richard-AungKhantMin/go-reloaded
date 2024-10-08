package main

func NumMod(num int, index int, words []string, f func(string) string) {

	for starter := index - num; starter < index; starter++ {
		words[starter] = f(words[starter])
	}
}
