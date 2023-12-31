package hnet

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// NewMsgPackage 创建一个 Message 消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// GetMsgId 获取消息的 ID
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// GetDataLen GetDateLen 获取消息的长度
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// GetData 获取消息的内容
func (m *Message) GetData() []byte {
	return m.Data
}

// SetMsgId 设计消息的 ID
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

// SetDataLen 设计消息的长度
func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

// SetData 设计消息的内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}
