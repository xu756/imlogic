// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package im

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *MediaType) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MediaType[number], err)
}

func (x *MediaType) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Uid, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MediaType) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Url, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MetaMsg) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MetaMsg[number], err)
}

func (x *MetaMsg) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LinkId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MetaMsg) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.UserId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MetaMsg) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Status = WsStatus(v)
	return offset, nil
}

func (x *MetaMsg) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.HostName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MetaMsg) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Device, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MessageRes) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MessageRes[number], err)
}

func (x *MessageRes) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *MessageRes) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.MsgId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Message[number], err)
}

func (x *Message) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LinkId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.MsgId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Timestamp, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.ChatType = ChatType(v)
	return offset, nil
}

func (x *Message) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Sender, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Receiver, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.ChatId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Message) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Message) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.MsgType = MsgType(v)
	return offset, nil
}

func (x *Message) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	var v MediaType
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Media = append(x.Media, &v)
	return offset, nil
}

func (x *SendMsgTooneReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMsgTooneReq[number], err)
}

func (x *SendMsgTooneReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LinkId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMsgTooneReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Message = &v
	return offset, nil
}

func (x *SendMsgToGroupReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMsgToGroupReq[number], err)
}

func (x *SendMsgToGroupReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.LinkIds = append(x.LinkIds, v)
	return offset, err
}

func (x *SendMsgToGroupReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Message
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Message = &v
	return offset, nil
}

func (x *SendMsgToGroupRes) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMsgToGroupRes[number], err)
}

func (x *SendMsgToGroupRes) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.LinkIds = append(x.LinkIds, v)
	return offset, err
}

func (x *MediaType) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MediaType) fastWriteField1(buf []byte) (offset int) {
	if x.Uid == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUid())
	return offset
}

func (x *MediaType) fastWriteField2(buf []byte) (offset int) {
	if x.Url == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUrl())
	return offset
}

func (x *MetaMsg) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *MetaMsg) fastWriteField1(buf []byte) (offset int) {
	if x.LinkId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLinkId())
	return offset
}

func (x *MetaMsg) fastWriteField2(buf []byte) (offset int) {
	if x.UserId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetUserId())
	return offset
}

func (x *MetaMsg) fastWriteField3(buf []byte) (offset int) {
	if x.Status == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, int32(x.GetStatus()))
	return offset
}

func (x *MetaMsg) fastWriteField4(buf []byte) (offset int) {
	if x.HostName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetHostName())
	return offset
}

func (x *MetaMsg) fastWriteField5(buf []byte) (offset int) {
	if x.Device == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetDevice())
	return offset
}

func (x *MessageRes) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *MessageRes) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *MessageRes) fastWriteField2(buf []byte) (offset int) {
	if x.MsgId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsgId())
	return offset
}

func (x *Message) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *Message) fastWriteField1(buf []byte) (offset int) {
	if x.LinkId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLinkId())
	return offset
}

func (x *Message) fastWriteField2(buf []byte) (offset int) {
	if x.MsgId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetMsgId())
	return offset
}

func (x *Message) fastWriteField3(buf []byte) (offset int) {
	if x.Timestamp == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetTimestamp())
	return offset
}

func (x *Message) fastWriteField4(buf []byte) (offset int) {
	if x.ChatType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, int32(x.GetChatType()))
	return offset
}

func (x *Message) fastWriteField5(buf []byte) (offset int) {
	if x.Sender == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetSender())
	return offset
}

func (x *Message) fastWriteField6(buf []byte) (offset int) {
	if x.Receiver == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.GetReceiver())
	return offset
}

func (x *Message) fastWriteField7(buf []byte) (offset int) {
	if x.ChatId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 7, x.GetChatId())
	return offset
}

func (x *Message) fastWriteField8(buf []byte) (offset int) {
	if x.GroupId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetGroupId())
	return offset
}

func (x *Message) fastWriteField9(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 9, x.GetContent())
	return offset
}

func (x *Message) fastWriteField10(buf []byte) (offset int) {
	if x.MsgType == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 10, int32(x.GetMsgType()))
	return offset
}

func (x *Message) fastWriteField11(buf []byte) (offset int) {
	if x.Media == nil {
		return offset
	}
	for i := range x.GetMedia() {
		offset += fastpb.WriteMessage(buf[offset:], 11, x.GetMedia()[i])
	}
	return offset
}

func (x *SendMsgTooneReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SendMsgTooneReq) fastWriteField1(buf []byte) (offset int) {
	if x.LinkId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetLinkId())
	return offset
}

func (x *SendMsgTooneReq) fastWriteField2(buf []byte) (offset int) {
	if x.Message == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *SendMsgToGroupReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SendMsgToGroupReq) fastWriteField1(buf []byte) (offset int) {
	if len(x.LinkIds) == 0 {
		return offset
	}
	for i := range x.GetLinkIds() {
		offset += fastpb.WriteString(buf[offset:], 1, x.GetLinkIds()[i])
	}
	return offset
}

