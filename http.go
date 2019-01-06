package main

import (
	"io"
	"log"
	"net/http"
)


//curl -v http://web.archive.org/web/20160513214322/http:/justpaste.it/u8j7//
//*   Trying 207.241.224.26...
//* TCP_NODELAY set
//* Connected to web.archive.org (207.241.224.26) port 80 (#0)
//> GET /web/20160513214322/http:/justpaste.it/u8j7// HTTP/1.1
//> Host: web.archive.org
//> User-Agent: curl/7.55.1
//> Accept: */*
//> 
//< HTTP/1.1 200 OK
//< Server: nginx/1.15.5
//< Date: Sat, 05 Jan 2019 00:49:43 GMT
//< Content-Type: text/html; charset=UTF-8
//< Content-Length: 35579
//< Connection: keep-alive
//< X-Archive-Orig-content-length: -1
//< X-Archive-Orig-strict-transport-security: max-age=15552000; includeSubDomains; preload
//< X-Archive-Orig-x-content-type-options: nosniff
//< X-Archive-Orig-set-cookie: PHPSESSID=ht1e342e9re2dvetp0o6n7e8oe8pbs3jdh31lpqd3ne7iuubo7j0; path=/; HttpOnly
//< X-Archive-Orig-expires: Thu, 19 Nov 1981 08:52:00 GMT
//< X-Archive-Orig-vary: Accept-Encoding
//< X-Archive-Orig-server: cloudflare-nginx
//< X-Archive-Orig-connection: close
//< X-Archive-Orig-pragma: no-cache
//< X-Archive-Orig-cache-control: no-store, no-cache, must-revalidate
//< X-Archive-Orig-date: Fri, 13 May 2016 21:43:22 GMT
//< X-Archive-Orig-cf-ray: 2a29491c158a1207-SJC
//< X-Archive-Guessed-Content-Type: text/html
//< X-Archive-Guessed-Charset: utf-8
//< Memento-Datetime: Fri, 13 May 2016 21:43:22 GMT
//< Link: <https://justpaste.it/u8j7>; rel="original", <http://web.archive.org/web/timemap/link/https://justpaste.it/u8j7>; rel="timemap"; type="application/link-format", <http://web.archive.org/web/https://justpaste.it/u8j7>; rel="timegate", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="first memento"; datetime="Fri, 13 May 2016 21:43:22 GMT", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="memento"; datetime="Fri, 13 May 2016 21:43:22 GMT", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="last memento"; datetime="Fri, 13 May 2016 21:43:22 GMT"
//< Content-Security-Policy: default-src 'self' 'unsafe-eval' 'unsafe-inline' data: blob: archive.org web.archive.org analytics.archive.org pragma.archivelab.org
//< X-Archive-Src: liveweb-20160514001702/live-20160513205105-wwwb-app11.us.archive.org.warc.gz
//< X-App-Server: wwwb-app12
//< X-ts: ----
//< X-location: All
//< X-Cache-Key: httpweb.archive.org/web/20160513214322/http:/justpaste.it/u8j7//CA
//< X-Page-Cache: MISS
//< 


func main1() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
        //w.Header().Set("Content-Type", "text/html; charset=utf-8")
//        w.Header().Set("Link", `<https://justpaste.it/u8j7>; rel="original", <http://web.archive.org/web/timemap/link/https://justpaste.it/u8j7>; rel="timemap"; type="application/link-format", <http://web.archive.org/web/https://justpaste.it/u8j7>; rel="timegate", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="first memento"; datetime="Fri, 13 May 2016 21:43:22 GMT", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="memento"; datetime="Fri, 13 May 2016 21:43:22 GMT", <http://web.archive.org/web/20160513214322/https://justpaste.it/u8j7>; rel="last memento"; datetime="Fri, 13 May 2016 21:43:22 GMT"`)
//        w.Header().Set("X-Cache-Key", `httpweb.archive.org/web/20160513214322/http:/justpaste.it/u8j7//CA`)
//        w.Header().Set("X-Archive-Orig-vary", "Accept-Encoding")
//        w.Header().Set("X-location", "All")
		io.WriteString(w, req.RequestURI + "\n")
	}

	http.HandleFunc("/web", helloHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}


type MyMux struct {
//    http.ServeMux
}


func (m *MyMux)ServeHTTP(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, req.RequestURI + "\n")
}

/*
func (m *MyMux)Handler(r *http.Request) (h http.Handler, pattern string) {
    println("MyMux)Handler")
    return http.HandlerFunc(m.ServeHTTP), "/"
}
*/

func main() {
	http.ListenAndServe(":8081", new(MyMux))
}
