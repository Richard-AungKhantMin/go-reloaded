package functions

import "log"

func IsErrNil(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
