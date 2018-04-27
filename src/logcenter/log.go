package main

import (
	"strconv"
)

func LogParamsToLogInfo(logParams *LogParams, logInfo *LogInfo) {
	if logParams != nil {
		logInfo.LogParams = *logParams
	}
}

/*-- 总日志转换成类型日志--*/
//转换成登陆日志
func ChangeLogToLogin(info LogInfo) LogLogin {
	return LogLogin{
		Time:          info.OnlineTime,
		Type:          info.Param3,
		UserName:      info.Uuid,
		UserId:        strconv.FormatUint(info.UserId, 10),
		UserType:      "",
		AreaName:      uint32(0),                                   //登陆区名称
		UserLv:        info.UserLevel,                              //用户等级
		VipLv:         info.VipLevel,                               //VIP等级
		LeaderRank:    "",                                          //排名
		Stage:         strconv.FormatUint(uint64(info.Param1), 10), //关卡位置
		DailyStage:    strconv.FormatUint(uint64(info.Param2), 10), //日常关卡位置
		DeviceId:      info.DeviceId,                               //设备ID
		DeviceMobile:  "",                                          //机型
		ClientIp:      info.ClientIp,                               //终端IP
		DeviceType:    "",                                          //应用类别（1是安卓，2是IOS）
		ChannelName:   info.ChannelId,                              //渠道ID
		DeviceCarrier: "",                                          //运营商
		DeviceOsVer:   "",
		IsNew:         info.Param4, //账号为新，为1，否则为0
		IsAreaNew:     1,           //角色微信，为1，否则为0
	}
}

//转换成创角注册日志
func ChangeLogToRegister(info LogInfo) LogRegister {
	return LogRegister{
		Time:          info.RegisterTime,
		UserName:      info.Uuid,
		UserId:        strconv.FormatUint(info.UserId, 10),
		UserType:      "",
		DeviceId:      info.DeviceId,  //设备ID
		DeviceMobile:  "",             //机型
		ClientIp:      info.ClientIp,  //终端IP
		DeviceType:    "",             //应用类别（1是安卓，2是IOS）
		ChannelName:   info.ChannelId, //渠道ID
		DeviceCarrier: "",             //运营商
		DeviceOsVer:   "",             //设备具体系统
	}
}

//装换成物品产出日志
func ChangeLogToOutPut(info LogInfo) LogOutPut {
	return LogOutPut{
		Time:          info.LogSendTime,
		Type:          strconv.FormatUint(uint64(info.Param4), 10),
		ItemId:        strconv.FormatUint(uint64(info.Param2), 10), //道具ID
		ItemType:      info.Param1,                                 //道具类型
		ItemName:      info.Param19,                                //道具名称
		ItemNum:       info.Param6,                                 //道具数量
		UserName:      info.Uuid,
		UserId:        strconv.FormatUint(info.UserId, 10),
		UserType:      "",
		AreaName:      uint32(0),      //登陆区名称
		UserLv:        info.UserLevel, //用户等级
		VipLv:         info.VipLevel,  //VIP等级
		DeviceId:      info.DeviceId,  //设备ID
		DeviceMobile:  "",             //机型
		ClientIp:      info.ClientIp,  //终端IP
		DeviceType:    "",             //应用类别（1是安卓，2是IOS）
		ChannelName:   info.ChannelId, //渠道ID
		DeviceCarrier: "",             //运营商
		DeviceOsVer:   "",             //设备具体系统
	}
}

//转换成物品消耗日志
func ChangeLogToConsume(info LogInfo) LogItemConsume {
	return LogItemConsume{
		Time:          info.LogSendTime,
		Type:          strconv.FormatUint(uint64(info.Param4), 10),
		ItemId:        strconv.FormatUint(uint64(info.Param2), 10), //道具ID
		ItemType:      info.Param1,                                 //道具类型
		ItemName:      info.Param19,                                //道具名称
		ItemNum:       info.Param6,                                 //道具数量
		UserName:      info.Uuid,
		UserId:        strconv.FormatUint(info.UserId, 10),
		UserType:      "",
		AreaName:      uint32(0),      //登陆区名称
		UserLv:        info.UserLevel, //用户等级
		VipLv:         info.VipLevel,  //VIP等级
		DeviceId:      info.DeviceId,  //设备ID
		DeviceMobile:  "",             //机型
		ClientIp:      info.ClientIp,  //终端IP
		DeviceType:    "",             //应用类别（1是安卓，2是IOS）
		ChannelName:   info.ChannelId, //渠道ID
		DeviceCarrier: "",             //运营商
		DeviceOsVer:   "",             //设备具体系统
	}
}

