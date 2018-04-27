package main

import (
	l4g "base/log4go"
	"encoding/json"
	"sync"
)

const (
	LOGSENDSEP = "+++" //日志分隔符
)

//日志管理器
type LogM struct {
	in_count  uint32
	in        sync.RWMutex
	out_count uint32
	out       sync.RWMutex
}

func (this *LogM) InAdd(count uint32) {
	this.in.RLock()
	defer this.in.RUnlock()
	this.in_count += count
}

func (this *LogM) InRead() uint32 {
	this.in.RLock()
	defer this.in.RUnlock()
	return this.in_count
}

func (this *LogM) OutAdd(count uint32) {
	this.out.RLock()
	defer this.out.RUnlock()
	this.out_count += count
}

func (this *LogM) OutRead() uint32 {
	this.out.RLock()
	defer this.out.RUnlock()
	return this.out_count
}

func (this *LogM) InTest() {
	this.InAdd(uint32(1))
	l4g.Info("infos in_count is up to %d", this.InRead())
}

func (this *LogM) OutTest() {
	this.OutAdd(uint32(1))
	l4g.Info("infos out_count is up to %d", this.OutRead())
}

func UnmarshalLog(infos []byte) []LogInfo {
	var result []LogInfo
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLog json.Unmarshal error %v %v", e, string(infos))
		return nil
	}
	l4g.Info("UnmarshalLog result is %+v", result)
	return result
}

func UnmarshalLogLogin(infos []byte) (LogLogin, bool) {
	var result LogLogin
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLoginLog json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLoginLog result is %+v", result)

	return result, true
}

func UnmarshalLogRegister(infos []byte) (LogRegister, bool) {
	var result LogRegister
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogRegister json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogRegister result is %+v", result)

	return result, true
}

func UnmarshalLogUserOnline(infos []byte) (LogUserOnline, bool) {
	var result LogUserOnline
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogUserOnline json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogUserOnline result is %+v", result)

	return result, true
}
func UnmarshalLogOutPut(infos []byte) (LogOutPut, bool) {
	var result LogOutPut
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogOutPut json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogOutPut result is %+v", result)

	return result, true
}

func UnmarshalLogItemConsume(infos []byte) (LogItemConsume, bool) {
	var result LogItemConsume
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogItemConsume json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogItemConsume result is %+v", result)

	return result, true
}

func UnmarshalLogChapterStage(infos []byte) (LogChapterStage, bool) {
	var result LogChapterStage
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogChapterStage json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogChapterStage result is %+v", result)

	return result, true
}

func UnmarshalLogOnlineTime(infos []byte) (LogOnlineTime, bool) {
	var result LogOnlineTime
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogOnlineTime json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogOnlineTime result is %+v", result)

	return result, true
}

func UnmarshalLogLevel(infos []byte) (LogLevel, bool) {
	var result LogLevel
	if e := json.Unmarshal(infos, &result); e != nil {
		l4g.Error("UnmarshalLogLevel json.Unmarshal error %v", e)
		return result, false
	}
	l4g.Info("UnmarshalLogLevel result is %+v", result)

	return result, true
}

func UnmarshalLogToEs(infos []byte, topics string) (interface{}, bool) {
	switch topics {
	case KAFKA_LOGIN:
		return UnmarshalLogLogin(infos)
	case KAFKA_REGISTER:
		return UnmarshalLogRegister(infos)
	case KAFKA_OUTPUT:
		return UnmarshalLogOutPut(infos)
	case KAFKA_CONSUME:
		return UnmarshalLogItemConsume(infos)
	case KAFKA_USERONLINE:
		return UnmarshalLogUserOnline(infos)
	case KAFKA_STAGE:
		return UnmarshalLogChapterStage(infos)
	case KAFKA_ONLINETIME:
		return UnmarshalLogOnlineTime(infos)
	case KAFKA_LEVEL:
		return UnmarshalLogLevel(infos)
	default:
		return nil, false
	}
}
