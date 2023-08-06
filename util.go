package main

func typeof(kind int) string {
	for key, value := range types {
		if key == kind {
			return value
		}
	}
	return "TYPE NOT FOUND"
}
