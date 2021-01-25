package util

/**
* @Author: super
* @Date: 2021-01-25 14:18
* @Description:
**/

import "github.com/skip2/go-qrcode"

func GenerateQRCodeByte(str string) ([]byte, error) {
	return qrcode.Encode(str, qrcode.Highest, 256)
}
