package common

//
// 文件名: date_time.go<br/>
// 创建时间: 2016年8月3日-下午4:54:01<br/>
// 简介: <br/>
// 详情: 时间工具类<br/>
// Copyright (C) 2013 duhaibo0404@gmail.com. All Rights Reserved.<br/>
//
import (
	"errors"
	"fmt"
	"time"
)

var (
	FIRST_DAY      int32    = 2361222 //对应 1752-09-14 这天
	SECS_PER_DAY   int32    = 86400
	MSECS_PER_DAY  int32    = 86400000
	SECS_PER_HOUR  int32    = 3600
	MSECS_PER_HOUR int32    = 3600000
	SECS_PER_MIN   int32    = 60
	MSECS_PER_MIN  int32    = 60000
	MSECS_PER_SEC  int32    = 1000
	FIRST_YEAR     int32    = 1752
	anDaysInMonth  []int32  = []int32{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	astrWeeks      []string = []string{"", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

type TDate struct {
	nDateVal int32
}

//获取当前日期
func CurDate() TDate {
	nYear, nMonth, nDay := time.Now().Date()
	val := gregToJulian(int32(nYear), int32(nMonth), int32(nDay))
	return TDate{val}
}

func NewDefaultDate() *TDate {
	return &TDate{0}
}

//创建指定日期
func NewSpecDate(nYear, nMonth, nDay int32) (*TDate, error) {
	tRetDate := &TDate{0}

	if !tRetDate.SetDate(nYear, nMonth, nDay) {
		tRetDate = nil
		return nil, errors.New("invalid param")
	}

	return tRetDate, nil
}

//按照%04d-%02d-%02d格式创建字符串指定日期
func NewStringDate(strDate string) (*TDate, error) {
	var nYear, nMonth, nDay int32

	if _, err := fmt.Sscanf(strDate, "%04d-%02d-%02d", &nYear, &nMonth, &nDay); err != nil {
		return nil, err
	}

	tRetDate := &TDate{0}
	if !tRetDate.SetDate(nYear, nMonth, nDay) {
		tRetDate = nil
		return nil, errors.New("invalid date")
	}

	return tRetDate, nil
}

func (this TDate) IsNull() bool {
	return this.nDateVal == 0
}

func (this TDate) IsValid() bool {
	return (this.nDateVal >= FIRST_DAY)
}

func gregToJulian(nYear, nMonth, nDay int32) int32 {
	var nC, nYa int32

	if nYear <= 99 {
		nYear += 1900
	}

	if nMonth > 2 {
		nMonth -= 3
	} else {
		nMonth += 9
		nYear--
	}

	nC = nYear
	nC /= 100
	nYa = nYear - 100*nC

	return 1721119 + nDay + ((146097 * nC) / 4) + ((1461 * nYa) / 4) + ((153*nMonth + 2) / 5)
}

func julianToGreg(nJulian int32, pnYear, pnMonth, pnDay *int32) {
	var nX int32
	var nJ int32 = nJulian - 1721119

	*pnYear = ((nJ * 4) - 1) / 146097
	nJ = (nJ * 4) - (146097 * *pnYear) - 1
	nX = nJ / 4
	nJ = ((nX * 4) + 3) / 1461
	*pnYear = (100 * *pnYear) + nJ
	nX = (nX * 4) + 3 - (1461 * nJ)
	nX = (nX + 4) / 4
	*pnMonth = ((5 * nX) - 3) / 153
	nX = (5 * nX) - 3 - (153 * *pnMonth)
	*pnDay = (nX + 5) / 5

	if *pnMonth < 10 {
		*pnMonth += 3
	} else {
		*pnMonth -= 9
		*pnYear++
	}
}

func (this TDate) Year() int32 {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)
	return nYear
}

func (this TDate) Month() int32 {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)
	return nMonth
}

func (this TDate) Day() int32 {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)
	return nDay
}

