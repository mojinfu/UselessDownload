package UselessDownload

import (
	"io/ioutil" 
	"net/http"
	//"strings"
	"bytes"
)
//var myTimeout_120ms = time.Duration(120* time.Millisecond)
func Download_POST_WAX(myUrl string,myJosnBody []byte) (string , error) {
	httpreq ,errr := http.NewRequest("POST",myUrl, bytes.NewReader(myJosnBody))
	if errr!=nil{
		return "" ,errr
	}
	httpreq.Header.Add("Content-Type","application/json")
	resp, errr := MyClient_WithoutProxy_2S.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr!=nil{
		return "" ,errr
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr!=nil{
		return "" ,errr	
	}
	return  string(body) , nil
}

func Download_POST_Business(myUrl string,myJosnBody []byte) (string , error) {
	httpreq ,errr := http.NewRequest("POST",myUrl, bytes.NewReader(myJosnBody))
	if errr!=nil{
		return "" ,errr
	}
	resp, errr := MyClient_WithoutProxy_2S.Do(httpreq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr!=nil{
		return "" ,errr
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr!=nil{
		return "" ,errr	
	}
	return  string(body) , nil
}

func Do_Request(myReq *http.Request)(string , error){
	resp, errr := MyClient_WithoutProxy.Do(myReq)
	if resp != nil {
		defer resp.Body.Close()
	}
	if errr!=nil{
		return "" ,errr
	}
	body, errr := ioutil.ReadAll(resp.Body)
	if errr!=nil{
		return "" ,errr	
	}
	return  string(body) , nil
}