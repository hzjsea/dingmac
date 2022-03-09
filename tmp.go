package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os/exec"
	"time"
)

/*
弹窗类型
*/
type PopupWindow struct {
	message string
	action string
	icon string
}

/*
通知类型
*/
type NoticeWindow struct {
	message string
	action string
	mainTitle string
	subTitle string
	soundName string
}

type Notify interface {
	String() string
	Send() (err error, n time.Time, msg string)
}

func (w *PopupWindow) String() string{
	popMessage := "%s \"%s\" %s"
	return fmt.Sprintf(popMessage, w.action, w.message, w.icon)
}

func (w *NoticeWindow ) String() string{
	notifyMessage := "%s \"%s\" with title \"%s\" subtitle \"%s\" sound name \"%s\""
	return fmt.Sprintf(notifyMessage, w.action, w.message, w.mainTitle, w.subTitle, w.soundName)
}

func (w *NoticeWindow) Send() (err error, n time.Time, msg string){
	cmd := exec.Command("/usr/bin/osascript", "-e" , w.String() )
	err = cmd.Run()
	if err != nil{
		return err, time.Now(), cmd.String()
	}
	return nil, time.Now(), ""
}

func (w *PopupWindow) Send() (err error, n time.Time, msg string){
	cmd := exec.Command("/usr/bin/osascript", "-e" , w.String() )
	err = cmd.Run()
	if err != nil{
		return err, time.Now(), cmd.String()
	}
	return nil, time.Now(), ""
}


func PollDing(notify Notify){
	c := cron.New()
	c.AddFunc("@hourly", func() {
	//c.AddFunc("@every 1s", func() {
		err, n, msg := notify.Send()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(n, msg)
	})
	c.Start()
	c.Run()
}


