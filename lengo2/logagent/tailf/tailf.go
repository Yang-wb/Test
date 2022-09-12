package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	"time"
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail *tail.Tail
	conf CollectConf
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
}

type TextMsg struct {
	Msg   string
	Topic string
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) error {

	if len(conf) == 0 {
		return fmt.Errorf("invalid config for log collect, conf:%v", conf)
	}

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	for _, v := range conf {
		obj := &TailObj{
			conf: v,
		}

		tails, err := tail.TailFile(v.LogPath, tail.Config{
			ReOpen: true,
			Follow: true,
			//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})
		if err != nil {
			return fmt.Errorf("tail file err:%v", err)
		}

		obj.tail = tails

		tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

		go readFormTail(obj)
	}

	return nil
}

func readFormTail(tailObj *TailObj) {
	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textMsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}

		tailObjMgr.msgChan <- textMsg
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return msg
}
