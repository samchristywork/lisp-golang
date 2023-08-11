package env

import (
	"lisp/model"
)

func typeof(kind int) string {
	for key, value := range model.Types {
		if key == kind {
			return value
		}
	}
	return "Type not found"
}
