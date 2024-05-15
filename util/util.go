package util

import (
	"os"

	"crypto/rand"
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

const DEFAULT_FILE_PERM = 0666
const (
	PW_COST      = 32768
	PW_BLOCKSIZE = 8
	PW_PARALLEL  = 3
	PW_HASH_LEN  = 64
)

func IsValidFsPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else {
		return false
	}
}

func DecodePostBody[T any](req *http.Request) (*T, error) {
	dec := json.NewDecoder(req.Body)
	var obj T
	if err := dec.Decode(&obj); err == nil {
		return &obj, nil
	} else {
		return nil, err
	}
}

func GenSalt(size uint32) []byte {
	bs := make([]byte, size)
	_, err := rand.Read(bs)
	if err != nil {
		log.Infof("EOF received when generating salt of size: %d :: %s", size, err)
	}
	return bs
}

// func GenSaltStr(size int) string {
//
// }
