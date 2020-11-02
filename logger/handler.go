package logger

import (
	"reflect"
)

func HandleErrorX(err error, action ...interface{}) {
	l := len(action)
	if l == 0 {
		HandleError(err, nil)
	}
	if l > 1 {
		Fatal("Can't Accept More than 1 Error Handler Action")
	}
	atc := reflect.TypeOf(action[0])
	if atc.String() == "func()" {
		action[0].(func())()
	} else {
		Warn("HandleErrorX Have No func() To Run...")
	}
}

func HandleDefaultErr(err error) {
	HandleError(err, nil)
}
func HandleError(err error, action func()) {
	if err != nil {
		if action != nil {
			action()
		} else {
			Error(err)
		}
	}
}
func HandleErrorF(err error, format string, v ...interface{}) {
	if err != nil {
		Errorf(format, v...)
	}
}

func HandleFatalF(err error, format string, v ...interface{}) {
	if err != nil {
		Fatalf(format, v...)
	}
}
