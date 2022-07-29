package constant

type GrowCurve struct {
	GROW_CURVE_NONE             uint16
	GROW_CURVE_HP               uint16
	GROW_CURVE_ATTACK           uint16
	GROW_CURVE_STAMINA          uint16
	GROW_CURVE_STRIKE           uint16
	GROW_CURVE_ANTI_STRIKE      uint16
	GROW_CURVE_ANTI_STRIKE1     uint16
	GROW_CURVE_ANTI_STRIKE2     uint16
	GROW_CURVE_ANTI_STRIKE3     uint16
	GROW_CURVE_STRIKE_HURT      uint16
	GROW_CURVE_ELEMENT          uint16
	GROW_CURVE_KILL_EXP         uint16
	GROW_CURVE_DEFENSE          uint16
	GROW_CURVE_ATTACK_BOMB      uint16
	GROW_CURVE_HP_LITTLEMONSTER uint16
	GROW_CURVE_ELEMENT_MASTERY  uint16
	GROW_CURVE_PROGRESSION      uint16
	GROW_CURVE_DEFENDING        uint16
	GROW_CURVE_MHP              uint16
	GROW_CURVE_MATK             uint16
	GROW_CURVE_TOWERATK         uint16
	GROW_CURVE_HP_S5            uint16
	GROW_CURVE_HP_S4            uint16
	GROW_CURVE_HP_2             uint16
	GROW_CURVE_ATTACK_S5        uint16
	GROW_CURVE_ATTACK_S4        uint16
	GROW_CURVE_ATTACK_S3        uint16
	GROW_CURVE_STRIKE_S5        uint16
	GROW_CURVE_DEFENSE_S5       uint16
	GROW_CURVE_DEFENSE_S4       uint16
	GROW_CURVE_ATTACK_101       uint16
	GROW_CURVE_ATTACK_102       uint16
	GROW_CURVE_ATTACK_103       uint16
	GROW_CURVE_ATTACK_104       uint16
	GROW_CURVE_ATTACK_105       uint16
	GROW_CURVE_ATTACK_201       uint16
	GROW_CURVE_ATTACK_202       uint16
	GROW_CURVE_ATTACK_203       uint16
	GROW_CURVE_ATTACK_204       uint16
	GROW_CURVE_ATTACK_205       uint16
	GROW_CURVE_ATTACK_301       uint16
	GROW_CURVE_ATTACK_302       uint16
	GROW_CURVE_ATTACK_303       uint16
	GROW_CURVE_ATTACK_304       uint16
	GROW_CURVE_ATTACK_305       uint16
	GROW_CURVE_CRITICAL_101     uint16
	GROW_CURVE_CRITICAL_102     uint16
	GROW_CURVE_CRITICAL_103     uint16
	GROW_CURVE_CRITICAL_104     uint16
	GROW_CURVE_CRITICAL_105     uint16
	GROW_CURVE_CRITICAL_201     uint16
	GROW_CURVE_CRITICAL_202     uint16
	GROW_CURVE_CRITICAL_203     uint16
	GROW_CURVE_CRITICAL_204     uint16
	GROW_CURVE_CRITICAL_205     uint16
	GROW_CURVE_CRITICAL_301     uint16
	GROW_CURVE_CRITICAL_302     uint16
	GROW_CURVE_CRITICAL_303     uint16
	GROW_CURVE_CRITICAL_304     uint16
	GROW_CURVE_CRITICAL_305     uint16
}

func GetGrowCurveConst() (r *GrowCurve) {
	r = new(GrowCurve)
	r.GROW_CURVE_NONE = 0
	r.GROW_CURVE_HP = 1
	r.GROW_CURVE_ATTACK = 2
	r.GROW_CURVE_STAMINA = 3
	r.GROW_CURVE_STRIKE = 4
	r.GROW_CURVE_ANTI_STRIKE = 5
	r.GROW_CURVE_ANTI_STRIKE1 = 6
	r.GROW_CURVE_ANTI_STRIKE2 = 7
	r.GROW_CURVE_ANTI_STRIKE3 = 8
	r.GROW_CURVE_STRIKE_HURT = 9
	r.GROW_CURVE_ELEMENT = 10
	r.GROW_CURVE_KILL_EXP = 11
	r.GROW_CURVE_DEFENSE = 12
	r.GROW_CURVE_ATTACK_BOMB = 13
	r.GROW_CURVE_HP_LITTLEMONSTER = 14
	r.GROW_CURVE_ELEMENT_MASTERY = 15
	r.GROW_CURVE_PROGRESSION = 16
	r.GROW_CURVE_DEFENDING = 17
	r.GROW_CURVE_MHP = 18
	r.GROW_CURVE_MATK = 19
	r.GROW_CURVE_TOWERATK = 20
	r.GROW_CURVE_HP_S5 = 21
	r.GROW_CURVE_HP_S4 = 22
	r.GROW_CURVE_HP_2 = 23
	r.GROW_CURVE_ATTACK_S5 = 31
	r.GROW_CURVE_ATTACK_S4 = 32
	r.GROW_CURVE_ATTACK_S3 = 33
	r.GROW_CURVE_STRIKE_S5 = 34
	r.GROW_CURVE_DEFENSE_S5 = 41
	r.GROW_CURVE_DEFENSE_S4 = 42
	r.GROW_CURVE_ATTACK_101 = 1101
	r.GROW_CURVE_ATTACK_102 = 1102
	r.GROW_CURVE_ATTACK_103 = 1103
	r.GROW_CURVE_ATTACK_104 = 1104
	r.GROW_CURVE_ATTACK_105 = 1105
	r.GROW_CURVE_ATTACK_201 = 1201
	r.GROW_CURVE_ATTACK_202 = 1202
	r.GROW_CURVE_ATTACK_203 = 1203
	r.GROW_CURVE_ATTACK_204 = 1204
	r.GROW_CURVE_ATTACK_205 = 1205
	r.GROW_CURVE_ATTACK_301 = 1301
	r.GROW_CURVE_ATTACK_302 = 1302
	r.GROW_CURVE_ATTACK_303 = 1303
	r.GROW_CURVE_ATTACK_304 = 1304
	r.GROW_CURVE_ATTACK_305 = 1305
	r.GROW_CURVE_CRITICAL_101 = 2101
	r.GROW_CURVE_CRITICAL_102 = 2102
	r.GROW_CURVE_CRITICAL_103 = 2103
	r.GROW_CURVE_CRITICAL_104 = 2104
	r.GROW_CURVE_CRITICAL_105 = 2105
	r.GROW_CURVE_CRITICAL_201 = 2201
	r.GROW_CURVE_CRITICAL_202 = 2202
	r.GROW_CURVE_CRITICAL_203 = 2203
	r.GROW_CURVE_CRITICAL_204 = 2204
	r.GROW_CURVE_CRITICAL_205 = 2205
	r.GROW_CURVE_CRITICAL_301 = 2301
	r.GROW_CURVE_CRITICAL_302 = 2302
	r.GROW_CURVE_CRITICAL_303 = 2303
	r.GROW_CURVE_CRITICAL_304 = 2304
	r.GROW_CURVE_CRITICAL_305 = 2305
	return r
}
