package Hash_local

import (
	"crypto/md5"
	"encoding/hex"
)

type hasherInfo struct {
}

func NewHashInfo() *hasherInfo {
	return &hasherInfo{}
}

func (*hasherInfo) HashMd5(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
