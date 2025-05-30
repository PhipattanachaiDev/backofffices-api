package cryptography

import (
	"strings"
)

// Cryptography service structure
type Cryptography struct {
	RefArray []string
}

var RefArray = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Encryption encrypts the input text using a simple character-based encryption algorithm
func Encryption(textLine string, txtPW string) string {
	var (
		rtnStr    string
		iTemp     int
		lenPw     int
		iPw       int = 1
		iResult1  int
		iResult2  int
		iResult   int
		doEncrypt bool
	)

	txtPW = strings.ToUpper(txtPW)
	textLine = strings.ToUpper(textLine)
	lenPw = len(txtPW)

	for i := 0; i < len(textLine); i++ {
		iTemp = int(textLine[i])

		iPwTemp := int(txtPW[(iPw - 1)])
		doEncrypt = true

		if iTemp >= 'A' && iTemp <= 'Z' {
			iResult1 = iTemp - 'A' + 1
		} else if iTemp >= '0' && iTemp <= '9' {
			iResult1 = iTemp - '0' + 27
		} else {
			doEncrypt = false
		}

		if doEncrypt {
			if iPwTemp >= 'A' && iPwTemp <= 'Z' {
				iResult2 = iPwTemp - 'A' + 1
			} else if iPwTemp >= '0' && iPwTemp <= '9' {
				iResult2 = iPwTemp - '0' + 27
			}

			iResult = iResult1 + iResult2
			if iResult > 36 {
				iResult -= 36
			}

			rtnStr += string(RefArray[iResult-1])
		} else {
			rtnStr += string(rune(iTemp))
		}

		iPw++
		if iPw > lenPw {
			iPw = 1
		}
	}

	return rtnStr
}

// Decryption decrypts the input text using the reverse of the encryption algorithm
func Decryption(textLine string, txtPW string) string {
	var (
		rtnStr    string
		iTemp     int
		lenPw     int
		iPw       int = 1
		iResult1  int
		iResult2  int
		iResult   int
		doDecrypt bool
	)

	txtPW = strings.ToUpper(txtPW)
	textLine = strings.ToUpper(textLine)
	lenPw = len(txtPW)

	for i := 0; i < len(textLine); i++ {
		iTemp = int(textLine[i])

		iPwTemp := int(txtPW[(iPw - 1)])
		doDecrypt = true

		if iTemp >= 'A' && iTemp <= 'Z' {
			iResult1 = iTemp - 'A' + 1
		} else if iTemp >= '0' && iTemp <= '9' {
			iResult1 = iTemp - '0' + 27
		} else {
			doDecrypt = false
		}

		if doDecrypt {
			if iPwTemp >= 'A' && iPwTemp <= 'Z' {
				iResult2 = iPwTemp - 'A' + 1
			} else if iPwTemp >= '0' && iPwTemp <= '9' {
				iResult2 = iPwTemp - '0' + 27
			}

			iResult = iResult1 - iResult2
			if iResult < 1 {
				iResult += 36
			}

			rtnStr += string(RefArray[iResult-1])
		} else {
			rtnStr += string(rune(iTemp))
		}

		iPw++
		if iPw > lenPw {
			iPw = 1
		}
	}

	return rtnStr
}
