package main

import (
	"ForumRemindBot/mirai"
	"github.com/Mrs4s/MiraiGo/client"
	"log"
	"os"
	"os/signal"
)

func main(){
	//监听结束信号，ctrl+c停止
	interrupt:=make(chan os.Signal,1)
	signal.Notify(interrupt,os.Interrupt)

	c:=client.NewClient(mirai.BotNum,mirai.Password)
	c.AllowSlider=true

	r,err:=c.Login()
	if err!=nil{
		log.Println(err)
	}
	if !r.Success{
		log.Println(r.ErrorMessage)
	}

	go mirai.Reminder(c)
	<-interrupt

}
