package solveError

import "log"

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recover error: ", err)
	}
}
