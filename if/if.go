package _if

import "log"

func ErrThenExit(err error, v ...any) {
	if err != nil {
		if len(v) == 0 {
			log.Fatalln(err)
		} else {
			log.Fatalln(v...)
		}
	}
}

func TrueThenExit(b bool, v ...interface{}) {
	if b {
		log.Fatalln(v...)
	}
}
