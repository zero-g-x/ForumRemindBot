package mirai

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"time"
)

const(
	BotNum=123456789 //bot QQ号
	Password="admin"//密码
	TargetGroup=987654321 //群组号码
)

var LastPost=217

func Reminder(c *client.QQClient){
	for true{
		now:=time.Now()
		nextTime:=time.Date(now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),0,0,
			now.Location()).Add(time.Hour)
		delta:=nextTime.Unix()-now.Unix()
		time.Sleep(time.Duration(delta)*time.Second)

		feed:=NewFeed()
		postUrls:=PostUrls(feed.Urls)
		newPosts:=NewPosts(postUrls)
		for _,p:=range newPosts{
			address:=p.Loc
			body:=string(Html(address))
			title:=GetTitle(body)
			text:=message.NewText("论坛更新:\n"+title+"\n"+address)
			msg:=message.SendingMessage{
				Elements: []message.IMessageElement{text},
			}
			c.SendGroupMessage(TargetGroup,&msg)
		}
	}
}