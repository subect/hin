package hnet

import (
	"fmt"
	"hin/hiface"
)

type MsgHandler struct {
	// 存放每个MsgID所对应的处理方法
	Apis map[uint32]hiface.IRouter
}

// NewMsgHandler 创建MsgHandler
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]hiface.IRouter),
	}
}

// DoMsgHandler 以非阻塞方式处理消息
func (mh *MsgHandler) DoMsgHandler(request hiface.IRequest) {
	// 1 从 Request 中找到 msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		return
	}
	// 2 根据 MsgID 调度对应 router 业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandler) AddRouter(msgId uint32, router hiface.IRouter) {
	// 1 判断当前 msg 绑定的 API 处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeat api, msgId = " + string(msgId))
	}
	// 2 添加 msg 与 api 的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}
