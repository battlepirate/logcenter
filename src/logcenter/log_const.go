package main

import ()

//kafka的const值
const (
	KAFKA_LOGIN         = "dn_login"
	KAFKA_REGISTER      = "dn_register"
	KAFKA_OUTPUT        = "dn_output"
	KAFKA_CONSUME       = "dn_itemconsume"
	KAFKA_USERONLINE    = "dn_useronline"
	KAFKA_STAGE         = "dn_chapter_stage"
	KAFKA_ONLINETIME    = "dn_onlinetime"
	KAFKA_LEVEL         = "dn_level"
	KAFKA_SILVERCONSUME = "dn_silverConsume"
)

//elasticsearch的const值
const (
	//ES_INDEX      = "logs"
	ES_TYPE       = "data"
	ES_LOGIN      = "dn_login"
	ES_REGISTER   = "dn_register"
	ES_OUTPUT     = "dn_output"
	ES_CONSUME    = "dn_itemconsume"
	ES_USERONLINE = "dn_useronline"
)

//TABLE 资源属性
const (
	LOG_RESOURCE_ATTR_LEVEL uint32 = 1 //等级
)

//日志主类型
const (
	LOG_MAIN_USER_CHANGE     uint32 = 1 //玩家角色变化
	LOG_MAIN_RECHARGE_RELATE uint32 = 2 //充值支付消耗
	LOG_MAIN_RESOURCE_CHANGE uint32 = 3 //资源变化
	LOG_MAIN_ACTION          uint32 = 4 //玩家行为
	LOG_MAIN_SERVER_INFO     uint32 = 5 //服务器基础信息
	LOG_MAIN_OTHER           uint32 = 6 //其他
)

//日志子类型
//USER CHANGE 玩家变化
const (
	LOG_CHILD_USER_CREATE  uint32 = 1 //用户创角色、激活
	LOG_CHILD_USER_LOGIN   uint32 = 2 //玩家登录
	LOG_CHILD_USER_OFFLINE uint32 = 3 //玩家下线
	LOG_CHILD_USER_UPGRADE uint32 = 4 //玩家升级
)

//RECHARGE 充值，金币的变化
const (
	LOG_CHILD_PAY_RECHARGE    uint32 = 1 //充值付费(一种行为)
	LOG_CHILD_PAY_GIFT        uint32 = 2 //非充值金币获取
	LOG_CHILD_PAY_CONSUME     uint32 = 3 //金币消耗
	LOG_CHILD_PAY_PURCHASE    uint32 = 4 //充值金币获取
	LOG_CHILD_PAY_RECHARGE_EX uint32 = 5 //充值付费类型附加(LOG_CHILD_PAY_RECHARGE类型字段已全满)
)

//RESOURCE 资源（不包括金币）
const (
	LOG_CHILD_RESOURCE_ADD          uint32 = 1 //资源产出
	LOG_CHILD_RESOURCE_DEC          uint32 = 2 //资源消耗
	LOG_CHILD_RESOURCE_ATTR_UP      uint32 = 3 //资源属性升级
	LOG_CHILD_RESOURCE_STATE_CHANGE uint32 = 4 //资源状态变化
	LOG_CHILD_RESOURCE_DIAMOND_ADD  uint32 = 5 //钻石产出
	LOG_CHILD_RESOURCE_DIAMOND_DEC  uint32 = 6 //钻石消耗
)

//SERVER 服务器基本信息
const (
	LOG_CHILD_SERVER_PCU uint32 = 1 //服务器在线人数，每5分钟发一次
)

//OTHER 其他补充日志 放这里
const (
	LOG_OTHER_USER_VAILD uint32 = 1 //玩家变为有效用户
	LOG_OTHER_TIMER_RANK uint32 = 2 //定时发送排行榜
	LOG_OTHER_DEVICE_ID  uint32 = 3 //用户打开客户端发送设备好
)

