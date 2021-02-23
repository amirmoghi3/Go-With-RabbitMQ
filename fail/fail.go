package fail

import "log"

//ShowError Fail
func ShowError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

//FailOnError Fail
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
