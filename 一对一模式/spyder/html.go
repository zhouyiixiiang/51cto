package spyder

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Html struct {

}

func (item *Html)GetUrl(url string)[]string{
	resp,err:=http.Get(url)
	if err !=nil{
		return []string{}
	}
	b,err:=ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err !=nil{
		return []string{}
	}
	file,err:=os.Create("1.txt")
	file.Write(b)
	file.Close()
	newFile,_:=os.Open("1.txt")
	fmt.Println(newFile.Name())
	//doc,err:=html.Parse(newFile)
	tmpList:=[]string{}
	return tmpList
}

func (item *Html)GetMail(url string)[]string{
	resp,err:=http.Get(url)
	if err!=nil{
		return []string{}
	}
	b,err:=ioutil.ReadAll(resp.Body)
	reg:=`[\w-]+(\.[\w-]+)*@[\w-]+(\.[w-]+)+` //提取email
	rgx:=regexp.MustCompile(reg) //编译
	tmplist:=rgx.FindAllString(string(b),-1)
	return tmplist
}
func (item *Html)GetQQ(url string)[]string{
	resp,err:=http.Get(url)
	if err!=nil{
		return []string{}
	}
	b,err:=ioutil.ReadAll(resp.Body)
	reg:=`[1-9][0-9]{4,11}` //提取email
	rgx:=regexp.MustCompile(reg) //编译
	tmplist:=rgx.FindAllString(string(b),-1)
	return tmplist
}