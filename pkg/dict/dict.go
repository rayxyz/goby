package dict

// Signs
const (
	Blank             = ""
	Space             = " "
	Slash             = "/"
	Dash              = "-"
	DashCN            = "——"
	DashCNHalf        = "—"
	Comma             = ","
	MonoQuote         = "'"
	Quote             = "\""
	Semicolon         = ";"
	Dot               = "."
	QuestionMark      = "?"
	ExclamationMark   = "!"
	DollarSymbol      = "$"
	AtSignSymbol      = "@"
	NumberSymbol      = "#"
	AndSymbol         = "&"
	SteriodSymbol     = "*"
	PlusSymbol        = "+"
	SubtractSymbol    = "-"
	EquivalenceSymbol = "="
	LeftArrow         = "<-"
	RightArrow        = "->"
	UpArrow           = ""
	Underscore        = "_"
)

//
const (
	MaxFileSize     = 1 << 20 * 10
	MaxUploadMemory = 1 << 20 * 100
)

// User
const (
	UserSexMale   = 1
	UserSexFemale = 2

	UserStateActive = 1
	UserStateLocked = 2
)

// Logging
const (
	LogTypeNormal = "normal"
	LogTypeError  = "error"
)

// For API Gateway
const (
	RateLimitRuleUnitMinute = 1
	RateLimitRuleUnitHour   = 2
	RateLimitRuleUnitDay    = 3
	RateLimitRuleUnitWeek   = 4
	RateLimitRuleUnitMonth  = 5
)

// RateLimitRuleUnitMap :
var RateLimitRuleUnitMap = map[int32]string{
	RateLimitRuleUnitMinute: "minute",
	RateLimitRuleUnitHour:   "hour",
	RateLimitRuleUnitDay:    "day",
	RateLimitRuleUnitWeek:   "week",
	RateLimitRuleUnitMonth:  "month",
}

// i18n
const (
	LangZHCN = "zh-CN"
	LangENUS = "en-US"

	TranslationStateALL          = 0
	TranslationStateUntranslated = 1
	TranslationStateTranslated   = 2
)
