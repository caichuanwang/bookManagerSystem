package untils

import (
	"crypto/md5"
	"fmt"
)

func CryptoWithMD5(data string) string {
	has := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", has)
}
