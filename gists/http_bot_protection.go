func isBot(r *http.Request) bool {
	agent := r.Header.Get("User-Agent")
	return strings.Contains(agent, "Baidu") || strings.Contains(agent, "bingbot") ||
	strings.Contains(agent, "Ezooms") || strings.Contains(agent, "Googlebot")
}

type noWwwHandler struct {
	Handler http.Handler
}

func (h *noWwwHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Some bots (especially Baidu) don't seem to respect robots.txt and swamp gitweb.cgi,
	// so explicitly protect it from bots.
	if strings.Contains(r.URL.RawPath, "/code/") && strings.Contains(r.URL.RawPath, "?") && isBot(r) {
		http.Error(rw, "bye", http.StatusUnauthorized)
		log.Printf("bot denied")
		return

	}

	host := strings.ToLower(r.Host)
	if host == "www.camlistore.org" {
		http.Redirect(rw, r, "http://camlistore.org"+r.URL.RawPath, http.StatusFound)
		return
	}
	h.Handler.ServeHTTP(rw, r)
}