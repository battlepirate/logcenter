package main

import ()

//总日志结构
type LogInfo struct {
	ServerId uint64 //服务器ID   TODO 前三个字段不能动，应用模块已经使用
	OpId     uint32 //运营商ID
	UserId   uint64 //玩家ID

	EventId   uint64 //流水ID
	MainType  uint32 //日志主类型
	ChildType uint32 //日志子类型

	RealServerId uint64 //真实服务器ID（合服后主服务器ID）
	OpgameId     uint32 //混服组ID
	AdId         string //广告ID
	ChannelId    string //用户来源 CPS ID
	Account      string //平台账号
	DeviceId     string //设备号
	LogSendTime  uint32 //日志发送时间

	ClientIp        string //客户端IP
	Uuid            string //玩家UUID
	UserName        string //玩家名字
	UserLevel       uint32 //玩家等级
	RegisterTime    uint32 //玩家注册时间
	PurchaseDiamond uint32 //购买钻石数
	DonateDiamond   uint32 //玩家赠送钻石数
	CorpId          uint64 //军团ID
	RolePaid        bool   //是否为付费玩家
	RoleValid       bool   //是否为有效玩家
	VipLevel        uint32 //VIP等级
	RoleExp         uint32 //玩家经验
	RoleEnergy      uint32 //玩家体力
	OnlineTime      uint32 //本次上线时间点
	OfflineTime     uint32 //上次下线时间点
	LastOnlineTime  uint32 //上次上线时间点

	LogParams
	Nanosec uint64 //TODO 最后字段不能动，应用模块已经使用
}

type LogParams struct {
	Param1  uint32
	Param2  uint32
	Param3  uint32
	Param4  uint32
	Param5  uint32
	Param6  uint32
	Param7  uint32
	Param8  uint32
	Param9  uint32
	Param10 uint32
	Param11 uint64
	Param12 uint64
	Param13 uint64
	Param14 uint64
	Param15 uint64
	Param16 string
	Param17 string
	Param18 string
	Param19 string
	Param20 string
}

//登陆日志
type LogLogin struct {
	Time          uint32 `json:"time"`          //时间戳
	Type          uint32 `json:"type"`          //登陆类型 1: 通过角色ID 2：通过角色名称 3：通过平台账号
	UserName      string `json:"userName"`      //登陆用户
	UserId        string `json:"userId"`        //账号ID
	UserType      string `json:"userType"`      //账号类型
	AreaName      uint32 `json:"areaName"`      //登陆区名称
	UserLv        uint32 `json:"userLev"`       //用户等级
	VipLv         uint32 `json:"vipLev"`        //VIP等级
	LeaderRank    string `json:"leaderRank"`    //排名
	Stage         string `json:"stage"`         //关卡位置
	DailyStage    string `json:"eliteStage"`    //日常关卡位置
	DeviceId      string `json:"deviceId"`      //设备ID
	DeviceMobile  string `json:"deviceMobile"`  //机型
	ClientIp      string `json:"clientIp"`      //终端IP
	DeviceType    string `json:"deviceType"`    //应用类别（1是安卓，2是IOS）
	ChannelName   string `json:"channelName"`   //渠道ID
	DeviceCarrier string `json:"deviceCarrier"` //运营商
	DeviceOsVer   string `json:"deviceOsVer"`   //设备具体系统
	IsNew         uint32 `json:"isNew"`         //账号为新，为1，否则为0
	IsAreaNew     uint32 `json:"isAreaNew"`     //角色微信，为1，否则为0
}

//创角注册日志
type LogRegister struct {
	Time          uint32 `json:"time"`          //时间戳
	UserName      string `json:"userName"`      //登陆用户
	UserId        string `json:"userId"`        //账号ID
	UserType      string `json:"userType"`      //账号类型
	DeviceId      string `json:"deviceId"`      //设备ID
	DeviceMobile  string `json:"deviceMobile"`  //机型
	ClientIp      string `json:"clientIp"`      //终端IP
	DeviceType    string `json:"deviceType"`    //应用类别（1是安卓，2是IOS）
	ChannelName   string `json:"channelName"`   //渠道ID
	DeviceCarrier string `json:"deviceCarrier"` //运营商
	DeviceOsVer   string `json:"deviceOsVer"`   //设备具体系统
}

