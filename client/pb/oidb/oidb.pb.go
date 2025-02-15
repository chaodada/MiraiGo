// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb.proto

package oidb

type OIDBSSOPkg struct {
	Command       int32  `protobuf:"varint,1,opt"`
	ServiceType   int32  `protobuf:"varint,2,opt"`
	Result        int32  `protobuf:"varint,3,opt"`
	Bodybuffer    []byte `protobuf:"bytes,4,opt"`
	ErrorMsg      string `protobuf:"bytes,5,opt"`
	ClientVersion string `protobuf:"bytes,6,opt"`
}

func (x *OIDBSSOPkg) GetCommand() int32 {
	if x != nil {
		return x.Command
	}
	return 0
}

func (x *OIDBSSOPkg) GetServiceType() int32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *OIDBSSOPkg) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *OIDBSSOPkg) GetBodybuffer() []byte {
	if x != nil {
		return x.Bodybuffer
	}
	return nil
}

func (x *OIDBSSOPkg) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

func (x *OIDBSSOPkg) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

type D8A0RspBody struct {
	OptUint64GroupCode int64             `protobuf:"varint,1,opt"`
	MsgKickResult      []*D8A0KickResult `protobuf:"bytes,2,rep"`
}

func (x *D8A0RspBody) GetOptUint64GroupCode() int64 {
	if x != nil {
		return x.OptUint64GroupCode
	}
	return 0
}

func (x *D8A0RspBody) GetMsgKickResult() []*D8A0KickResult {
	if x != nil {
		return x.MsgKickResult
	}
	return nil
}

type D8A0KickResult struct {
	OptUint32Result    int32 `protobuf:"varint,1,opt"`
	OptUint64MemberUin int64 `protobuf:"varint,2,opt"`
}

func (x *D8A0KickResult) GetOptUint32Result() int32 {
	if x != nil {
		return x.OptUint32Result
	}
	return 0
}

func (x *D8A0KickResult) GetOptUint64MemberUin() int64 {
	if x != nil {
		return x.OptUint64MemberUin
	}
	return 0
}

type D8A0KickMemberInfo struct {
	OptUint32Operate   int32  `protobuf:"varint,1,opt"`
	OptUint64MemberUin int64  `protobuf:"varint,2,opt"`
	OptUint32Flag      int32  `protobuf:"varint,3,opt"`
	OptBytesMsg        []byte `protobuf:"bytes,4,opt"`
}

func (x *D8A0KickMemberInfo) GetOptUint32Operate() int32 {
	if x != nil {
		return x.OptUint32Operate
	}
	return 0
}

func (x *D8A0KickMemberInfo) GetOptUint64MemberUin() int64 {
	if x != nil {
		return x.OptUint64MemberUin
	}
	return 0
}

func (x *D8A0KickMemberInfo) GetOptUint32Flag() int32 {
	if x != nil {
		return x.OptUint32Flag
	}
	return 0
}

func (x *D8A0KickMemberInfo) GetOptBytesMsg() []byte {
	if x != nil {
		return x.OptBytesMsg
	}
	return nil
}

type D8A0ReqBody struct {
	OptUint64GroupCode int64                 `protobuf:"varint,1,opt"`
	MsgKickList        []*D8A0KickMemberInfo `protobuf:"bytes,2,rep"`
	KickList           []int64               `protobuf:"varint,3,rep"`
	KickFlag           int32                 `protobuf:"varint,4,opt"`
	KickMsg            []byte                `protobuf:"bytes,5,opt"`
}

func (x *D8A0ReqBody) GetOptUint64GroupCode() int64 {
	if x != nil {
		return x.OptUint64GroupCode
	}
	return 0
}

func (x *D8A0ReqBody) GetMsgKickList() []*D8A0KickMemberInfo {
	if x != nil {
		return x.MsgKickList
	}
	return nil
}

func (x *D8A0ReqBody) GetKickList() []int64 {
	if x != nil {
		return x.KickList
	}
	return nil
}

func (x *D8A0ReqBody) GetKickFlag() int32 {
	if x != nil {
		return x.KickFlag
	}
	return 0
}

func (x *D8A0ReqBody) GetKickMsg() []byte {
	if x != nil {
		return x.KickMsg
	}
	return nil
}

