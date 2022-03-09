package main

import "time"

type Message struct {
	action string
}


type defaultMessage Message

/*
弹窗类型
*/
type PopupWindow struct {
	defaultMessage
	action string
}

/*
通知类型
*/
type NoticeWindow struct {
	defaultMessage
	action string
	mainTitle string
	subTitle string
	soundName string
}

type Notify interface {
	String() string
	Send() (err error, n time.Duration, msg string)
}

func (w *PopupWindow) String() string{
	return ""
}

func (w *NoticeWindow ) String() string{
	return ""
}

func (w *NoticeWindow) Send() (err error, n time.Time, msg string){
	return nil, time.Now(), ""
}

func (w *PopupWindow) Send() (err error, n time.Time, msg string){
	return nil, time.Now(), ""
}


// 定时
type CronTask struct {

}