//  Created by 摸金校尉 on 17/11/20.
package UselessDownload

import (
	"time"
	"net/http"
	"net/url"
	"math/rand"
	"errors"
	"net/http/cookiejar"
)

var MyJar *cookiejar.Jar
var myProxy *url.URL
var FiddlerProxy *url.URL
//var MyClient *http.Client
var MyClient_WithoutProxy *http.Client
var transport_WithoutProxy http.Transport
var MyClient_WithoutProxy_2S *http.Client
var transport_WithoutProxy_2S http.Transport
var MyClient_WithoutProxy_500MS  *http.Client
var transport_WithoutProxy_500MS  http.Transport
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var myTimeout = time.Duration(15000* time.Millisecond)
var myTimeout_2S = time.Duration(2000 * time.Millisecond)
var myTimeout_500MS = time.Duration(500* time.Millisecond)
type proxyStruct struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Proxy []string `json:"data"`
}
var proxyServerERR =errors.New("proxyServerERR")
var jsonunmarshalERR =errors.New("jsonunmarshalERR")
var fiddlerProxyAdress string = "http://127.0.0.1:8888"