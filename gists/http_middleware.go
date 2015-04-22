package gists

//https://justinas.org/writing-http-middleware-in-go/
func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
    return &SingleHost{handler: handler, allowedHost: allowedHost}
}

func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    host := r.Host
    if host == s.allowedHost {
        s.handler.ServeHTTP(w, r)
    } else {
        w.WriteHeader(403)
    }
}

type SingleHost struct {
    handler     http.Handler
    allowedHost string
}


singleHosted = NewSingleHost(myHandler, "example.com")
http.ListenAndServe(":8080", singleHosted)

/*
func SingleHost(handler http.Handler, allowedHost string) http.Handler {
    ourFunc := func(w http.ResponseWriter, r *http.Request) {
        host := r.Host
        if host == allowedHost {
            handler.ServeHTTP(w, r)
        } else {
            w.WriteHeader(403)
        }
    }
    return http.HandlerFunc(ourFunc)
}
*/