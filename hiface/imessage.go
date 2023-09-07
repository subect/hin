package hiface

// 将请求的消息封装到 message 中

type IMessage interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgId() uint32   //获取消息id
	GetData() []byte    //获取消息内容
	SetMsgId(uint32)    //设计消息id
	SetData([]byte)     //设计消息内容
	SetDataLen(uint32)  //设置消息数据段长度
}