//物品产出日志
type LogOutPut struct {
	Time          uint32 `json:"time"`          //时间戳
	Type          string `json:"type"`          //产出来源
	ItemId        string `json:"itemId"`        //道具ID
	ItemType      uint32 `json:"itemType"`      //道具类型
	ItemName      string `json:"itemName"`      //道具名称
	ItemNum       uint32 `json:"itemNum"`       //道具数量
	UserName      string `json:"userName"`      //登陆用户
	UserId        string `json:"userId"`        //账号ID
	UserType      string `json:"userType"`      //账号类型
	AreaName      uint32 `json:"areaName"`      //登陆区名称
	UserLv        uint32 `json:"userLev"`       //用户等级
	VipLv         uint32 `json:"vipLev"`        //VIP等级
	DeviceId      string `json:"deviceId"`      //设备ID
	DeviceMobile  string `json:"deviceMobile"`  //机型
	ClientIp      string `json:"clientIp"`      //终端IP
	DeviceType    string `json:"deviceType"`    //应用类别（1是安卓，2是IOS）
	ChannelName   string `json:"channelName"`   //渠道ID
	DeviceCarrier string `json:"deviceCarrier"` //运营商
	DeviceOsVer   string `json:"deviceOsVer"`   //设备具体系统
}

//在线人数日志
type LogUserOnline struct {
	Time          uint32 `json:"time"`          //时间戳
	AreaName      uint32 `json:"areaName"`      //登陆区名称
	Count         uint32 `json:"count"`         //在线玩家数量
	DeviceId      string `json:"deviceId"`      //设备ID
	DeviceMobile  string `json:"deviceMobile"`  //机型
	ClientIp      string `json:"clientIp"`      //终端IP
	DeviceType    string `json:"deviceType"`    //应用类别（1是安卓，2是IOS）
	ChannelName   string `json:"channelName"`   //渠道ID
	DeviceCarrier string `json:"deviceCarrier"` //运营商
	DeviceOsVer   string `json:"deviceOsVer"`   //设备具体系统
}

//物品消耗日志
type LogItemConsume struct {
	Time          uint32 `json:"time"`          //时间戳
	Type          string `json:"type"`          //产出来源
	ItemId        string `json:"itemId"`        //道具ID
	ItemType      uint32 `json:"itemType"`      //道具类型
	ItemName      string `json:"itemName"`      //道具名称
	ItemNum       uint32 `json:"itemNum"`       //道具数量
	UserName      string `json:"userName"`      //登陆用户
	UserId        string `json:"userId"`        //账号ID
	UserType      string `json:"userType"`      //账号类型
	AreaName      uint32 `json:"areaName"`      //登陆区名称
	UserLv        uint32 `json:"userLev"`       //用户等级
	VipLv         uint32 `json:"vipLev"`        //VIP等级
	DeviceId      string `json:"deviceId"`      //设备ID
	DeviceMobile  string `json:"deviceMobile"`  //机型
	ClientIp      string `json:"clientIp"`      //终端IP
	DeviceType    string `json:"deviceType"`    //应用类别（1是安卓，2是IOS）
	ChannelName   string `json:"channelName"`   //渠道ID
	DeviceCarrier string `json:"deviceCarrier"` //运营商
	DeviceOsVer   string `json:"deviceOsVer"`   //设备具体系统
}

