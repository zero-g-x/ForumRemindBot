package mirai

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Feed struct{
	XMLName xml.Name `xml:"urlset"`
	Urls []Url `xml:"url"`
}

type Url struct{
	XMLName xml.Name `xml:"url"`
	Loc string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func Xml(url string)[]byte{
	client:=http.Client{}
	req,_:=http.NewRequest("GET",url,nil)

	res,err:=client.Do(req)
	if err!=nil{
		fmt.Println("http get error:",err)
	}

	body,err1:=ioutil.ReadAll(res.Body)
	if err1!=nil{
		fmt.Println("read error:",err1)
	}

	return body
}

func NewFeed()(newFeed Feed){
	url:="https://hustport.com/sitemap.xml"
	xmlContent:=Xml(url)

	err:=xml.Unmarshal(xmlContent,&newFeed)
	if err!=nil{
		fmt.Println("error:",err)
	}
	return newFeed
}

func NewUrls(urls []Url)(time string,newUrl string){
	for i:=0;i<2;i++{
		if urls[i].Loc!="https://hustport.com/"{
			time=urls[i].LastMod
			newUrl=urls[i].Loc
			return
		}
	}
	return
}

func PostUrls(allUrls []Url)(urls []Url){
	for i:=0;i<len(allUrls);i++{
		if len(allUrls[i].Loc)>21&&allUrls[i].Loc[21]=='d'{
			urls = append(urls,allUrls[i])
		}
	}
	return urls
}

func NewPosts(postUrls []Url)(urls []Url){
	var ids []int
	for i:=0;i<len(postUrls);i++{
		str1:=strings.Split(postUrls[i].Loc,"/d/")[1]
		str2:=strings.Split(str1,"-")[0]
		id,_:=strconv.Atoi(str2)
		if id>LastPost{
			urls = append(urls, postUrls[i])
			ids = append(ids, id)
		}
	}
	for _,n:=range ids{
		if n>LastPost{
			LastPost=n
		}
	}
	return urls
}