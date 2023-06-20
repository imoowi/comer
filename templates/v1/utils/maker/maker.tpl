/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package maker

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

func MakeSn(prefix string) string {
	now := time.Now()
	date := now.Format(`20060102`)
	r := rand.Intn(1000)
	timeTick := time.Now().UnixNano() / 1e6
	code := prefix + fmt.Sprintf(`%s%d%03d`, date, timeTick, r)
	return code
}

func MakeRandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
