// 支付宝 生活号接入
package alipay_fuwu

import (
	"bytes"
	"crypto"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"

	//"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	//"strings"

	"surpass/internal/protocol"

	"surpass/internal/component/thirdlogin_provider"
)

const (
	appId = "..." //prod


	myPrivKey = `
-----BEGIN RSA PRIVATE KEY-----
...
-----END RSA PRIVATE KEY-----
`
	myPubKey = `...`
  
	//prod
	alipayPubKey = `...`

	//prod
	alipayGateway = "https://openapi.alipay.com/gateway.do"

	//qa
	//alipayGateway = "https://openapi.alipaydev.com/gateway.do"
)

func init() {
	provider.MustRegister(provider.AlipayFuWu, &fuwu{})
}

type (
	LoginReq struct {
		Code string
	}

	alipayGatewayCheck struct {
		Service  string `json:"service"`
		Sign     string `json:"sign"`
		Charset  string `json:"charset"`
		SignType string `json:"sign_type"`

		//Content AlipayBizContent
		Content string `json:"content"`
	}
)

type fuwu struct{}

func (p *fuwu) Setup() error {
	return nil
}

func (p *fuwu) GatewayCheck(appId uint64, data string) (string, error) {

	check := alipayGatewayCheck{}

	err := json.Unmarshal([]byte(data), &check)
	if err != nil {
		return "", err
	}

	raw := fmt.Sprintf("biz_content=%s&charset=%s&service=%s&sign_type=%s", check.Content, check.Charset, check.Service, check.SignType)

	err = verify(raw, check.Sign, alipayPubKey)
	if err != nil {
		return "", err
	}

	raw = fmt.Sprintf("<biz_content>%s</biz_content><success>true</success>", myPubKey)

	return makeResponse(raw, myPrivKey)
}

func (p *fuwu) UserInfo(code string, data interface{}) (*protocol.ThirdUserInfo, error) {
	uid, tok, err := fetchAccessToken(code)
	if err != nil {
		return nil, err
	}

	userInfo, err := fetchUserInfo(uid, tok)
	//fmt.Println(userInfo.Avatar, userInfo.Province, userInfo.City, userInfo.Account, userInfo.Nickname)
	return userInfo, err

}

func genPubKey(key string) (pubKey *rsa.PublicKey, err error) {

	encodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return
	}

	pkix, err := x509.ParsePKIXPublicKey(encodedKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pxix key")
	}

	ok := false

	if pubKey, ok = pkix.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("aliPubKey can not be parsed to rsa.PublicKey")
	}
	return
}

// verify 验签函数
func verify(body, sign, aliPubKey string) error {
	decoded, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(body))

	pubKey, err := genPubKey(aliPubKey)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, h.Sum(nil), decoded)
}

func rsaSign(content, privKey string) (string, error) {
	p, _ := pem.Decode([]byte(privKey))
	key, _ := x509.ParsePKCS1PrivateKey(p.Bytes)

	hashed := sha256.Sum256([]byte(content))
	signed, _ := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	return base64.StdEncoding.EncodeToString(signed), nil
}

func makeResponse(content, privKey string) (string, error) {
	builder := `<?xml version="1.0" encoding="GBK" ?>
				<alipay>
					<response>%s</response>
					<sign>%s</sign>
					<sign_type>RSA2</sign_type>
				</alipay>`

	sign, err := rsaSign(content, privKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(builder, content, sign), nil
}

func signIt(m url.Values) (string, error) {

	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	buf := &bytes.Buffer{}
	for _, k := range keys {
		fmt.Fprintf(buf, "%s=%s&", k, m[k][0])
	}

	buf.Truncate(buf.Len() - 1)

	plain := buf.String()

	signed, err := rsaSign(plain, myPrivKey)
	if err != nil {
		return "", err
	}

	return signed, nil

}

type accessToken struct {
	Sign  string `json:"sign"`
	Token struct {
		AlipayUserId string `json:"alipay_user_id"`
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
	} `json:"alipay_system_oauth_token_response"`
}

//https://docs.open.alipay.com/197/105238/
type userInfo struct {
	Sign string `json:"sign"`
	Info struct {
		Avatar   string `json:"avatar"`
		NickName string `json:"nick_name"`
		Gender   string `json:"gender"`
		City     string `json:"city"`
		Province string `json:"province"`
	} `json:"alipay_user_info_share_response"`
}

func fetchAccessToken(code string) (string, string, error) {

	data := url.Values{}
	data["app_id"] = []string{appId}
	data["method"] = []string{"alipay.system.oauth.token"}
	data["charset"] = []string{"GBK"}
	data["sign_type"] = []string{"RSA2"}
	data["version"] = []string{"1.0"}
	data["grant_type"] = []string{"authorization_code"}

	data["timestamp"] = []string{time.Now().Format("2006-01-02 15:04:05")}
	data["code"] = []string{code}
	sign, err := signIt(data)

	if err != nil {
		return "", "", err
	}
	data["sign"] = []string{sign}

	resp, err := http.PostForm(alipayGateway, data)
	if err != nil {
		fmt.Println(err)
		return "", "", err

	}

	defer resp.Body.Close()

	raw, _ := ioutil.ReadAll(resp.Body)

	d := json.NewDecoder(bytes.NewReader(raw))
	tok := accessToken{}
	err = d.Decode(&tok)
	if err != nil {
		return "", "", err
	}

	return tok.Token.AlipayUserId, tok.Token.AccessToken, nil

}

func GBKToUTF8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func fetchUserInfo(uid string, accessToken string) (*protocol.ThirdUserInfo, error) {

	data := url.Values{}
	data["app_id"] = []string{appId}
	data["method"] = []string{"alipay.user.info.share"}

	data["charset"] = []string{"GBK"}
	data["sign_type"] = []string{"RSA2"}
	data["version"] = []string{"1.0"}
	data["timestamp"] = []string{time.Now().Format("2006-01-02 15:04:05")}
	data["auth_token"] = []string{accessToken}

	sign, err := signIt(data)

	if err != nil {
		return nil, err
	}

	data["sign"] = []string{sign}

	resp, err := http.PostForm(alipayGateway, data)
	if err != nil {
		fmt.Println(err)
		return nil, err

	}

	defer resp.Body.Close()

	gbkRaw, _ := ioutil.ReadAll(resp.Body)

	raw, _ := GBKToUTF8(gbkRaw)

	d := json.NewDecoder(bytes.NewReader(raw))

	var u userInfo
	err = d.Decode(&u)
	if err != nil {
		fmt.Println(err)

	}

	const (
		male   = 1
		female = 2
	)

	var gender int32 = female

	if u.Info.Gender == "m" {
		gender = male
	}
	return &protocol.ThirdUserInfo{
		Account: uid,

		Nickname: u.Info.NickName,
		Gender:   gender,
		City:     u.Info.City,
		Province: u.Info.Province,
		Avatar:   u.Info.Avatar,
	}, nil

}
