package wol

import (
	"net"
)

// 网络唤醒
func Wol(macStr string) error {
	macAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return err
	}
	wolMsg := wolMsgCreate(macAddr)
	err = updBroadcast(wolMsg)
	return err
}

// 构造WOL数据包
func wolMsgCreate(macAddr net.HardwareAddr) (wolMsg []byte) {
	wolMsg = make([]byte, 0, 102)
	wolMsg = append(wolMsg, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF)
	for i := 0; i < 16; i++ {
		wolMsg = append(wolMsg, macAddr...)
	}
	return wolMsg
}

// UDP本地网络广播
func updBroadcast(msg []byte) error {
	conn, err := net.Dial("udp", "255.255.255.255:0")
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Write(msg)
	return err
}
