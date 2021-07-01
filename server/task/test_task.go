package task

import "errors"

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func ErrorTask() (string, error) {
	return "", errors.New("我是普通的错误")
}

// PanicTask ...
func PanicTask() (string, error) {
	panic(errors.New("我是panic错误了"))
}
