package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"
)

// GenMD5HashCode : 生成MD5编码
func GenMD5HashCode(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	hashCode := h.Sum(nil)
	hashCodeString := fmt.Sprintf("%x", hashCode)

	return hashCodeString
}

// EncodeAsSha1 :
func EncodeAsSha1(data, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	// return string(mac.Sum(nil))
	return hex.EncodeToString(mac.Sum(nil))
}

// EncodeAsSha512 :
func EncodeAsSha512(data, salt string) string {
	mac := hmac.New(sha512.New, []byte(salt))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// EncodeAsSha512ToBase64 :
func EncodeAsSha512ToBase64(data, salt string) string {
	mac := hmac.New(sha512.New, []byte(salt))
	mac.Write([]byte(data))
	base64edMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return base64edMAC
}

// GenInternalCommSha1Hash :
func GenInternalCommSha1Hash(secret string) (string, string) {
	nanoSecs := strconv.Itoa(time.Now().Nanosecond())
	sha1Hash := EncodeAsSha1(nanoSecs, secret)
	return sha1Hash, nanoSecs
}
