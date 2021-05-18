package services

import (
	"gin-test/model"
	"gin-test/model/date"
	"go.uber.org/zap"
	"sync"
)

var MessageService = newMessageService()

var messageLog = zap.L()

func newMessageService() *messageService {
	return &messageService{
		messagesChan: make(chan *model.Message),
	}
}

type messageService struct {
	messagesChan        chan *model.Message
	messagesConsumeOnce sync.Once
}

// 生产，将消息数据放入chan
func (s *messageService) Produce(fromId, toId int64, title, content, quoteContent string, msgType int, extraDataMap map[string]interface{}) {
	s.Consume()

	//to := cache.UserCache.Get(toId)
	//if to == nil || to.Type != constants.UserTypeNormal {
	//	return
	//}

	var (
		extraData string
		//err       error
	)
	//if extraData, err = json.ToStr(extraDataMap); err != nil {
	//	messageLog.Error("格式化extraData错误", err)
	//}
	s.messagesChan <- &model.Message{
		FromId:       fromId,
		UserId:       toId,
		Title:        title,
		Content:      content,
		QuoteContent: quoteContent,
		Type:         msgType,
		ExtraData:    extraData,
		Status:       1,
		CreateTime:   date.NowTimestamp(),
	}
}

// 消费，消费chan中的消息
func (s *messageService) Consume() {
	s.messagesConsumeOnce.Do(func() {
		go func() {
			messageLog.Info("开始消费系统消息...")
			for {
				msg := <-s.messagesChan
				messageLog.Info("处理消息", zap.Int64("FromId", msg.FromId), zap.Int64("toId:", msg.UserId))

				//if err := s.Create(msg); err != nil {
				//	messageLog.Info("创建消息发生异常...", err)
				//} else {
				//	s.SendEmailNotice(msg)
				//}
			}
		}()
	})
}
