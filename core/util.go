package core

import (
	"lisp/model"
)

func typeof(kind int) string {
	for key, value := range model.Types {
		if key == kind {
			return value
		}
	}
	return "TYPE NOT FOUND"
}
