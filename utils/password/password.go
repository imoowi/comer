/*
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package password

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	saltPassword    = "qkhPAGA13HocW3GAEWwb"
	defaultPassword = "123456"
)

func GeneratePassword(str string) (password string) {
	// md5
	m := md5.New()
	m.Write([]byte(str))
	mByte := m.Sum(nil)

	// hmac
	h := hmac.New(sha256.New, []byte(saltPassword))
	h.Write(mByte)
	password = hex.EncodeToString(h.Sum(nil))

	return
}

func ResetPassword() (password string) {
	m := md5.New()
	m.Write([]byte(defaultPassword))
	mStr := hex.EncodeToString(m.Sum(nil))

	password = GeneratePassword(mStr)

	return
}

func GenerateLoginToken(id int32) (token string) {
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d%s", id, saltPassword)))
	token = hex.EncodeToString(m.Sum(nil))

	return
}

func GenerateMd5Str(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