//获取目前日期对应周几，返回(1-7)
func (this TDate) DayOfWeek() int32 {
	return (((this.nDateVal+1)%7)+6)%7 + 1
}

//获取目前日期是所在年份中的第多少天(比如元旦表示第一天)
func (this TDate) DayOfYear() int32 {
	return this.nDateVal - gregToJulian(this.Year(), 1, 1) + 1
}

//判断指定年份是否是闰年
func (this TDate) IsLeapYear(nYear int32) bool {
	return (nYear%4 == 0 && nYear%100 != 0 || nYear%400 == 0)
}

//获取目前日期所在月份的总共天数
func (this TDate) DaysInMonth() int32 {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)
	if nMonth == 2 && this.IsLeapYear(nYear) {
		return 29
	} else {
		return anDaysInMonth[nMonth]
	}
}

//获取目前日期所在年份的总共天数
func (this TDate) DaysInYear() int32 {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)
	if this.IsLeapYear(nYear) {
		return 366
	} else {
		return 365
	}
}

//获取目前日期是周几的字符串
func (this TDate) DayOfWeekToString() string {
	return astrWeeks[this.DayOfWeek()]
}

func (this TDate) isValid(nYear, nMonth, nDay int32) bool {
	if nYear >= 0 && nYear <= 99 {
		nYear += 1900
	}

	if nYear < FIRST_YEAR || (nYear == FIRST_YEAR && (nMonth < 9 || (nMonth == 9 && nDay < 14))) {
		return false
	}

	return (nDay > 0 && nMonth > 0 && nMonth <= 12) && (nDay <= anDaysInMonth[nMonth] || (nDay == 29 && nMonth == 2 && this.IsLeapYear(nYear)))
}

//将目前日志设置为指定日期
func (this *TDate) SetDate(nYear, nMonth, nDay int32) bool {
	if !this.isValid(nYear, nMonth, nDay) {
		return false
	}

	this.nDateVal = gregToJulian(nYear, nMonth, nDay)
	return true
}

//返回目前日期加上指定天数对应的日期，参数可正可负
func (this TDate) AddDays(nDays int32) TDate {
	return TDate{this.nDateVal + nDays}
}

//返回目前日期距离指定日期的天数差，为负值时表示目前日期相比指定日期是未来时间
func (this TDate) DaysTo(ptDstDate *TDate) int32 {
	return ptDstDate.nDateVal - this.nDateVal
}

//将目前日期按照%04d-%02d-%02d格式转换为字符串
func (this TDate) ToString() string {
	var nYear, nMonth, nDay int32
	julianToGreg(this.nDateVal, &nYear, &nMonth, &nDay)

	return fmt.Sprintf("%04d-%02d-%02d", nYear, nMonth, nDay)
}

//以下就是一些日期比较函数
func (this TDate) IsEqual(ptDate *TDate) bool {
	return this.nDateVal == ptDate.nDateVal
}

func (this TDate) IsNotEqual(ptDate *TDate) bool {
	return this.nDateVal != ptDate.nDateVal
}

func (this TDate) IsLess(ptDate *TDate) bool {
	return this.nDateVal < ptDate.nDateVal
}

func (this TDate) IsLessOrEqual(ptDate *TDate) bool {
	return this.nDateVal <= ptDate.nDateVal
}

func (this TDate) IsGreat(ptDate *TDate) bool {
	return this.nDateVal > ptDate.nDateVal
}

func (this TDate) IsGreatOrEqual(ptDate *TDate) bool {
	return this.nDateVal >= ptDate.nDateVal
}

type TTime struct {
	nTimeVal int32
}

func (this TTime) IsValid() bool {
	return this.nTimeVal < MSECS_PER_DAY
}

func (this TTime) isValid(nHour, nMin, nSec, nMs int32) bool {
	return (nHour < 24 && nMin < 60 && nSec < 60 && nMs < 1000)
}

