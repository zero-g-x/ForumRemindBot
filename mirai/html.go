package mirai

import (
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type All struct{
	Resorce []Type `json:"resources"`
}

type Type struct{
	Type string `json:"type"`
	Id string `json:"id"`
	Info Attributes `json:"attributes"`
}

type Attributes struct{
	Title string `json:"title"`
	Time string `json:"lastPostedAt"`
}

func Html(url string)(body []byte){
	client:=http.Client{}
	req,_:=http.NewRequest("GET",url,nil)
	res,_:=client.Do(req)
	defer res.Body.Close()
	body,_=ioutil.ReadAll(res.Body)
	return body
}

func GetTitle(body string)string{
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(body))
	if err!=nil{
		log.Fatal(err)
	}
	content:=dom.Find("title").Last()
	return strings.Split(content.Text(),"- HUSTPORT")[0]
}

func GetScript(body string)string{
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(body))
	if err!=nil{
		log.Fatal(err)
	}
	content:=dom.Find("script").Last()
	return content.Text()
}

func GetJson(con string)string{
	str1:=strings.Split(con,"forum-dark-6122925d.css\",")[1]
	str2:=strings.Split(str1,"\"session\":")[0]
	str3:=str2[:len(str2)-1]
	return "{"+str3+"}"
}

