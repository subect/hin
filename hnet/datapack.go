package hnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hin/hiface"
	"hin/utils"
)

// DataPack 封包、拆包模块
// 直接面向TCP连接中的数据流，用于处理TCP粘包问题
type DataPack struct{}

// NewDataPack 封包拆包实例的初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32(4字节) + DataLen uint32(4字节)
	return 8
}

// Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg hiface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	// 将dataLen写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// 将msgId写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 将data数据写进dataBuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (hiface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	// 只解压head信息，得到dataLen和msgId
	msg := &Message{}

	// 读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读msgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断dataLen是否已经超出了我们允许的最大包长度
	if utils.GlobalObject.MaxPackage > 0 && msg.DataLen > utils.GlobalObject.MaxPackage {
		return nil, errors.New("too large msg data received")
	}

	return msg, nil
}