//将当前时间设置为指定时间
func (this *TTime) SetTime(nHour, nMin, nSec, nMs int32) bool {
	if !this.isValid(nHour, nMin, nSec, nMs) {
		return false
	} else {
		this.nTimeVal = (nHour*SECS_PER_HOUR+nMin*SECS_PER_MIN+nSec)*MSECS_PER_SEC + nMs
		return true
	}
}

//返回当前时间
func CurTime() TTime {
	tNow := time.Now()
	nHour, nMinute, nSecond := tNow.Clock()
	return TTime{int32(nHour)*MSECS_PER_HOUR + int32(nMinute)*MSECS_PER_MIN + int32(nSecond)*MSECS_PER_SEC + int32(tNow.Nanosecond())/1000000}
}

//生成默认时间，即零点时间
func NewDefaultTime() *TTime {
	return &TTime{0}
}

//生成指定时间
func NewSpecTime(nHour, nMin, nSec, nMs int32) (*TTime, error) {
	tRetTime := &TTime{0}

	if !tRetTime.SetTime(nHour, nMin, nSec, nMs) {
		tRetTime = nil
		return nil, errors.New("invalid param")
	}

	return tRetTime, nil
}

//按照%02d:%02d:%02d格式格式创建字符串指定时间(精确到秒)
func NewNormalTime(strTime string) (*TTime, error) {
	var nHour, nMin, nSec int32

	if _, err := fmt.Sscanf(strTime, "%02d:%02d:%02d", &nHour, &nMin, &nSec); err != nil {
		return nil, err
	}

	tRetTime := &TTime{0}
	if !tRetTime.SetTime(nHour, nMin, nSec, 0) {
		tRetTime = nil
		return nil, errors.New("invalid time")
	}

	return tRetTime, nil
}

//按照%02d:%02d:%02d:%03d格式格式创建字符串指定时间(精确到毫秒)
func NewDetailTime(strTime string) (*TTime, error) {
	var nHour, nMin, nSec, nMs int32

	_, err := fmt.Sscanf(strTime, "%02d:%02d:%02d:%03d", &nHour, &nMin, &nSec, &nMs)
	if err != nil {
		return nil, err
	}

	tRetTime := &TTime{0}
	if !tRetTime.SetTime(nHour, nMin, nSec, nMs) {
		tRetTime = nil
		return nil, errors.New("invalid time")
	}

	return tRetTime, nil
}

//返回当前时间中的小时数
func (this TTime) Hour() int32 {
	return this.nTimeVal / MSECS_PER_HOUR
}

//返回当前时间中的分钟数
func (this TTime) Minute() int32 {
	return (this.nTimeVal % MSECS_PER_HOUR) / MSECS_PER_MIN
}

//返回当前时间中的秒数
func (this TTime) Second() int32 {
	return (this.nTimeVal / MSECS_PER_SEC) % SECS_PER_MIN
}

//返回当前时间中的毫秒数
func (this TTime) MilliSec() int32 {
	return this.nTimeVal % MSECS_PER_SEC
}

func (this TTime) AddSecs(nSecs int32) TTime {
	return this.AddMilliSecs(nSecs * MSECS_PER_SEC)
}

func (this TTime) AddMilliSecs(nMs int32) TTime {
	tTmpTime := TTime{0}

	if nMs < 0 {
		nNegDays := (MSECS_PER_DAY - nMs) / MSECS_PER_DAY
		tTmpTime.nTimeVal = (this.nTimeVal + nMs + nNegDays*MSECS_PER_DAY) % MSECS_PER_DAY
	} else {
		tTmpTime.nTimeVal = (this.nTimeVal + nMs) % MSECS_PER_DAY
	}

	return tTmpTime
}

func (this TTime) SecsTo(ptDstTime *TTime) int32 {
	return (ptDstTime.nTimeVal - this.nTimeVal) / MSECS_PER_SEC
}