//转换成在线人数日志
func ChangeLogToUserOnline(info LogInfo) LogUserOnline {
	return LogUserOnline{
		Time:          info.LogSendTime,
		AreaName:      uint32(0),      //登陆区名称
		Count:         info.Param1,    //在线玩家数量
		DeviceId:      info.DeviceId,  //设备ID
		DeviceMobile:  "",             //机型
		ClientIp:      info.ClientIp,  //终端IP
		DeviceType:    "",             //应用类别（1是安卓，2是IOS）
		ChannelName:   info.ChannelId, //渠道ID
		DeviceCarrier: "",             //运营商
		DeviceOsVer:   "",             //设备具体系统
	}
}

//转换成关卡事件日志
func ChangeLogToStage(info LogInfo) LogChapterStage {
	result := LogChapterStage{
		Time:         info.LogSendTime,
		UserName:     info.Uuid,
		UserId:       strconv.FormatUint(info.UserId, 10),
		UserType:     "",
		AreaName:     uint32(0),      //登陆区名称
		UserLv:       info.UserLevel, //用户等级
		VipLv:        info.VipLevel,  //VIP等级
		Efficiency:   info.Param1,    //关卡ID
		Stage:        info.Param16,   //关卡名称
		IsSuccess:    info.Param3,    //0：成功，1，失败，2，开始
		DeviceId:     info.DeviceId,  //设备ID
		DeviceMobile: "",             //机型
		ClientIp:     info.ClientIp,  //终端IP
		DeviceType:   "",             //应用类别（1是安卓，2是IOS）
		ChannelName:  info.ChannelId, //渠道ID
		AppVersion:   "",             //版本号
	}
	switch info.Param2 {
	case LOG_STAGE_MAIN:
		result.StageType = LOG_STAGE_MAIN_STRING
	case LOG_STAGE_DAILY:
		result.StageType = LOG_STAGE_DAILY_STRING
	}
	return result
}

//转换成在线时长日志
func ChangeLogToOnlineTime(info LogInfo) LogOnlineTime {
	return LogOnlineTime{
		Time:         info.LogSendTime,
		UserName:     info.Uuid,
		UserId:       strconv.FormatUint(info.UserId, 10),
		UserType:     "",
		AreaName:     uint32(0), //登陆区名称
		Duration:     info.Param2,
		DeviceId:     info.DeviceId,  //设备ID
		DeviceMobile: "",             //机型
		ClientIp:     info.ClientIp,  //终端IP
		DeviceType:   "",             //应用类别（1是安卓，2是IOS）
		ChannelName:  info.ChannelId, //渠道ID
		AppVersion:   "",             //版本号
	}
}

//转换成等级日志
func ChangeLogToLevel(info LogInfo) LogLevel {
	return LogLevel{
		Time:         info.LogSendTime,
		UserName:     info.Uuid,
		UserId:       strconv.FormatUint(info.UserId, 10),
		UserType:     "",
		AreaName:     uint32(0),      //登陆区名称
		UserLv:       info.UserLevel, //用户等级
		VipLv:        info.VipLevel,  //VIP等级
		DeviceId:     info.DeviceId,  //设备ID
		DeviceMobile: "",             //机型
		ClientIp:     info.ClientIp,  //终端IP
		DeviceType:   "",             //应用类别（1是安卓，2是IOS）
		ChannelName:  info.ChannelId, //渠道ID
		AppVersion:   "",             //版本号
	}
}

//转换成充值货币进出日志
func ChangeLogToSilverConsume(info LogInfo) LogSilverConsume {
	result := LogSilverConsume{
		Time:       info.OnlineTime,
		UserName:   info.Uuid,
		UserId:     strconv.FormatUint(info.UserId, 10),
		UserType:   "",
		AreaName:   uint32(0),                                   //登陆区名称
		UserLv:     info.UserLevel,                              //用户等级
		VipLv:      info.VipLevel,                               //VIP等级
		LeaderRank: "",                                          //排名
		Stage:      strconv.FormatUint(uint64(info.Param8), 10), //关卡位置
		DailyStage: strconv.FormatUint(uint64(info.Param9), 10), //日常关卡位置

		ItemId:   strconv.FormatUint(uint64(info.Param2), 10), //道具ID
		ItemType: info.Param1,                                 //道具类型
		ItemName: info.Param19,
		ItemNum:  info.Param6, //道具名称

		DeviceId:     info.DeviceId,  //设备ID
		DeviceMobile: "",             //机型
		ClientIp:     info.ClientIp,  //终端IP
		DeviceType:   "",             //应用类别（1是安卓，2是IOS）
		ChannelName:  info.ChannelId, //渠道ID
	}

	switch info.ChildType {
	case LOG_CHILD_RESOURCE_DIAMOND_ADD:
		result.Consume = int(info.Param6)
	case LOG_CHILD_RESOURCE_DIAMOND_DEC:
		result.Consume = -int(info.Param6)
		switch info.Param4 {
		case LOG_CHILD_ACTION_SHOP_COST:
			result.ConsumeType = "商店购买消耗"
		case LOG_CHILD_ACTION_RECRUIT:
			result.ConsumeType = "抽卡消耗"
		}
	}

	return result
}