type D89AReqBody struct {
	GroupCode           int64          `protobuf:"varint,1,opt"`
	StGroupInfo         *D89AGroupinfo `protobuf:"bytes,2,opt"`
	OriginalOperatorUin int64          `protobuf:"varint,3,opt"`
	ReqGroupOpenAppid   int32          `protobuf:"varint,4,opt"`
}

func (x *D89AReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *D89AReqBody) GetStGroupInfo() *D89AGroupinfo {
	if x != nil {
		return x.StGroupInfo
	}
	return nil
}

func (x *D89AReqBody) GetOriginalOperatorUin() int64 {
	if x != nil {
		return x.OriginalOperatorUin
	}
	return 0
}

func (x *D89AReqBody) GetReqGroupOpenAppid() int32 {
	if x != nil {
		return x.ReqGroupOpenAppid
	}
	return 0
}

type D89AGroupinfo struct {
	GroupExtAdmNum         int32                       `protobuf:"varint,1,opt"`
	Flag                   int32                       `protobuf:"varint,2,opt"`
	IngGroupName           []byte                      `protobuf:"bytes,3,opt"`
	IngGroupMemo           []byte                      `protobuf:"bytes,4,opt"`
	IngGroupFingerMemo     []byte                      `protobuf:"bytes,5,opt"`
	IngGroupAioSkinUrl     []byte                      `protobuf:"bytes,6,opt"`
	IngGroupBoardSkinUrl   []byte                      `protobuf:"bytes,7,opt"`
	IngGroupCoverSkinUrl   []byte                      `protobuf:"bytes,8,opt"`
	GroupGrade             int32                       `protobuf:"varint,9,opt"`
	ActiveMemberNum        int32                       `protobuf:"varint,10,opt"`
	CertificationType      int32                       `protobuf:"varint,11,opt"`
	IngCertificationText   []byte                      `protobuf:"bytes,12,opt"`
	IngGroupRichFingerMemo []byte                      `protobuf:"bytes,13,opt"`
	StGroupNewguidelines   *D89AGroupNewGuidelinesInfo `protobuf:"bytes,14,opt"`
	GroupFace              int32                       `protobuf:"varint,15,opt"`
	AddOption              int32                       `protobuf:"varint,16,opt"`
	// Types that are assignable to ShutupTime:
	//	*D89AGroupinfo_Val
	ShutupTime           isD89AGroupinfo_ShutupTime `protobuf_oneof:"shutupTime"`
	GroupTypeFlag        int32                      `protobuf:"varint,18,opt"`
	StringGroupTag       []byte                     `protobuf:"bytes,19,opt"`
	MsgGroupGeoInfo      *D89AGroupGeoInfo          `protobuf:"bytes,20,opt"`
	GroupClassExt        int32                      `protobuf:"varint,21,opt"`
	IngGroupClassText    []byte                     `protobuf:"bytes,22,opt"`
	AppPrivilegeFlag     int32                      `protobuf:"varint,23,opt"`
	AppPrivilegeMask     int32                      `protobuf:"varint,24,opt"`
	StGroupExInfo        *D89AGroupExInfoOnly       `protobuf:"bytes,25,opt"`
	GroupSecLevel        int32                      `protobuf:"varint,26,opt"`
	GroupSecLevelInfo    int32                      `protobuf:"varint,27,opt"`
	SubscriptionUin      int64                      `protobuf:"varint,28,opt"`
	AllowMemberInvite    int32                      `protobuf:"varint,29,opt"`
	IngGroupQuestion     []byte                     `protobuf:"bytes,30,opt"`
	IngGroupAnswer       []byte                     `protobuf:"bytes,31,opt"`
	GroupFlagext3        int32                      `protobuf:"varint,32,opt"`
	GroupFlagext3Mask    int32                      `protobuf:"varint,33,opt"`
	GroupOpenAppid       int32                      `protobuf:"varint,34,opt"`
	NoFingerOpenFlag     int32                      `protobuf:"varint,35,opt"`
	NoCodeFingerOpenFlag int32                      `protobuf:"varint,36,opt"`
	RootId               int64                      `protobuf:"varint,37,opt"`
	MsgLimitFrequency    int32                      `protobuf:"varint,38,opt"`
}

func (x *D89AGroupinfo) GetGroupExtAdmNum() int32 {
	if x != nil {
		return x.GroupExtAdmNum
	}
	return 0
}