func (this TTime) MilliSecsTo(ptDstTime *TTime) int32 {
	return ptDstTime.nTimeVal - this.nTimeVal
}

//将时间转换为精确到毫秒的字符串，格式为%02d:%02d:%02d:%03d
func (this TTime) ToDetailTime() string {
	return fmt.Sprintf("%02d:%02d:%02d:%03d", this.Hour(), this.Minute(), this.Second(), this.MilliSec())
}

//将时间转换为精确到秒的字符串，格式为%02d:%02d:%02d
func (this TTime) ToNormalTime() string {
	return fmt.Sprintf("%02d:%02d:%02d", this.Hour(), this.Minute(), this.Second())
}

//以下为时间的一些比较函数
func (this TTime) IsEqual(ptTime *TTime) bool {
	return this.nTimeVal == ptTime.nTimeVal
}

func (this TTime) IsNotEqual(ptTime *TTime) bool {
	return this.nTimeVal != ptTime.nTimeVal
}

func (this TTime) IsLess(ptTime *TTime) bool {
	return this.nTimeVal < ptTime.nTimeVal
}

func (this TTime) IsLessOrEqual(ptTime *TTime) bool {
	return this.nTimeVal <= ptTime.nTimeVal
}

func (this TTime) IsGreat(ptTime *TTime) bool {
	return this.nTimeVal > ptTime.nTimeVal
}

func (this TTime) IsGreatOrEqual(ptTime *TTime) bool {
	return this.nTimeVal >= ptTime.nTimeVal
}

type TDateTime struct {
	tDate TDate
	tTime TTime
}

func CurDateTime() TDateTime {
	return TDateTime{CurDate(), CurTime()}
}

//生成指定数值的日期时间
func NewDateTime(nYear, nMonth, nDay, nHour, nMin, nSec, nMs int32) (*TDateTime, error) {
	ptDate, errDate := NewSpecDate(nYear, nMonth, nDay)
	if errDate != nil {
		return nil, errDate
	}

	ptTime, errTime := NewSpecTime(nHour, nMin, nSec, nMs)
	if errTime != nil {
		ptDate = nil
		return nil, errTime
	}

	return &TDateTime{*ptDate, *ptTime}, nil
}

//根据TDate TTime生成日期时间
func NewSpecDateTime(ptDate *TDate, ptTime *TTime) *TDateTime {
	return &TDateTime{*ptDate, *ptTime}
}

//根据%04d-%02d-%02d %02d:%02d:%02d格式的字符串生成日期时间(精确到秒)
func NewNormalDateTime(strDateTime string) (*TDateTime, error) {
	var nYear, nMonth, nDay, nHour, nMin, nSec int32

	if _, err := fmt.Sscanf(strDateTime, "%04d-%02d-%02d %02d:%02d:%02d", &nYear, &nMonth, &nDay, &nHour, &nMin, &nSec); err != nil {
		return nil, err
	}

	return NewDateTime(nYear, nMonth, nDay, nHour, nMin, nSec, 0)
}

//根据%04d-%02d-%02d %02d:%02d:%02d:%03d格式的字符串生成日期时间(精确到毫秒)
func NewDetailDateTime(strDateTime string) (*TDateTime, error) {
	var nYear, nMonth, nDay, nHour, nMin, nSec, nMs int32

	if _, err := fmt.Sscanf(strDateTime, "%04d-%02d-%02d %02d:%02d:%02d:%03d", &nYear, &nMonth, &nDay, &nHour, &nMin, &nSec, &nMs); err != nil {
		return nil, err
	}

	return NewDateTime(nYear, nMonth, nDay, nHour, nMin, nSec, nMs)
}

//将目前日期时间设置为指定数值
func (this *TDateTime) SetDateTime(nYear, nMonth, nDay, nHour, nMin, nSec, nMs int32) bool {
	ptDate := &this.tDate
	ptTime := &this.tTime

	return ptDate.SetDate(nYear, nMonth, nDay) && ptTime.SetTime(nHour, nMin, nSec, nMs)
}

