package tailf

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
	_package "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/package"
	"sync"
	"time"
)

const (
	StatusNormal = 1
	StatusDelete = 2
)

type CollectConf struct {
	LogPath string
	Topic   string
}

type TailObj struct {
	tail     *tail.Tail
	conf     CollectConf
	status   int
	exitChan chan int
}

type TailObjMgr struct {
	tailObjs []*TailObj
	msgChan  chan *TextMsg
	lock     sync.Mutex
}

type TextMsg struct {
	Msg   string
	Topic string
}

var (
	tailObjMgr *TailObjMgr
)

func InitTail(conf []CollectConf, chanSize int) error {

	tailObjMgr = &TailObjMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	if len(conf) == 0 {
		logs.Error(fmt.Errorf("invalid config for log collect, conf:%v", conf))
	}

	for _, v := range conf {
		createNewTask(v)
	}

	return nil
}

func readFormTail(tailObj *TailObj) {
	for true {
		select {
		case line, ok := <-tailObj.tail.Lines:
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
		case <-tailObj.exitChan:
			logs.Warn("tail obj will exited, conf:%v", tailObj.conf)
			return
		}
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailObjMgr.msgChan
	return msg
}

func createNewTask(conf CollectConf) {
	obj := &TailObj{
		conf:     conf,
		exitChan: make(chan int, 1),
	}

	tails, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Error("collect filename[%s] failed, err:%v", conf.LogPath, err)
		return
	}

	obj.tail = tails

	tailObjMgr.tailObjs = append(tailObjMgr.tailObjs, obj)

	go readFormTail(obj)
}

func UpdateConfig(conf []CollectConf) error {
	tailObjMgr.lock.Lock()
	defer tailObjMgr.lock.Unlock()
	for _, oneConf := range conf {
		var isRunning = false
		for _, obj := range tailObjMgr.tailObjs {
			if oneConf.LogPath == obj.conf.LogPath {
				isRunning = true
				break
			}
		}

		if isRunning {
			continue
		}

		createNewTask(oneConf)
	}

	var tailObjs []*TailObj
	for _, obj := range tailObjMgr.tailObjs {
		obj.status = StatusDelete
		for _, oneConf := range conf {
			if oneConf.LogPath == obj.conf.LogPath {
				obj.status = StatusNormal
				break
			}
		}

		if obj.status == StatusDelete {
			obj.exitChan <- 1
			continue
		}
		tailObjs = append(tailObjs, obj)
	}

	tailObjMgr.tailObjs = tailObjs
	return nil
}
