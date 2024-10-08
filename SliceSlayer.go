package main

func SliceSlayer(words []string) []string {
	var edited []string

	for _, each := range words {
		if each != "" {
			edited = append(edited, each)
		}
	}
	return edited
}
