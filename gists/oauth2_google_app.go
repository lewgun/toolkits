func TestGooglePlay(t *testing.T) {

	v := url.Values{}
	v.Set("receipt-data", raw)

	rspn, err := http.PostForm("http://127.0.0.1:8080/appstore", v)

	if err != nil {
		t.Fatal(err)
	}

	defer rspn.Body.Close()
	data, err := ioutil.ReadAll(rspn.Body)
	fmt.Println(string(data))


	const (
		AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
	)

	const (
		CLIENT_ID     = "your client id"
		CLIENT_SECRET = "your client secret "
		REFRESH_TOKEN = "your token"
	)

	const (
		PACKAGE_NAME = "packageName"
		PRODUCT_ID   = "productId"
		TOKEN        = "token"
	)

	newOAuthHttpCleint := func() *http.Client {
		oAuthClient := oauth2.NewClient(oauth2.NoContext, oauthutil.NewRefreshTokenSource(&oauth2.Config{
			Scopes:       []string{AndroidpublisherScope},
			Endpoint:     google.Endpoint,
			ClientID:     CLIENT_ID,
			ClientSecret: CLIENT_SECRET,
			RedirectURL:  oauthutil.TitleBarRedirectURL,
		}, REFRESH_TOKEN))

		return oAuthClient
	}

	url := func() string {
		baseUrl := "https://www.googleapis.com/androidpublisher/v2/applications"
		urlBuf := bytes.NewBufferString(baseUrl)
		if !strings.HasSuffix(baseUrl, "/") {
			urlBuf.WriteString("/")
		}

		m := make(map[string]string)
		m[PACKAGE_NAME] = "com.mf.dandehg.gp"
		m[PRODUCT_ID] = "ad.ko.101"
		m[TOKEN] = "dbfjnhfogpnicbcdbndgnlko.AO-J1OwXlp7amrb8KZTGZvaQNta1pLdhhYDVSKDUuJAneR5XVj68ng40Yv0xv9BUYhJQ8AWjllJrD-utRzye7cn-vJFxMhZjoOUDAlnRRvgd84R7EkUAalc"

		urlBuf.WriteString(m[PACKAGE_NAME] + "/")
		urlBuf.WriteString("purchases/products/")
		urlBuf.WriteString(m[PRODUCT_ID] + "/")
		urlBuf.WriteString("tokens/")
		urlBuf.WriteString(m[TOKEN])

		return urlBuf.String()

	}

	os.Setenv("HTTP_PROXY", "http://192.168.6.72:1080")

	client := newOAuthHttpCleint()

	r, err := http.NewRequest("GET", url(), nil)
	rspn, err := client.Do(r)
	if err != nil {
		fmt.Println(err)

	}

	defer rspn.Body.Close()

	data, err := ioutil.ReadAll(rspn.Body)
	fmt.Println(string(data))

}
