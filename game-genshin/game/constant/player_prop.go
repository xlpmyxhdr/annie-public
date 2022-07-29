package constant

type PlayerProperty struct {
	PROP_EXP                             uint16
	PROP_BREAK_LEVEL                     uint16
	PROP_SATIATION_VAL                   uint16
	PROP_SATIATION_PENALTY_TIME          uint16
	PROP_LEVEL                           uint16
	PROP_LAST_CHANGE_AVATAR_TIME         uint16
	PROP_MAX_SPRING_VOLUME               uint16
	PROP_CUR_SPRING_VOLUME               uint16
	PROP_IS_SPRING_AUTO_USE              uint16
	PROP_SPRING_AUTO_USE_PERCENT         uint16
	PROP_IS_FLYABLE                      uint16
	PROP_IS_WEATHER_LOCKED               uint16
	PROP_IS_GAME_TIME_LOCKED             uint16
	PROP_IS_TRANSFERABLE                 uint16
	PROP_MAX_STAMINA                     uint16
	PROP_CUR_PERSIST_STAMINA             uint16
	PROP_CUR_TEMPORARY_STAMINA           uint16
	PROP_PLAYER_LEVEL                    uint16
	PROP_PLAYER_EXP                      uint16
	PROP_PLAYER_HCOIN                    uint16
	PROP_PLAYER_SCOIN                    uint16
	PROP_PLAYER_MP_SETTING_TYPE          uint16
	PROP_IS_MP_MODE_AVAILABLE            uint16
	PROP_PLAYER_WORLD_LEVEL              uint16
	PROP_PLAYER_RESIN                    uint16
	PROP_PLAYER_WAIT_SUB_HCOIN           uint16
	PROP_PLAYER_WAIT_SUB_SCOIN           uint16
	PROP_IS_ONLY_MP_WITH_PS_PLAYER       uint16
	PROP_PLAYER_MCOIN                    uint16
	PROP_PLAYER_WAIT_SUB_MCOIN           uint16
	PROP_PLAYER_LEGENDARY_KEY            uint16
	PROP_IS_HAS_FIRST_SHARE              uint16
	PROP_PLAYER_FORGE_POINT              uint16
	PROP_CUR_CLIMATE_METER               uint16
	PROP_CUR_CLIMATE_TYPE                uint16
	PROP_CUR_CLIMATE_AREA_ID             uint16
	PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE   uint16
	PROP_PLAYER_WORLD_LEVEL_LIMIT        uint16
	PROP_PLAYER_WORLD_LEVEL_ADJUST_CD    uint16
	PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM uint16
	PROP_PLAYER_HOME_COIN                uint16
	PROP_PLAYER_WAIT_SUB_HOME_COIN       uint16
}

func GetPlayerPropertyConst() (r *PlayerProperty) {
	r = new(PlayerProperty)
	r.PROP_EXP = 1001
	r.PROP_BREAK_LEVEL = 1002
	r.PROP_SATIATION_VAL = 1003
	r.PROP_SATIATION_PENALTY_TIME = 1004
	r.PROP_LEVEL = 4001
	r.PROP_LAST_CHANGE_AVATAR_TIME = 10001
	r.PROP_MAX_SPRING_VOLUME = 10002
	r.PROP_CUR_SPRING_VOLUME = 10003
	r.PROP_IS_SPRING_AUTO_USE = 10004
	r.PROP_SPRING_AUTO_USE_PERCENT = 10005
	r.PROP_IS_FLYABLE = 10006
	r.PROP_IS_WEATHER_LOCKED = 10007
	r.PROP_IS_GAME_TIME_LOCKED = 10008
	r.PROP_IS_TRANSFERABLE = 10009
	r.PROP_MAX_STAMINA = 10010
	r.PROP_CUR_PERSIST_STAMINA = 10011
	r.PROP_CUR_TEMPORARY_STAMINA = 10012
	r.PROP_PLAYER_LEVEL = 10013
	r.PROP_PLAYER_EXP = 10014
	r.PROP_PLAYER_HCOIN = 10015 // 原石
	r.PROP_PLAYER_SCOIN = 10016 // 摩拉
	r.PROP_PLAYER_MP_SETTING_TYPE = 10017
	r.PROP_IS_MP_MODE_AVAILABLE = 10018
	r.PROP_PLAYER_WORLD_LEVEL = 10019
	r.PROP_PLAYER_RESIN = 10020
	r.PROP_PLAYER_WAIT_SUB_HCOIN = 10022
	r.PROP_PLAYER_WAIT_SUB_SCOIN = 10023
	r.PROP_IS_ONLY_MP_WITH_PS_PLAYER = 10024
	r.PROP_PLAYER_MCOIN = 10025 // 创世结晶
	r.PROP_PLAYER_WAIT_SUB_MCOIN = 10026
	r.PROP_PLAYER_LEGENDARY_KEY = 10027
	r.PROP_IS_HAS_FIRST_SHARE = 10028
	r.PROP_PLAYER_FORGE_POINT = 10029
	r.PROP_CUR_CLIMATE_METER = 10035
	r.PROP_CUR_CLIMATE_TYPE = 10036
	r.PROP_CUR_CLIMATE_AREA_ID = 10037
	r.PROP_CUR_CLIMATE_AREA_CLIMATE_TYPE = 10038
	r.PROP_PLAYER_WORLD_LEVEL_LIMIT = 10039
	r.PROP_PLAYER_WORLD_LEVEL_ADJUST_CD = 10040
	r.PROP_PLAYER_LEGENDARY_DAILY_TASK_NUM = 10041
	r.PROP_PLAYER_HOME_COIN = 10042
	r.PROP_PLAYER_WAIT_SUB_HOME_COIN = 10043
	return r
}
