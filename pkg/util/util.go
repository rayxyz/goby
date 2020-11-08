package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"time"

	"goby/pkg/crypto"

	"github.com/microcosm-cc/bluemonday"
	log "github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/gomail.v2"
)

// EncodeStringWithSalt :
func EncodeStringWithSalt(str, salt string) string {
	return crypto.EncodeAsSha512ToBase64(str, salt)
}

// EncodeString :
func EncodeString(str string) (string, string) {
	nanoSecStr := big.NewInt(time.Now().UnixNano()).String()
	sha := sha1.New()
	sha.Write([]byte(nanoSecStr))
	salt := fmt.Sprintf("%x", sha.Sum(nil))

	mac := hmac.New(sha512.New, []byte(salt))
	mac.Write([]byte(str))
	base64edMAC := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	return base64edMAC, salt
}

// GenUUID :
func GenUUID() string {
	return uuid.Must(uuid.NewV4(), nil).String()
}

// GetDayOfWeekCN : 将一周中的某天转成中文
func GetDayOfWeekCN(dayOfWeek int) string {
	data := map[int]string{
		0: "日",
		1: "一",
		2: "二",
		3: "三",
		4: "四",
		5: "五",
		6: "六",
	}
	return data[dayOfWeek]
}

// GenMD5HashCode : 生成MD5编码
func GenMD5HashCode(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	hashCode := h.Sum(nil)
	hashCodeString := fmt.Sprintf("%x", hashCode)
	return hashCodeString
}

// MailCC :
type MailCC struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

// MailPayload :
type MailPayload struct {
	Host       string    `json:"host"`
	Port       int       `json:"port"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	From       string    `json:"from"`
	To         []string  `json:"to"`
	CcList     []*MailCC `json:"cc"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	AttachList []string  `json:"attach_list"`
}

// SendEmail :
func SendEmail(payload *MailPayload) error {
	log.Info("Sending E-mail....")
	// log.Println("payload => ", payload)
	log.Info("From: ", payload.From)
	log.Info("To: ", payload.To)
	m := gomail.NewMessage()
	m.SetHeader("From", payload.From)
	m.SetHeader("To", payload.To...)
	for _, v := range payload.CcList {
		m.SetAddressHeader("Cc", v.Address, v.Name)
	}
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/html", payload.Body)
	for _, v := range payload.AttachList {
		m.Attach(v)
	}

	d := gomail.NewDialer(payload.Host, payload.Port, payload.UserName, payload.Password)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Error("send Email error => ", err)
		return err
	}

	log.Info("E-mail has been successfullly sent.")

	return nil
}

// ExtractHTMLText :
func ExtractHTMLText(content string, length int) string {
	content = bluemonday.StrictPolicy().Sanitize(content)
	runes := []rune(content)
	if len(runes) > length {
		content = string(runes[:length]) + "..."
	}
	return content
}

// Truncate :
func Truncate(s string, i int, dots bool) string {
	if i > 0 {
		runes := []rune(s)
		if len(runes) > i {
			s = string(runes[:i])
			if dots {
				s += "..."
			}
		}
	}
	return s
}
