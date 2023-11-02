package ynet

import (
	"Yis/utils"
	"Yis/yiface"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type DataPack struct {
}

func NewDataPack() yiface.IDataPack {
	pack := &DataPack{}
	return pack
}

// GetHeadLen 获取消息头长度
func (dp *DataPack) GetHeadLen() uint32 {
	//id:uint32，length:uint32
	return 8
}

// Pack 封包方法
func (dp *DataPack) Pack(msg yiface.IMessage) ([]byte, error) {
	//创建一个byte缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写id
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetId())
	if err != nil {
		fmt.Println("binary.Write error : ", err)
		return nil, err
	}
	//写length
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetLength())
	if err != nil {
		fmt.Println("binary.Write2 error : ", err)
		return nil, err
	}
	//写data
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		fmt.Println("binary.Write3 error : ", err)
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// UnPack 拆包方法，只解析出消息头部
func (dp *DataPack) UnPack(data []byte) (yiface.IMessage, error) {
	dataBuff := bytes.NewReader(data)

	msg := &Message{}
	//取出id
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		fmt.Println("binary Read error : ", err)
		return nil, err
	}
	//取出length
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Length)
	if err != nil {
		fmt.Println("binary Read2 error : ", err)
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.GetLength() > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg size")
	}

	return msg, nil
}
