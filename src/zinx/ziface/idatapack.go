package ziface

//IDataPack 处理TCP粘包
type IDataPack interface {
	GetHeadLen() uint32
	Pack(msg IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