//返回目前日期时间加上指定天数后的日期时间，参数可正可负
func (this TDateTime) AddDays(nDays int32) TDateTime {
	return TDateTime{this.tDate.AddDays(nDays), this.tTime}
}

//返回目前日期时间加上指定时数后的日期时间，参数可正可负
func (this TDateTime) AddHours(sdwHours int64) TDateTime {
	return this.AddMilliSecs(sdwHours * int64(MSECS_PER_HOUR))
}

//返回目前日期时间加上指定分钟数后的日期时间，参数可正可负
func (this TDateTime) AddMinutes(sdwMins int64) TDateTime {
	return this.AddMilliSecs(sdwMins * int64(MSECS_PER_MIN))
}

//返回目前日期时间加上指定秒数后的日期时间，参数可正可负
func (this TDateTime) AddSecs(sdwSecs int64) TDateTime {
	return this.AddMilliSecs(sdwSecs * int64(MSECS_PER_SEC))
}

//返回目前日期时间加上指定毫秒数后的日期时间，参数可正可负
func (this TDateTime) AddMilliSecs(sdwMs int64) TDateTime {
	nDateVal := this.tDate.nDateVal
	nTimeVal := this.tTime.nTimeVal
	var nSign int32 = 1

	if sdwMs < 0 {
		sdwMs = -sdwMs
		nSign = -1
	}

	if sdwMs >= int64(MSECS_PER_DAY) {
		nDateVal += nSign * int32(sdwMs/int64(MSECS_PER_DAY))
		sdwMs %= int64(MSECS_PER_DAY)
	}

	nTimeVal += nSign * int32(sdwMs)
	if nTimeVal < 0 {
		nTimeVal = MSECS_PER_DAY - nTimeVal - 1
		nDateVal -= nTimeVal / MSECS_PER_DAY
		nTimeVal %= MSECS_PER_DAY
		nTimeVal = MSECS_PER_DAY - nTimeVal - 1
	} else if nTimeVal >= MSECS_PER_DAY {
		nDateVal += nTimeVal / MSECS_PER_DAY
		nTimeVal = nTimeVal % MSECS_PER_DAY
	}

	return TDateTime{tDate: TDate{nDateVal}, tTime: TTime{nTimeVal}}
}

//获取目前日期时间对应的格林尼治时间
func (this TDateTime) GetUnix() uint32 {
	tSysTime := time.Date(int(this.tDate.Year()), time.Month(this.tDate.Month()), int(this.tDate.Day()),
		int(this.tTime.Hour()), int(this.tTime.Minute()), int(this.tTime.Second()), 0, time.Local)
	return uint32(tSysTime.Unix())
}

//将目前日期时间设置为指定格林尼治时间对应的日期时间
func (this *TDateTime) SetUnix(dwUnix uint32) bool {
	tSysTime := time.Unix(int64(dwUnix), 0)
	nYear, nMonth, nDay := tSysTime.Date()
	nHour, nMinute, nSecond := tSysTime.Clock()
	return this.SetDateTime(int32(nYear), int32(nMonth), int32(nDay),
		int32(nHour), int32(nMinute), int32(nSecond), 0)
}

//计算两个日期时间之间的相隔天数
func (this TDateTime) DaysTo(ptDateTime *TDateTime) int32 {
	return this.tDate.DaysTo(&ptDateTime.tDate)
}

//计算两个日期时间之间的相隔秒数
func (this TDateTime) SecsTo(ptDateTime *TDateTime) int64 {
	return int64(this.tTime.SecsTo(&ptDateTime.tTime)) + int64(this.tDate.DaysTo(&ptDateTime.tDate))*int64(SECS_PER_DAY)
}