func (x *SendMsgToGroupReq) fastWriteField2(buf []byte) (offset int) {
	if x.Message == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetMessage())
	return offset
}

func (x *SendMsgToGroupRes) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SendMsgToGroupRes) fastWriteField1(buf []byte) (offset int) {
	if len(x.LinkIds) == 0 {
		return offset
	}
	for i := range x.GetLinkIds() {
		offset += fastpb.WriteString(buf[offset:], 1, x.GetLinkIds()[i])
	}
	return offset
}

func (x *MediaType) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MediaType) sizeField1() (n int) {
	if x.Uid == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUid())
	return n
}

func (x *MediaType) sizeField2() (n int) {
	if x.Url == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUrl())
	return n
}

func (x *MetaMsg) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *MetaMsg) sizeField1() (n int) {
	if x.LinkId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetLinkId())
	return n
}

func (x *MetaMsg) sizeField2() (n int) {
	if x.UserId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetUserId())
	return n
}

func (x *MetaMsg) sizeField3() (n int) {
	if x.Status == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, int32(x.GetStatus()))
	return n
}

func (x *MetaMsg) sizeField4() (n int) {
	if x.HostName == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetHostName())
	return n
}

func (x *MetaMsg) sizeField5() (n int) {
	if x.Device == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetDevice())
	return n
}

func (x *MessageRes) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *MessageRes) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *MessageRes) sizeField2() (n int) {
	if x.MsgId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsgId())
	return n
}

func (x *Message) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *Message) sizeField1() (n int) {
	if x.LinkId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetLinkId())
	return n
}

func (x *Message) sizeField2() (n int) {
	if x.MsgId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetMsgId())
	return n
}

func (x *Message) sizeField3() (n int) {
	if x.Timestamp == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetTimestamp())
	return n
}

func (x *Message) sizeField4() (n int) {
	if x.ChatType == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, int32(x.GetChatType()))
	return n
}

func (x *Message) sizeField5() (n int) {
	if x.Sender == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetSender())
	return n
}

func (x *Message) sizeField6() (n int) {
	if x.Receiver == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.GetReceiver())
	return n
}

func (x *Message) sizeField7() (n int) {
	if x.ChatId == 0 {
		return n
	}
	n += fastpb.SizeInt64(7, x.GetChatId())
	return n
}

func (x *Message) sizeField8() (n int) {
	if x.GroupId == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetGroupId())
	return n
}

func (x *Message) sizeField9() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(9, x.GetContent())
	return n
}

func (x *Message) sizeField10() (n int) {
	if x.MsgType == 0 {
		return n
	}
	n += fastpb.SizeInt32(10, int32(x.GetMsgType()))
	return n
}

func (x *Message) sizeField11() (n int) {
	if x.Media == nil {
		return n
	}
	for i := range x.GetMedia() {
		n += fastpb.SizeMessage(11, x.GetMedia()[i])
	}
	return n
}

func (x *SendMsgTooneReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SendMsgTooneReq) sizeField1() (n int) {
	if x.LinkId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetLinkId())
	return n
}

func (x *SendMsgTooneReq) sizeField2() (n int) {
	if x.Message == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetMessage())
	return n
}

func (x *SendMsgToGroupReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SendMsgToGroupReq) sizeField1() (n int) {
	if len(x.LinkIds) == 0 {
		return n
	}
	for i := range x.GetLinkIds() {
		n += fastpb.SizeString(1, x.GetLinkIds()[i])
	}
	return n
}

func (x *SendMsgToGroupReq) sizeField2() (n int) {
	if x.Message == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetMessage())
	return n
}

func (x *SendMsgToGroupRes) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SendMsgToGroupRes) sizeField1() (n int) {
	if len(x.LinkIds) == 0 {
		return n
	}
	for i := range x.GetLinkIds() {
		n += fastpb.SizeString(1, x.GetLinkIds()[i])
	}
	return n
}

var fieldIDToName_MediaType = map[int32]string{
	1: "Uid",
	2: "Url",
}

var fieldIDToName_MetaMsg = map[int32]string{
	1: "LinkId",
	2: "UserId",
	3: "Status",
	4: "HostName",
	5: "Device",
}

var fieldIDToName_MessageRes = map[int32]string{
	1: "Success",
	2: "MsgId",
}

var fieldIDToName_Message = map[int32]string{
	1:  "LinkId",
	2:  "MsgId",
	3:  "Timestamp",
	4:  "ChatType",
	5:  "Sender",
	6:  "Receiver",
	7:  "ChatId",
	8:  "GroupId",
	9:  "Content",
	10: "MsgType",
	11: "Media",
}

var fieldIDToName_SendMsgTooneReq = map[int32]string{
	1: "LinkId",
	2: "Message",
}

var fieldIDToName_SendMsgToGroupReq = map[int32]string{
	1: "LinkIds",
	2: "Message",
}

var fieldIDToName_SendMsgToGroupRes = map[int32]string{
	1: "LinkIds",
}
