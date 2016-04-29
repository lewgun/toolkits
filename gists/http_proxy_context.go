// ProxyURL returns a proxy function (for use in a Transport)
// that always returns the same URL.
func ProxyURL(fixedURL *url.URL) func(* http.Request) (*url.URL, error) {
	return func(*http.Request) (*url.URL, error) {
		return fixedURL, nil
	}
}
func (gp *GooglePlay) context() context.Context {
	const defaultTimeout = time.Duration(5e9)
	transport := &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(network, addr, defaultTimeout)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(defaultTimeout))
			return c, nil
		},
		Proxy: ProxyURL(gp.proxy),
	}
	cli := &http.Client{
		Transport: transport,
	}
	return  context.WithValue(oauth2.NoContext, oauth2.HTTPClient, cli )
}
func (gp *GooglePlay) oauthHTTPClient(mfID string) (*http.Client, error) {
	iter, err := storemgr.Do("GooglePlay", mfID)
	if err != nil {
		return nil, err
	}
	gplay, ok := iter.(*types.GooglePlay)
	if !ok {
		return nil, fmt.Errorf("can't get google play config for %s", mfID)
	}
	oauth2Conf := &oauth2.Config{
		Scopes:      []string{AndroidpublisherScope},
		Endpoint:    google.Endpoint,
		ClientID:    gplay.ClientID,
		RedirectURL: titleBarRedirectURL,
	}

	tok := &oauth2.Token{
		RefreshToken: gplay.RefreshToken,
		Expiry:       time.Now().Add(-1),
	}

	var (
		c *http.Client
		ctx context.Context
	)
	if gp.proxy != nil {
		ctx = gp.context()
	} else {
		ctx = oauth2.NoContext
	}

	c = oauth2.NewClient(ctx, oauth2Conf.TokenSource(ctx, tok))  //!!!!!!!!!!!!!, ctx, ctx

	return c, nil
}


/*
    client, err := gp.oauthHTTPClient(gr.Id)
	if err != nil {
		return
	}

	url := gp.playIAPUrl(gp.config.URL, gr)
	r, _ := http.NewRequest("GET", url, nil)
	fmt.Println("CreateGPlayOrder5555555")
	rspn, err := client.Do(r)
*/
println