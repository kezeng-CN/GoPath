package znet

//Message 协议结构
type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

//NewMsgPackage 返回Message包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//GetDataLen 数据处理
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//GetMsgID 数据处理
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

//GetData 数据处理
func (msg *Message) GetData() []byte {
	return msg.Data
}

//SetDataLen 数据处理
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//SetMsgID 数据处理
func (msg *Message) SetMsgID(msgid uint32) {
	msg.ID = msgid
}

//SetData 数据处理
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
