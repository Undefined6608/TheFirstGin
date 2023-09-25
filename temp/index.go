package temp

import "fmt"

var tempCode = make(map[string]string)

func SetTempCode(email string, code string) {
	fmt.Println(email, code)
	tempCode[email] = code
}

func GetTempCode(email string) string {
	return tempCode[email]
}

func DeleteTempCode(email string) {
	delete(tempCode, email)
}