//ACTION 玩家行为(理论上 一条CS协议 一个行为)
//TODO  新行为添加到最后面，不能中间插入
const (
	LOG_CHILD_ACTION_PROCESS_MAIL           uint32 = 1  //邮件领取
	LOG_CHILD_ACTION_USER_MAIL              uint32 = 2  //用户邮件
	LOG_CHILD_ACTION_DEL_MAIL               uint32 = 3  //手动删除邮件
	LOG_CHILD_ACTION_RECHARGE               uint32 = 4  //充值
	LOG_CHILD_ACTION_ADD_FRIEND             uint32 = 5  //添加好友或者黑名单
	LOG_CHILD_ACTION_DEL_FRIEND             uint32 = 6  //删除好友或黑名单
	LOG_CHILD_ACTION_GIVE_SPIRIT            uint32 = 7  //赠送好友精力
	LOG_CHILD_ACTION_RECEIVE_SPIRIT         uint32 = 8  //领取好友精力
	LOG_CHILD_ACTION_ADD_FRIEND_REQ         uint32 = 9  //发送好友申请
	LOG_CHILD_ACTION_DEL_FRIEND_REQ         uint32 = 10 //删除好友申请
	LOG_CHILD_ACTION_FRIEND_PK              uint32 = 11 //好友切磋
	LOG_CHILD_ACTION_USE_ITEM               uint32 = 12 //道具使用
	LOG_CHILD_ACTION_LEAD_LEVEL_UP          uint32 = 13 //主角升级
	LOG_CHILD_ACTION_CREATE                 uint32 = 14 //创建角色
	LOG_CHILD_ACTION_CREATE_AWARD           uint32 = 15 //建号发放
	LOG_CHILD_ACTION_CHAT                   uint32 = 16 //聊天
	LOG_CHILD_ACTION_GIFT_CODE_REWARD       uint32 = 17 //礼品码兑换
	LOG_CHILD_ACTION_GM_DELETE              uint32 = 18 //gm后台删除玩家道具
	LOG_CHILD_ACTION_SELL_ITEM              uint32 = 19 //出售道具
	LOG_CHILD_ACTION_BATTLE                 uint32 = 20 //战斗奖励
	LOG_CHILD_ACTION_RECEIVE_CHAPTER_BOX    uint32 = 21 //章节宝箱
	LOG_CHILD_ACTION_BATTLE_COST            uint32 = 22 //战斗消耗体力
	LOG_CHILD_ACTION_INTERACTIVE_TOUCH      uint32 = 23 //交互功能事件完成奖励
	LOG_CHILD_ACTION_PATROL_AWARD           uint32 = 24 //巡逻事件奖励
	LOG_CHILD_ACTION_PATROL_REDUCE_RESOURCE uint32 = 25 //巡逻消耗加成资源
	LOG_CHILD_ACTION_HERO_LEVEL_UP          uint32 = 26 //英雄升级
	LOG_CHILD_ACTION_UPGRADE_HERO_AWAKE     uint32 = 27 //英雄觉醒
	LOG_CHILD_ACTION_UPGRADE_HERO_SKILL     uint32 = 28 //英雄技能升级
	LOG_CHILD_ACTION_GM_ADD_USER_AWARD      uint32 = 29 //GM后台添加玩家数据
	LOG_CHILD_ACTION_ITEM_CONVERT           uint32 = 30 //道具合成
	LOG_CHILD_ACTION_UPGRADE_HERO_STAR      uint32 = 31 //英雄升星
	LOG_CHILD_ACTION_COOK_MAKE_AWARD        uint32 = 32 //制作食谱奖励
	LOG_CHILD_ACTION_COOK_MAKE_REDUCE       uint32 = 33 //制作食谱消耗
	LOG_CHILD_ACTION_MUSIC_STRENGTH         uint32 = 34 //苍谱强化消耗
	LOG_CHILD_ACTION_RECRUIT                uint32 = 35 //添加招募日志
	LOG_CHILD_ACTION_COOKING_EAT_AWARD      uint32 = 36 //英雄进食奖励
	LOG_CHILD_ACTION_INTERACTIVE_AWARD      uint32 = 37 //恋爱交互奖励
	LOG_CHILD_ACTION_EXPEDITION_COST        uint32 = 38 //远征消耗体力
	LOG_CHILD_ACTION_EXPEDITION_AWARD       uint32 = 39 //远征奖励
	LOG_CHILD_ACTION_SHOP_AWARD             uint32 = 40 //商店购买
	LOG_CHILD_ACTION_SHOP_COST              uint32 = 41 //商店购买消耗
	LOG_CHILD_ACTION_ACHIEVETASK_AWARD      uint32 = 40 //成就任务奖励
	LOG_CHILD_ACTION_USER_RENAME_COST       uint32 = 41 //改名消耗
	LOG_CHILD_ACTION_STAGE                  uint32 = 42 //关卡事件
	LOG_CHILD_ACTION_QUICK_BUY_REDUCE       uint32 = 43 //快速购买消耗
	LOG_CHILD_ACTION_QUICK_BUY_AWARD        uint32 = 44 //快速购买奖励
)

//关卡行为常量
const (
	LOG_STAGE_MAIN         uint32 = 1 //主线
	LOG_STAGE_DAILY        uint32 = 2 //日常
	LOG_STAGE_MAIN_STRING  string = "主线"
	LOG_STAGE_DAILY_STRING string = "日常"

	LOG_STAGE_SUCCESS uint32 = 0 //成功
	LOG_STAGE_LOSE    uint32 = 1 //失败
	LOG_STAGE_BEGIN   uint32 = 2 //开始
)
