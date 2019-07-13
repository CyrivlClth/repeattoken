package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/CyrivlClth/repeattoken/token"
)

func generate(key string, data token.Data) (token string, err error) {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var str []string
	for _, k := range keys {
		var val string
		switch v := data[k].(type) {
		case string:
			val = v
		case int:
			val = strconv.Itoa(v)
		case float64:
			val = strconv.FormatFloat(v, 'f', -1, 64)
		case bool:
			val = strconv.FormatBool(v)
		default:
			val = ""
		}
		val = strings.TrimSpace(val)
		if val != "" {
			str = append(str, k+"="+val)
		}
	}
	h := md5.New()
	h.Write([]byte(strings.Join(str, "&") + "&key=" + key))
	token = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	return
}

type digest struct{}

// Generate 生成签名
func (d *digest) Generate(key string, data token.Data) (token string, err error) {
	return generate(key, data)
}

// Verify 验证签名
func (d *digest) Verify(key string, data token.Data) bool {
	sign, ok := data["sign"]
	if !ok {
		return false
	}
	z, err := json.Marshal(data)
	if err != nil {
		return false
	}
	var a token.Data
	err = json.Unmarshal(z, &a)
	if err != nil {
		return false
	}
	delete(a, "sign")
	nToken, err := d.Generate(key, a)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return nToken == sign
}

// New 生成验证器
// 该验证器验证签名时将拷贝data数据
func New() token.Token {
	return &digest{}
}

type fastDigest struct{}

// Generate 生成签名
func (f *fastDigest) Generate(key string, data token.Data) (token string, err error) {
	return generate(key, data)
}

// Verify 验证签名
func (f *fastDigest) Verify(key string, data token.Data) bool {
	sign, ok := data["sign"]
	if !ok {
		return false
	}
	delete(data, "sign")
	nToken, err := f.Generate(key, data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return nToken == sign
}

// NewFast 生成快速验证器
// 该验证器验证签名时会直接删除data中的sign字段，但验证速度更快
func NewFast() token.Token {
	return &fastDigest{}
}
