package spark

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/shenfz/TabakoAssistant/backend/global"
	"github.com/shenfz/TabakoAssistant/backend/pkg/zapLogger"
	"sync"
	"time"
)

/**
 * @Author shenfz
 * @Date 2024/3/4 10:52
 * @Email 1328919715@qq.com
 * @Description:
 **/

type SparkOption func(s *SparkWsConn)

type SparkWsConn struct {
	ctx           context.Context
	onWorking     bool
	handShakeTime time.Duration
	authURL       string
	wssConn       *websocket.Conn
	logger        *zapLogger.LoggerXC
	locker        *sync.RWMutex
	payloadCache  map[string][]Text
	sendChan      chan SparkReqMsg
	acceptChan    chan SparkRespMsg
}

func GetSparkWsConn(ctx context.Context, opts ...SparkOption) *SparkWsConn {
	sparkC := SparkWsConn{
		ctx:           ctx,
		onWorking:     false,
		handShakeTime: time.Second * 5,
		wssConn:       nil,
		logger:        zapLogger.NewLogger(zapLogger.SetDebugMode()),
		locker:        &sync.RWMutex{},
		payloadCache:  make(map[string][]Text),
		sendChan:      make(chan SparkReqMsg, 5),
		acceptChan:    make(chan SparkRespMsg, 5),
	}

	for _, opt := range opts {
		opt(&sparkC)
	}
	go sparkC.HandleStreamRespContent()

	sparkC.InitSparkWssConn()
	if sparkC.wssConn != nil {
		sparkC.SetOnWorking()
	}
	return &sparkC
}

func (s *SparkWsConn) InitSparkWssConn() {
	d := websocket.Dialer{
		HandshakeTimeout: s.handShakeTime,
	}
	authURL, err := AssembleAuthUrl()
	if err != nil {
		s.logger.Errorf("Auth url Fail, Error: %v", err)
		return
	}

	//握手并建立websocket 连接
	conn, resp, err := d.Dial(authURL, nil)
	if err != nil {
		s.logger.Errorf("Dail URL[%s] Fail, Error: %v", authURL, err)
		return
	}
	s.logger.Infof("Dail URL [%s]  Succeed, Resp: %+v ", authURL, resp)

	s.wssConn = conn
	s.authURL = authURL
	s.wssConn.SetCloseHandler(s.CloseHandler)
}

func (s *SparkWsConn) CloseHandler(code int, text string) error {
	s.logger.Warn(fmt.Sprintf("CLosed wss conn code[ %d ] msg[ %s ]", code, text))
	return global.AutoCloseWssConn
}

func (s *SparkWsConn) SetOnWorking() {
	s.onWorking = true
	go s.GorSend()
	go s.GorRead()
}

func (s *SparkWsConn) SetOffWorking() {
	s.onWorking = false
	s.wssConn.Close()
	s.wssConn = nil
}

func (s *SparkWsConn) Close() {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.onWorking = false
	s.wssConn.Close()
	close(s.acceptChan)
	close(s.sendChan)
	s.logger.Warn("Close Spark Wss Conn")
}

func (s *SparkWsConn) IsOnWorking() bool {
	return s.onWorking
}

func SetHandshakeTimeOut(t time.Duration) SparkOption {
	return func(s *SparkWsConn) {
		s.handShakeTime = t
	}
}

func (s *SparkWsConn) TestEcho() {
	req := NewSparkRequestMsg(EchoSetOption{})
	s.sendChan <- req
}

func (s *SparkWsConn) GorSend() {
	defer s.logger.Warn("Send channel closed ,exit gor")
	s.logger.Info("Start Send Gor")
	for s.onWorking {
		select {
		case sendObj, ok := <-s.sendChan:
			if !ok {
				return
			}
			// check payloadCache
			sendObj.ReqPayload.Message.Text = s.UpUserPayload(sendObj.ReqPayload.Message.Text)
			// do send
			err := s.wssConn.WriteJSON(sendObj)
			if err != nil {
				go s.Check(err)
				continue
			}
			s.logger.Debug("Get Obj send Succeed: " + fmt.Sprintf("%+v", sendObj))
		default:
		}
	}
}

func (s *SparkWsConn) Check(err error) {
	if errors.Is(err, global.AutoCloseWssConn) {
		s.logger.Error("auto closed wss conn by server, now restart conn")
	} else {
		s.logger.Error(err.Error() + ", now restart conn")
	}
	s.SetOffWorking()

	s.InitSparkWssConn()
	if s.wssConn != nil {
		s.SetOnWorking()
	}
}

func (s *SparkWsConn) GorRead() {
	defer s.logger.Warn("ReadChan closed, exit read gor")
	s.logger.Info("Start Read Gor")
	for s.onWorking {
		_, msg, err := s.wssConn.ReadMessage()
		if err != nil {
			s.logger.Error("read message error: " + err.Error())
			return
		}
		dstObj := SparkRespMsg{}
		err1 := json.Unmarshal(msg, &dstObj)
		if err1 != nil {
			s.logger.Errorf("Error parsing JSON: %v", err)
			continue
		}
		s.acceptChan <- dstObj
	}
}

func (s *SparkWsConn) HandleStreamRespContent() {
	defer s.logger.Warn("AcceptChan closed, exit HandleStreamRespContent")
	s.logger.Info("Start HandleStreamRespContent Gor")
	for {
		select {
		case obj, ok := <-s.acceptChan:
			if !ok {
				s.logger.Warnf("acceptChan Had been closed, can not handle any more")
				return
			}
			// check code is ok ?
			if obj.RespHeader.Code != 0 {
				s.logger.Errorf("code not ok, Resp Header: %+v ", obj.RespHeader)
				continue
			}
			// check status
			switch obj.RespChoices.Status {
			case 0:
				// start
				s.SaveAssistantContentBySeq(fmt.Sprintf("%d-%s", obj.Seq, global.SparkAPI_Role_Assistant), obj.Text)
			case 1:
				// middle
				s.SaveAssistantContentBySeq(fmt.Sprintf("%d-%s", obj.Seq, global.SparkAPI_Role_Assistant), obj.Text)
			case 2:
				// end
				s.SaveAssistantContentBySeq(fmt.Sprintf("%d-%s", obj.Seq, global.SparkAPI_Role_Assistant), obj.Text)
				//TODO emit event
				// runtime.EventsEmit(s.ctx, "showSpark", s.AssembleStreamContent())
				//
				s.logger.Info("Get Completely Content: " + s.AssembleStreamContent())
			default:
				s.logger.Warnf("Unrecognizable Status code From Spark Server: %v", obj.RespChoices.Status)
			}

		default:

		}

	}
}
