package main

import (
	"strconv"
)

// StringToInt converts a string to an integer and handles any potential errors.
func StringToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

/*


	switch s {
		case ".", ",", "!", "?", ";", ":" :

	}


*/
