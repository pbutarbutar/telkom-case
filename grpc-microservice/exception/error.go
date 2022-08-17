package exception

import "log"

func Panic(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
