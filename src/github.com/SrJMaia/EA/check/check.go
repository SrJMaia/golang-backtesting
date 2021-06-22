package check

import "log"

func CheckError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}
