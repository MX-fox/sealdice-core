package dice

import (
	"encoding/json"
)

var coc7TemplateData string = `
{
  "keyName": "coc7",
  "fullName": "克苏鲁的呼唤第七版",
  "authors": [
    "木落",
    "月森优姬",
    "于言诺"
  ],
  "version": "1.0.0",
  "updatedTime": "20230212",
  "templateVer": "1.0",
  "nameTemplate": {
    "coc": {
      "template": "{$t玩家_RAW} SAN{理智} HP{生命值}/{生命值上限} DEX{敏捷}",
      "helpText": "自动设置coc名片"
    },
    "cocL": {
      "template": "{$t玩家_RAW} san{理智} hp{生命值}/{生命值上限} dex{敏捷}",
      "helpText": "自动设置coc名片，小写"
    }
  },
  "attrSettings": {
    "top": [
      "力量",
      "敏捷",
      "体质",
      "体型",
      "外貌",
      "智力",
      "意志",
      "教育",
      "理智",
      "db",
      "克苏鲁神话",
      "生命值",
      "魔法值"
    ],
    "sortBy": "name",
    "ignores": [
      "生命值上限"
    ],
    "showAs": {
      "生命值": "{生命值}/{生命值上限}",
      "魔法值": "{魔法值}/{魔法值上限}"
    },
    "setter": null
  },
  "defaults": {
    "乔装": 5,
    "书法": 5,
    "人类学": 1,
    "会计": 5,
    "伪造": 5,
    "估价": 5,
    "侦查": 25,
    "信用评级": 0,
    "催眠": 1,
    "克苏鲁神话": 0,
    "写作": 5,
    "冲锋枪": 15,
    "制陶": 5,
    "剑": 20,
    "动物学": 1,
    "动物驯养": 5,
    "化学": 1,
    "医学": 1,
    "博物学": 10,
    "历史": 5,
    "厨艺": 5,
    "取悦": 15,
    "司法科学": 1,
    "吹真空管": 5,
    "喜剧": 5,
    "器乐": 5,
    "园艺": 5,
    "图书馆使用": 20,
    "地质学": 1,
    "声乐": 5,
    "天文学": 1,
    "妙手": 10,
    "学识": 1,
    "密码学": 1,
    "导航": 10,
    "工程学": 1,
    "弓": 15,
    "心理学": 10,
    "急救": 30,
    "恐吓": 15,
    "手枪": 20,
    "打字": 5,
    "技术制图": 5,
    "投掷": 20,
    "摄影": 5,
    "操作重型机械": 1,
    "攀爬": 20,
    "数学": 10,
    "斗殴": 25,
    "斧": 15,
    "日本刀": 20,
    "木匠": 5,
    "机枪": 10,
    "机械维修": 10,
    "极地": 10,
    "植物学": 1,
    "歌剧歌唱": 5,
    "气象学": 1,
    "汽车驾驶": 20,
    "沙漠": 10,
    "法律": 5,
    "海洋": 10,
    "游泳": 20,
    "潜水": 1,
    "潜行": 20,
    "火焰喷射器": 10,
    "炮术": 1,
    "爆破": 1,
    "物理学": 1,
    "理发": 5,
    "生存": 10,
    "生物学": 1,
    "电子学": 1,
    "电气维修": 10,
    "矛": 20,
    "神秘学": 5,
    "科学": 1,
    "粉刷匠和油漆工": 5,
    "精神分析": 1,
    "绞索": 15,
    "美术": 5,
    "考古学": 1,
    "耕作": 5,
    "聆听": 20,
    "舞蹈": 5,
    "船": 1,
    "艺术与手艺": 5,
    "药学": 1,
    "莫里斯舞蹈": 5,
    "表演": 5,
    "裁缝": 5,
    "计算机使用": 5,
    "话术": 5,
    "语言": 1,
    "说服": 10,
    "读唇": 1,
    "跳跃": 20,
    "连枷": 10,
    "追踪": 10,
    "速记": 5,
    "重武器": 10,
    "链锯": 10,
    "锁匠": 1,
    "雕塑": 5,
    "霰弹枪": 25,
    "鞭": 5,
    "飞行器": 1,
    "骑术": 5
  },
  "defaultsComputed": {
    "db": "(力量 + 体型) \u003c 65 ? -2, (力量 + 体型) \u003c 85 ? -1, (力量 + 体型) \u003c 125 ? 0, (力量 + 体型) \u003c 165 ? 1d14, (力量 + 体型) \u003c 205 ? 1d6, 1 ? ((力量 + 体型 - 205) / 80 + 2)d6",
    "母语": "教育",
    "生命值上限": "(体质 + 体型) / 10",
    "语言": "教育",
    "闪避": "敏捷 / 2"
  },
  "alias": {
    "db": [
      "DB",
      "伤害加值"
    ],
    "乔装": [
      "喬裝"
    ],
    "书法": [
      "書法"
    ],
    "人类学": [
      "人類學"
    ],
    "会计": [
      "會計"
    ],
    "伪造": [
      "偽造"
    ],
    "估价": [
      "估價"
    ],
    "体型": [
      "siz",
      "體型",
      "体形",
      "體形"
    ],
    "体质": [
      "con",
      "體質"
    ],
    "侦查": [
      "侦察",
      "偵查",
      "偵察"
    ],
    "信用评级": [
      "信誉",
      "信用",
      "信誉度",
      "cr",
      "信用評級",
      "信譽",
      "信譽度"
    ],
    "克苏鲁神话": [
      "cm",
      "克苏鲁",
      "克苏鲁神话知识",
      "克蘇魯",
      "克蘇魯神話",
      "克蘇魯神話知識"
    ],
    "写作": [
      "文学",
      "寫作",
      "文學"
    ],
    "冲锋枪": [
      "衝鋒槍"
    ],
    "剑": [
      "剑术",
      "劍",
      "劍術"
    ],
    "力量": [
      "str"
    ],
    "动物学": [
      "動物學"
    ],
    "动物驯养": [
      "驯养",
      "驯兽",
      "馴獸",
      "動物馴養",
      "馴養"
    ],
    "化学": [
      "化學"
    ],
    "医学": [
      "醫學"
    ],
    "博物学": [
      "自然",
      "自然学",
      "自然史",
      "自然學",
      "博物學"
    ],
    "历史": [
      "歷史"
    ],
    "厨艺": [
      "烹饪",
      "廚藝",
      "烹飪"
    ],
    "取悦": [
      "魅惑",
      "取悅"
    ],
    "司法科学": [
      "司法科學"
    ],
    "喜剧": [
      "喜劇"
    ],
    "器乐": [
      "器樂"
    ],
    "园艺": [
      "園藝"
    ],
    "图书馆使用": [
      "圖書館使用",
      "图书馆",
      "图书馆利用",
      "圖書館",
      "圖書館利用"
    ],
    "地质学": [
      "地理学",
      "地質學",
      "地理學"
    ],
    "声乐": [
      "聲樂"
    ],
    "外貌": [
      "app",
      "外表"
    ],
    "天文学": [
      "天文學"
    ],
    "妙手": [
      "藏匿",
      "盗窃",
      "盜竊"
    ],
    "学识": [
      "学问",
      "學識",
      "學問"
    ],
    "密码学": [
      "密碼學"
    ],
    "导航": [
      "领航",
      "領航",
      "導航"
    ],
    "工程学": [
      "工程學"
    ],
    "幸运": [
      "luck",
      "幸运值",
      "运气",
      "幸運",
      "運氣",
      "幸運值"
    ],
    "弓": [
      "弓术",
      "弓箭",
      "弓術"
    ],
    "心理学": [
      "心理學"
    ],
    "恐吓": [
      "恐嚇"
    ],
    "意志": [
      "pow"
    ],
    "手枪": [
      "手槍"
    ],
    "技术制图": [
      "技術製圖"
    ],
    "投掷": [
      "投擲"
    ],
    "护甲": [
      "装甲",
      "護甲",
      "裝甲"
    ],
    "摄影": [
      "攝影"
    ],
    "操作重型机械": [
      "重型操作",
      "重型机械",
      "重型",
      "重机",
      "操作重型機械",
      "重型機械",
      "重機"
    ],
    "攀爬": [
      "攀岩",
      "攀登"
    ],
    "敏捷": [
      "dex"
    ],
    "教育": [
      "edu",
      "知识",
      "知識"
    ],
    "数学": [
      "數學"
    ],
    "斗殴": [
      "鬥毆"
    ],
    "斧": [
      "斧头",
      "斧子",
      "斧頭"
    ],
    "智力": [
      "int",
      "灵感",
      "靈感"
    ],
    "木匠": [
      "木工"
    ],
    "机枪": [
      "機槍"
    ],
    "机械维修": [
      "机器维修",
      "机修",
      "機器維修",
      "機修",
      "機械維修"
    ],
    "极地": [
      "極地"
    ],
    "枪械": [
      "火器",
      "射击",
      "槍械",
      "射擊"
    ],
    "植物学": [
      "植物學"
    ],
    "歌剧歌唱": [
      "歌劇歌唱"
    ],
    "母语": [
      "母語"
    ],
    "气象学": [
      "氣象學"
    ],
    "汽车驾驶": [
      "汽車駕駛",
      "汽车",
      "驾驶",
      "汽車",
      "駕駛"
    ],
    "海洋": [
      "海上"
    ],
    "潜水": [
      "潛水"
    ],
    "潜行": [
      "躲藏",
      "潛行"
    ],
    "火焰喷射器": [
      "火焰噴射器"
    ],
    "炮术": [
      "炮術"
    ],
    "物理": [
      "物理学",
      "物理學"
    ],
    "理发": [
      "理髮"
    ],
    "理智": [
      "san",
      "san值",
      "理智值",
      "理智点数",
      "心智",
      "心智点数",
      "心智點數",
      "理智點數"
    ],
    "生命值": [
      "hp",
      "生命",
      "体力",
      "體力",
      "血量",
      "耐久值"
    ],
    "生命值上限": [
      "hpmax",
      "生命上限",
      "体力上限",
      "體力上限",
      "血量上限",
      "耐久值上限"
    ],
    "生物学": [
      "生物學"
    ],
    "电子学": [
      "電子學"
    ],
    "电气维修": [
      "电器维修",
      "电工",
      "電氣維修",
      "電器維修",
      "電工"
    ],
    "矛": [
      "投矛"
    ],
    "神秘学": [
      "神秘學"
    ],
    "科学": [
      "科學"
    ],
    "精神分析": [
      "心理分析"
    ],
    "绞索": [
      "绞具",
      "絞索",
      "絞具"
    ],
    "美术": [
      "美術"
    ],
    "考古学": [
      "考古學"
    ],
    "聆听": [
      "聆聽"
    ],
    "船": [
      "开船",
      "驾驶船",
      "開船",
      "駕駛船"
    ],
    "艺术与手艺": [
      "艺术和手艺",
      "艺术",
      "手艺",
      "工艺",
      "技艺",
      "藝術與手藝",
      "藝術和手藝",
      "藝術",
      "手藝",
      "工藝",
      "技藝"
    ],
    "药学": [
      "藥學"
    ],
    "裁缝": [
      "裁縫"
    ],
    "计算机使用": [
      "电脑使用",
      "計算機使用",
      "電腦使用",
      "计算机",
      "电脑",
      "計算機",
      "電腦"
    ],
    "话术": [
      "快速交谈",
      "話術",
      "快速交談"
    ],
    "语言": [
      "外语",
      "語言",
      "外語"
    ],
    "说服": [
      "辩论",
      "议价",
      "演讲",
      "說服",
      "辯論",
      "議價",
      "演講"
    ],
    "读唇": [
      "唇语",
      "讀唇",
      "唇語"
    ],
    "跳跃": [
      "跳躍"
    ],
    "追踪": [
      "跟踪",
      "追蹤",
      "跟蹤"
    ],
    "速记": [
      "速記"
    ],
    "链枷": [
      "连枷",
      "連枷",
      "鏈枷"
    ],
    "链锯": [
      "电锯",
      "油锯",
      "鏈鋸",
      "電鋸",
      "油鋸"
    ],
    "锁匠": [
      "开锁",
      "撬锁",
      "钳工",
      "鎖匠",
      "鉗工",
      "開鎖",
      "撬鎖"
    ],
    "闪避": [
      "閃避"
    ],
    "霰弹枪": [
      "步枪",
      "步霰",
      "步枪/霰弹枪",
      "散弹枪",
      "步槍",
      "霰彈槍",
      "步霰",
      "步槍/霰彈槍",
      "散彈槍"
    ],
    "飞行器": [
      "开飞行器",
      "驾驶飞行器",
      "飛行器",
      "開飛行器",
      "駕駛飛行器"
    ],
    "骑术": [
      "騎術"
    ],
    "魔法值": [
      "mp",
      "魔法",
      "魔力",
      "魔力值"
    ],
    "魔法值上限": [
      "mpmax",
      "魔法上限",
      "蓝量上限",
      "藍量上限"
    ]
  },
  "textMap": {
    "COC": {
      "设置测试_成功": [
        [
          "设置完成",
          1
        ]
      ]
    }
  },
  "textMapHelpInfo": null
}
`

var _coc7tmpl *CharacterTemplate

func getCoc7CharTemplate() *CharacterTemplate {
	if _coc7tmpl != nil {
		return _coc7tmpl
	}

	temp := &CharacterTemplate{}
	err := json.Unmarshal([]byte(coc7TemplateData), temp)
	if err != nil {
		return nil
	}

	// 因为 `` 的冲突，所以写在这里
	temp.AttrSettings.ShowAs["db"] = "{ (力量 + 体型) \u003c 65 ? '-2', (力量 + 体型) \u003c 85 ? '-1', (力量 + 体型) \u003c 125 ? '0', (力量 + 体型) \u003c 165 ? '1d4', (力量 + 体型) \u003c 205 ? '1d6', 1 ? `{((力量 + 体型 - 205) / 80 + 2)}d6` }"
	_coc7tmpl = temp

	return temp
}
