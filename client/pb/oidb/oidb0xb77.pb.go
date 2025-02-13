// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb0xb77.proto

package oidb

type DB77ReqBody struct {
	AppId       uint64           `protobuf:"varint,1,opt"`
	AppType     uint32           `protobuf:"varint,2,opt"`
	MsgStyle    uint32           `protobuf:"varint,3,opt"`
	SenderUin   uint64           `protobuf:"varint,4,opt"`
	ClientInfo  *DB77ClientInfo  `protobuf:"bytes,5,opt"`
	TextMsg     string           `protobuf:"bytes,6,opt"`
	ExtInfo     *DB77ExtInfo     `protobuf:"bytes,7,opt"`
	SendType    uint32           `protobuf:"varint,10,opt"`
	RecvUin     uint64           `protobuf:"varint,11,opt"`
	RichMsgBody *DB77RichMsgBody `protobuf:"bytes,12,opt"`
}

func (x *DB77ReqBody) GetAppId() uint64 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DB77ReqBody) GetAppType() uint32 {
	if x != nil {
		return x.AppType
	}
	return 0
}

func (x *DB77ReqBody) GetMsgStyle() uint32 {
	if x != nil {
		return x.MsgStyle
	}
	return 0
}

func (x *DB77ReqBody) GetSenderUin() uint64 {
	if x != nil {
		return x.SenderUin
	}
	return 0
}

func (x *DB77ReqBody) GetClientInfo() *DB77ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *DB77ReqBody) GetTextMsg() string {
	if x != nil {
		return x.TextMsg
	}
	return ""
}

func (x *DB77ReqBody) GetExtInfo() *DB77ExtInfo {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *DB77ReqBody) GetSendType() uint32 {
	if x != nil {
		return x.SendType
	}
	return 0
}

func (x *DB77ReqBody) GetRecvUin() uint64 {
	if x != nil {
		return x.RecvUin
	}
	return 0
}

func (x *DB77ReqBody) GetRichMsgBody() *DB77RichMsgBody {
	if x != nil {
		return x.RichMsgBody
	}
	return nil
}

type DB77ClientInfo struct {
	Platform           uint32 `protobuf:"varint,1,opt"`
	SdkVersion         string `protobuf:"bytes,2,opt"`
	AndroidPackageName string `protobuf:"bytes,3,opt"`
	AndroidSignature   string `protobuf:"bytes,4,opt"`
	IosBundleId        string `protobuf:"bytes,5,opt"`
	PcSign             string `protobuf:"bytes,6,opt"`
}

func (x *DB77ClientInfo) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *DB77ClientInfo) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

func (x *DB77ClientInfo) GetAndroidPackageName() string {
	if x != nil {
		return x.AndroidPackageName
	}
	return ""
}

func (x *DB77ClientInfo) GetAndroidSignature() string {
	if x != nil {
		return x.AndroidSignature
	}
	return ""
}

func (x *DB77ClientInfo) GetIosBundleId() string {
	if x != nil {
		return x.IosBundleId
	}
	return ""
}

func (x *DB77ClientInfo) GetPcSign() string {
	if x != nil {
		return x.PcSign
	}
	return ""
}

type DB77ExtInfo struct {
	CustomFeatureId []uint32 `protobuf:"varint,11,rep"`
	ApnsWording     string   `protobuf:"bytes,12,opt"`
	GroupSaveDbFlag uint32   `protobuf:"varint,13,opt"`
	ReceiverAppId   uint32   `protobuf:"varint,14,opt"`
	MsgSeq          uint64   `protobuf:"varint,15,opt"`
}

func (x *DB77ExtInfo) GetCustomFeatureId() []uint32 {
	if x != nil {
		return x.CustomFeatureId
	}
	return nil
}

func (x *DB77ExtInfo) GetApnsWording() string {
	if x != nil {
		return x.ApnsWording
	}
	return ""
}

func (x *DB77ExtInfo) GetGroupSaveDbFlag() uint32 {
	if x != nil {
		return x.GroupSaveDbFlag
	}
	return 0
}

func (x *DB77ExtInfo) GetReceiverAppId() uint32 {
	if x != nil {
		return x.ReceiverAppId
	}
	return 0
}

func (x *DB77ExtInfo) GetMsgSeq() uint64 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

type DB77RichMsgBody struct {
	Title      string `protobuf:"bytes,10,opt"`
	Summary    string `protobuf:"bytes,11,opt"`
	Brief      string `protobuf:"bytes,12,opt"`
	Url        string `protobuf:"bytes,13,opt"`
	PictureUrl string `protobuf:"bytes,14,opt"`
	Action     string `protobuf:"bytes,15,opt"`
	MusicUrl   string `protobuf:"bytes,16,opt"` //ImageInfo imageInfo = 17;
}

func (x *DB77RichMsgBody) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DB77RichMsgBody) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *DB77RichMsgBody) GetBrief() string {
	if x != nil {
		return x.Brief
	}
	return ""
}

func (x *DB77RichMsgBody) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DB77RichMsgBody) GetPictureUrl() string {
	if x != nil {
		return x.PictureUrl
	}
	return ""
}

func (x *DB77RichMsgBody) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *DB77RichMsgBody) GetMusicUrl() string {
	if x != nil {
		return x.MusicUrl
	}
	return ""
}
