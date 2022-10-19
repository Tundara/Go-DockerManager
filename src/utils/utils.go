package utils

import "os"

//CheckRoot is the function for check if user are root
func CheckRoot() bool {
	file := "/root/tmpshellcheck.oklm"
	_, err := os.Create(file)
	if err != nil {
		os.Remove(file)
		return false
	}
	os.Remove(file)
	return true
}
