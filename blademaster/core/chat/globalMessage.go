package chat

import (
	"net"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/kerlong"
	. "github.com/KouKouChan/CSO2-Server/servermanager"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

func OnChatGlobalMessage(p *InChatPacket, client net.Conn) {
	//找到对应用户
	uPtr := GetUserFromConnection(client)
	if uPtr == nil ||
		uPtr.Userid <= 0 {
		DebugInfo(2, "Error : Client from", client.RemoteAddr().String(), "sent GlobalMessage but not in server !")
		return
	}
	if uPtr.GetUserRoomID() <= 0 {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent GlobalMessage but not in room !")
		return
	}
	//找到对应房间
	rm := GetRoomFromID(uPtr.GetUserChannelServerID(), uPtr.GetUserChannelID(), uPtr.GetUserRoomID())
	if rm == nil || rm.Id <= 0 {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent GlobalMessage but not in room !")
		return
	}
	if !uPtr.CurrentIsIngame {
		DebugInfo(2, "Error : User", string(uPtr.IngameName), "sent GlobalMessage but not ingame !")
		return
	}
	//发送数据
	msg := BuildChatMessage(uPtr, p, ChatIngameGlobal)
	for _, v := range rm.Users {
		if v.CurrentIsIngame {
			SendPacket(BytesCombine(BuildHeader(v.CurrentSequence, PacketTypeChat), msg), v.CurrentConnection)
		}
	}
	DebugInfo(1, "User", string(uPtr.IngameName), "say global message <", string(p.Message), "> in room id", rm.Id)
}
