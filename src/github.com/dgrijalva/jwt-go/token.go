package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

// TimeFunc provides the current time when parsing token to validate "exp" claim (expiration time).
// You can override it to use another time value.  This is useful for testing or if your
// server uses a different time zone than your tokens.
// TimeFunc提供解析令牌以验证“exp”声明（到期时间）时的当前时间。
//您可以覆盖它以使用其他时间值。 这对于测试或者如果您的测试很有用
//服务器使用与令牌不同的时区。
var TimeFunc = time.Now

// Parse methods use this callback function to supply
// the key for verification.  The function receives the parsed,
// but unverified Token.  This allows you to use properties in the
// Header of the token (such as `kid`) to identify which key to use.
//解析方法使用此回调函数来提供
//验证的关键 该函数接收解析后的，
//但未经验证的令牌。 这允许您使用中的属性
//标记的标题（例如`kid`）标识要使用的密钥。
type Keyfunc func(*Token) (interface{}, error)

// A JWT Token.  Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
//一个JWT令牌。 根据您的使用情况，将使用不同的字段
//创建或解析/验证令牌。
type Token struct {
	Raw       string                 // 原始token.解析token时填充
	Method    SigningMethod          // 使用或将要使用的签名方法
	Header    map[string]interface{} // token的第一部分
	Claims    Claims                 //  token的第二部分
	Signature string                 // token的第三部分.解析token时填充
	Valid     bool                   // token有效吗？ 解析/验证token时填充
}

// 创建一个新的令牌。 采用签名方法
func New(method SigningMethod) *Token {
	return NewWithClaims(method, MapClaims{})
}

func NewWithClaims(method SigningMethod, claims Claims) *Token {
	return &Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

// 获取完整的签名令牌
func (t *Token) SignedString(key interface{}) (string, error) {
	var sig, sstr string
	var err error
	if sstr, err = t.SigningString(); err != nil {
		return "", err
	}
	if sig, err = t.Method.Sign(sstr, key); err != nil {
		return "", err
	}
	return strings.Join([]string{sstr, sig}, "."), nil
}

// 生成签名字符串.  This is the
// most expensive part of the whole deal.  Unless you
// need this for something special, just go straight for
// the SignedString.
func (t *Token) SigningString() (string, error) {
	var err error
	parts := make([]string, 2)
	for i, _ := range parts {
		var jsonValue []byte
		if i == 0 {
			if jsonValue, err = json.Marshal(t.Header); err != nil {
				return "", err
			}
		} else {
			if jsonValue, err = json.Marshal(t.Claims); err != nil {
				return "", err
			}
		}

		parts[i] = EncodeSegment(jsonValue)
	}
	return strings.Join(parts, "."), nil
}

 //解析，验证并返回令牌。
// keyFunc将接收已解析的令牌，并应返回用于验证的密钥。
// If everything is kosher, err will be nil
func Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
	return new(Parser).Parse(tokenString, keyFunc)
}

func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
	return new(Parser).ParseWithClaims(tokenString, claims, keyFunc)
}

// Encode JWT specific base64url encoding with padding stripped 使用填充剥离编码JWT特定的base64url编码
func EncodeSegment(seg []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
}

// Decode JWT specific base64url encoding with padding stripped 使用填充剥离解码JWT特定的base64url编码
func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}
