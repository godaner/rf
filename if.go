package rf

import "log"

func IfErrThenExit(err error, v ...any) {
	if err != nil {
		if len(v) == 0 {
			log.Fatalln(err)
		} else {
			log.Fatalln(v...)
		}
	}
}

func IfTrueThenExit(b bool, v ...interface{}) {
	if b {
		log.Fatalln(v...)
	}
}