//关卡事件
type LogChapterStage struct {
	Time         uint32 `json:"time"`         //时间戳
	UserName     string `json:"userName"`     //登陆用户
	UserId       string `json:"userId"`       //账号ID
	UserType     string `json:"userType"`     //账号类型
	AreaName     uint32 `json:"areaName"`     //登陆区名称
	UserLv       uint32 `json:"userLev"`      //用户等级
	VipLv        uint32 `json:"vipLev"`       //VIP等级
	Efficiency   uint32 `json:"efficiency"`   //关卡ID
	StageType    string `json:"stageType"`    //日常，主线
	Stage        string `json:"stage"`        //关卡名称
	IsSuccess    uint32 `json:"isSuccess"`    //0：成功，1，失败，2，开始
	DeviceId     string `json:"deviceId"`     //设备ID
	DeviceMobile string `json:"deviceMobile"` //机型
	ClientIp     string `json:"clientIp"`     //终端IP
	DeviceType   string `json:"deviceType"`   //应用类别（1是安卓，2是IOS）
	ChannelName  string `json:"channelName"`  //渠道ID
	AppVersion   string `json:"appVersion"`   //版本号
}

//在线时长
type LogOnlineTime struct {
	Time         uint32 `json:"time"`         //时间戳
	UserName     string `json:"userName"`     //登陆用户
	UserId       string `json:"userId"`       //账号ID
	UserType     string `json:"userType"`     //账号类型
	AreaName     uint32 `json:"areaName"`     //登陆区名称
	Duration     uint32 `json:"duration"`     //用户在线时间(秒)
	DeviceId     string `json:"deviceId"`     //设备ID
	DeviceMobile string `json:"deviceMobile"` //机型
	ClientIp     string `json:"clientIp"`     //终端IP
	DeviceType   string `json:"deviceType"`   //应用类别（1是安卓，2是IOS）
	ChannelName  string `json:"channelName"`  //渠道ID
	AppVersion   string `json:"appVersion"`   //版本号
}

//等级
type LogLevel struct {
	Time         uint32 `json:"time"`         //时间戳
	UserName     string `json:"userName"`     //登陆用户
	UserId       string `json:"userId"`       //账号ID
	UserType     string `json:"userType"`     //账号类型
	AreaName     uint32 `json:"areaName"`     //登陆区名称
	UserLv       uint32 `json:"userLev"`      //用户等级
	VipLv        uint32 `json:"vipLev"`       //VIP等级
	DeviceId     string `json:"deviceId"`     //设备ID
	DeviceMobile string `json:"deviceMobile"` //机型
	ClientIp     string `json:"clientIp"`     //终端IP
	DeviceType   string `json:"deviceType"`   //应用类别（1是安卓，2是IOS）
	ChannelName  string `json:"channelName"`  //渠道ID
	AppVersion   string `json:"appVersion"`   //版本号
}

//充值货币进出
type LogSilverConsume struct {
	Time       uint32 `json:"time"`       //时间戳
	UserName   string `json:"userName"`   //登陆用户
	UserId     string `json:"userId"`     //账号ID
	UserType   string `json:"userType"`   //账号类型
	AreaName   uint32 `json:"areaName"`   //登陆区名称
	UserLv     uint32 `json:"userLev"`    //用户等级
	VipLv      uint32 `json:"vipLev"`     //VIP等级
	LeaderRank string `json:"leaderRank"` //排名
	Stage      string `json:"stage"`      //关卡位置
	DailyStage string `json:"eliteStage"` //日常关卡位置

	ConsumeType string `json:"consumeType"` //消耗类型 用于干什么了
	Consume     int    `json:"consume"`     //消耗数量(产出拋正值,消耗拋负值)
	ItemId      string `json:"itemId"`      //道具ID
	ItemType    uint32 `json:"itemType"`    //道具类型
	ItemName    string `json:"itemName"`    //道具名称
	ItemNum     uint32 `json:"itemNum"`     //道具数量

	DeviceId     string `json:"deviceId"`     //设备ID
	DeviceMobile string `json:"deviceMobile"` //机型
	ClientIp     string `json:"clientIp"`     //终端IP
	DeviceType   string `json:"deviceType"`   //应用类别（1是安卓，2是IOS）
	ChannelName  string `json:"channelName"`  //渠道ID

}
