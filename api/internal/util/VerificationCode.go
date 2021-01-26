package util

import (
	"fmt"
	"math/rand"
	"time"
)

/**
* @Author: super
* @Date: 2021-01-26 21:51
* @Description:
**/

func GenerateVerificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}