//计算两个日期时间之间的相隔毫秒数
func (this TDateTime) MilliSecsTo(ptDateTime *TDateTime) int64 {
	return int64(this.tTime.MilliSecsTo(&ptDateTime.tTime)) + int64(this.tDate.DaysTo(&ptDateTime.tDate))*int64(MSECS_PER_DAY)
}

//将目前日期时间转换为精确到秒的字符串，格式为%04d-%02d-%02d %02d:%02d:%02d
func (this TDateTime) ToNormalDateTime() string {
	return this.tDate.ToString() + " " + this.tTime.ToNormalTime()
}

//将目前日期时间转换为精确到毫秒的字符串，格式为%04d-%02d-%02d %02d:%02d:%02d:%03d
func (this TDateTime) ToDetailDateTime() string {
	return this.tDate.ToString() + " " + this.tTime.ToDetailTime()
}

//以下为日期时间的一些比较函数
func (this TDateTime) IsEqual(ptDateTime *TDateTime) bool {
	return this.tDate.IsEqual(&ptDateTime.tDate) && this.tTime.IsEqual(&ptDateTime.tTime)
}

func (this TDateTime) IsNotEqual(ptDateTime *TDateTime) bool {
	return this.tDate.IsNotEqual(&ptDateTime.tDate) || this.tTime.IsNotEqual(&ptDateTime.tTime)
}

func (this TDateTime) IsLess(ptDateTime *TDateTime) bool {
	if this.tDate.IsLess(&ptDateTime.tDate) {
		return true
	}

	return this.tDate.IsEqual(&ptDateTime.tDate) && this.tTime.IsLess(&ptDateTime.tTime)
}

func (this TDateTime) IsLessOrEqual(ptDateTime *TDateTime) bool {
	if this.tDate.IsLess(&ptDateTime.tDate) {
		return true
	}

	return this.tDate.IsEqual(&ptDateTime.tDate) && this.tTime.IsLessOrEqual(&ptDateTime.tTime)
}

func (this TDateTime) IsGreat(ptDateTime *TDateTime) bool {
	if this.tDate.IsGreat(&ptDateTime.tDate) {
		return true
	}

	return this.tDate.IsEqual(&ptDateTime.tDate) && this.tTime.IsGreat(&ptDateTime.tTime)
}

func (this TDateTime) IsGreatOrEqual(ptDateTime *TDateTime) bool {
	if this.tDate.IsGreat(&ptDateTime.tDate) {
		return true
	}

	return this.tDate.IsEqual(&ptDateTime.tDate) && this.tTime.IsGreatOrEqual(&ptDateTime.tTime)
}

func (this TDateTime) IsValid() bool {
	return this.tDate.IsValid() && this.tTime.IsValid()
}

func (this *TDateTime) GetDate() *TDate {
	return &this.tDate
}

func (this *TDateTime) GetTime() *TTime {
	return &this.tTime
}

func (this *TDateTime) SetDate(ptDate *TDate) {
	this.tDate = *ptDate
}

func (this *TDateTime) SetTime(ptTime *TTime) {
	this.tTime = *ptTime
}

func (this TDateTime) Year() int32 {
	return this.tDate.Year()
}

func (this TDateTime) Month() int32 {
	return this.tDate.Month()
}

func (this TDateTime) Day() int32 {
	return this.tDate.Day()
}

func (this TDateTime) Hour() int32 {
	return this.tTime.Hour()
}

func (this TDateTime) Minute() int32 {
	return this.tTime.Minute()
}

func (this TDateTime) Second() int32 {
	return this.tTime.Second()
}

func (this TDateTime) MilliSec() int32 {
	return this.tTime.MilliSec()
}

func WithinTime(start_time uint32, sep_time uint32) bool {
	fot := start_time
	fot -= fot % (24 * 3600)
	end_time := fot + (sep_time * 24 * 3600)
	now := uint32(time.Now().Unix())
	if fot <= now && now <= end_time {
		return true
	}
	return false
}