func (x *D89AGroupinfo) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *D89AGroupinfo) GetIngGroupName() []byte {
	if x != nil {
		return x.IngGroupName
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupMemo() []byte {
	if x != nil {
		return x.IngGroupMemo
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupFingerMemo() []byte {
	if x != nil {
		return x.IngGroupFingerMemo
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupAioSkinUrl() []byte {
	if x != nil {
		return x.IngGroupAioSkinUrl
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupBoardSkinUrl() []byte {
	if x != nil {
		return x.IngGroupBoardSkinUrl
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupCoverSkinUrl() []byte {
	if x != nil {
		return x.IngGroupCoverSkinUrl
	}
	return nil
}

func (x *D89AGroupinfo) GetGroupGrade() int32 {
	if x != nil {
		return x.GroupGrade
	}
	return 0
}

func (x *D89AGroupinfo) GetActiveMemberNum() int32 {
	if x != nil {
		return x.ActiveMemberNum
	}
	return 0
}

func (x *D89AGroupinfo) GetCertificationType() int32 {
	if x != nil {
		return x.CertificationType
	}
	return 0
}

func (x *D89AGroupinfo) GetIngCertificationText() []byte {
	if x != nil {
		return x.IngCertificationText
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupRichFingerMemo() []byte {
	if x != nil {
		return x.IngGroupRichFingerMemo
	}
	return nil
}

func (x *D89AGroupinfo) GetStGroupNewguidelines() *D89AGroupNewGuidelinesInfo {
	if x != nil {
		return x.StGroupNewguidelines
	}
	return nil
}

func (x *D89AGroupinfo) GetGroupFace() int32 {
	if x != nil {
		return x.GroupFace
	}
	return 0
}

func (x *D89AGroupinfo) GetAddOption() int32 {
	if x != nil {
		return x.AddOption
	}
	return 0
}

func (m *D89AGroupinfo) GetShutupTime() isD89AGroupinfo_ShutupTime {
	if m != nil {
		return m.ShutupTime
	}
	return nil
}

func (x *D89AGroupinfo) GetVal() int32 {
	if x, ok := x.GetShutupTime().(*D89AGroupinfo_Val); ok {
		return x.Val
	}
	return 0
}

func (x *D89AGroupinfo) GetGroupTypeFlag() int32 {
	if x != nil {
		return x.GroupTypeFlag
	}
	return 0
}

func (x *D89AGroupinfo) GetStringGroupTag() []byte {
	if x != nil {
		return x.StringGroupTag
	}
	return nil
}

func (x *D89AGroupinfo) GetMsgGroupGeoInfo() *D89AGroupGeoInfo {
	if x != nil {
		return x.MsgGroupGeoInfo
	}
	return nil
}

func (x *D89AGroupinfo) GetGroupClassExt() int32 {
	if x != nil {
		return x.GroupClassExt
	}
	return 0
}

func (x *D89AGroupinfo) GetIngGroupClassText() []byte {
	if x != nil {
		return x.IngGroupClassText
	}
	return nil
}

func (x *D89AGroupinfo) GetAppPrivilegeFlag() int32 {
	if x != nil {
		return x.AppPrivilegeFlag
	}
	return 0
}

func (x *D89AGroupinfo) GetAppPrivilegeMask() int32 {
	if x != nil {
		return x.AppPrivilegeMask
	}
	return 0
}

func (x *D89AGroupinfo) GetStGroupExInfo() *D89AGroupExInfoOnly {
	if x != nil {
		return x.StGroupExInfo
	}
	return nil
}

func (x *D89AGroupinfo) GetGroupSecLevel() int32 {
	if x != nil {
		return x.GroupSecLevel
	}
	return 0
}

func (x *D89AGroupinfo) GetGroupSecLevelInfo() int32 {
	if x != nil {
		return x.GroupSecLevelInfo
	}
	return 0
}

func (x *D89AGroupinfo) GetSubscriptionUin() int64 {
	if x != nil {
		return x.SubscriptionUin
	}
	return 0
}

func (x *D89AGroupinfo) GetAllowMemberInvite() int32 {
	if x != nil {
		return x.AllowMemberInvite
	}
	return 0
}

func (x *D89AGroupinfo) GetIngGroupQuestion() []byte {
	if x != nil {
		return x.IngGroupQuestion
	}
	return nil
}

func (x *D89AGroupinfo) GetIngGroupAnswer() []byte {
	if x != nil {
		return x.IngGroupAnswer
	}
	return nil
}

func (x *D89AGroupinfo) GetGroupFlagext3() int32 {
	if x != nil {
		return x.GroupFlagext3
	}
	return 0
}

func (x *D89AGroupinfo) GetGroupFlagext3Mask() int32 {
	if x != nil {
		return x.GroupFlagext3Mask
	}
	return 0
}

func (x *D89AGroupinfo) GetGroupOpenAppid() int32 {
	if x != nil {
		return x.GroupOpenAppid
	}
	return 0
}

func (x *D89AGroupinfo) GetNoFingerOpenFlag() int32 {
	if x != nil {
		return x.NoFingerOpenFlag
	}
	return 0
}

func (x *D89AGroupinfo) GetNoCodeFingerOpenFlag() int32 {
	if x != nil {
		return x.NoCodeFingerOpenFlag
	}
	return 0
}

func (x *D89AGroupinfo) GetRootId() int64 {
	if x != nil {
		return x.RootId
	}
	return 0
}

func (x *D89AGroupinfo) GetMsgLimitFrequency() int32 {
	if x != nil {
		return x.MsgLimitFrequency
	}
	return 0
}

type isD89AGroupinfo_ShutupTime interface {
	isD89AGroupinfo_ShutupTime()
}

type D89AGroupinfo_Val struct {
	Val int32 `protobuf:"varint,17,opt"`
}

func (*D89AGroupinfo_Val) isD89AGroupinfo_ShutupTime() {}

type D89AGroupNewGuidelinesInfo struct {
	BoolEnabled bool   `protobuf:"varint,1,opt"`
	IngContent  []byte `protobuf:"bytes,2,opt"`
}

func (x *D89AGroupNewGuidelinesInfo) GetBoolEnabled() bool {
	if x != nil {
		return x.BoolEnabled
	}
	return false
}

func (x *D89AGroupNewGuidelinesInfo) GetIngContent() []byte {
	if x != nil {
		return x.IngContent
	}
	return nil
}

type D89AGroupExInfoOnly struct {
	TribeId          int32 `protobuf:"varint,1,opt"`
	MoneyForAddGroup int32 `protobuf:"varint,2,opt"`
}

func (x *D89AGroupExInfoOnly) GetTribeId() int32 {
	if x != nil {
		return x.TribeId
	}
	return 0
}

func (x *D89AGroupExInfoOnly) GetMoneyForAddGroup() int32 {
	if x != nil {
		return x.MoneyForAddGroup
	}
	return 0
}

type D89AGroupGeoInfo struct {
	CityId        int32  `protobuf:"varint,1,opt"`
	Longtitude    int64  `protobuf:"varint,2,opt"`
	Latitude      int64  `protobuf:"varint,3,opt"`
	IngGeoContent []byte `protobuf:"bytes,4,opt"`
	PoiId         int64  `protobuf:"varint,5,opt"`
}

func (x *D89AGroupGeoInfo) GetCityId() int32 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *D89AGroupGeoInfo) GetLongtitude() int64 {
	if x != nil {
		return x.Longtitude
	}
	return 0
}

func (x *D89AGroupGeoInfo) GetLatitude() int64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *D89AGroupGeoInfo) GetIngGeoContent() []byte {
	if x != nil {
		return x.IngGeoContent
	}
	return nil
}

func (x *D89AGroupGeoInfo) GetPoiId() int64 {
	if x != nil {
		return x.PoiId
	}
	return 0
}

type DED3ReqBody struct {
	ToUin     int64 `protobuf:"varint,1,opt"`
	GroupCode int64 `protobuf:"varint,2,opt"`
	MsgSeq    int32 `protobuf:"varint,3,opt"`
	MsgRand   int32 `protobuf:"varint,4,opt"`
	AioUin    int64 `protobuf:"varint,5,opt"`
}

func (x *DED3ReqBody) GetToUin() int64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *DED3ReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *DED3ReqBody) GetMsgSeq() int32 {
	if x != nil {
		return x.MsgSeq
	}
	return 0
}

func (x *DED3ReqBody) GetMsgRand() int32 {
	if x != nil {
		return x.MsgRand
	}
	return 0
}

func (x *DED3ReqBody) GetAioUin() int64 {
	if x != nil {
		return x.AioUin
	}
	return 0
}
