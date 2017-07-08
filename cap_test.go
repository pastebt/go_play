package captcha

import (
    "os"
    "fmt"
    "net/http"
    "testing"
)


// go test -v cap.go cap_test.go -args site key
//{...captcha.test", "-test.v=true", "site", "key"}


func TestVerify(tst *testing.T) {
    tst.Logf("%#v", os.Args)
    l := len(os.Args)
    if l < 3 {
        tst.Error("need more arguments")
    }

    site := os.Args[l - 2]
    println("site = %s", site)
    rc := RC{Key: os.Args[l - 1]}

    form := fmt.Sprintf(`
        <html>
            <head>
                <script src='https://www.google.com/recaptcha/api.js'></script>
            </head>
            <body>
                <form action="/submit" method="post">
                    <div class="g-recaptcha" data-sitekey="%s"></div>
                    <input type="submit">
                </form>
            </body>
        </html>
    `, site)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, form)
    })
    http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        isValid := rc.Verify(*r)
        if isValid {
            fmt.Fprintf(w, "Valid")
        } else {
            fmt.Fprintf(w, "Invalid! These errors ocurred: %v", rc.Err)
        }
    })

    err := http.ListenAndServe(":8100", nil)
    if err != nil { tst.Error(err) }
}
