package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type Encrypt struct {
	Salt     string
	S string
}

func (e *Encrypt) MD5EncryptS() string {
	h := md5.New()
	h.Write([]byte(e.Salt + e.S))
	return hex.EncodeToString(h.Sum(nil))
}

func (e *Encrypt) CheckMD5EncryptS(encryptS string) bool {
	return strings.EqualFold(e.MD5EncryptS(), encryptS)
}
