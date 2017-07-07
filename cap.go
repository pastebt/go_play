package main

import (
    //"os"
    "fmt"
    "time"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

// https://github.com/haisum/recaptcha/blob/master/recaptcha.go
// google response, for decode json
type GR struct {
    Ok    bool          `json:"success"`
    ErrorCodes []string `json:"error-codes"`
}


// post re-captcha response to
var RcURL = "https://www.google.com/recaptcha/api/siteverify"


// Recatpcha
type RC struct {
    Key     string
    Err     error
}


func (rc *RC)Verify(req http.Request) bool {
    resp := req.FormValue("g-recaptcha-response")
    return rc.VerifyString(resp)
}


func (rc *RC)VerifyString(dat string) bool {
    client := &http.Client{Timeout: 20 * time.Second}
    resp, err := client.PostForm(RcURL,
                                 url.Values{"secret": {rc.Key},
                                            "response": {dat}})
    if err != nil { rc.Err = err; return false }
    defer resp.Body.Close()

    gr := new(GR)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil { rc.Err = err; return false }
    if rc.Err = json.Unmarshal(body, gr); rc.Err != nil { return false }

    if !gr.Ok { rc.Err = fmt.Errorf(string(body)) }
    return gr.Ok
}
