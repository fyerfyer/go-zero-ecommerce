package encrypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5String(str string) (string, error) {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return "", err
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr), nil
}

func Md5Bytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
