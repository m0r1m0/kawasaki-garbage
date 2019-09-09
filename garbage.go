package kawasakigarbage

import "time"

// GarbageType g
type GarbageType int

const (
	// Normal is 普通ゴミ
	Normal GarbageType = iota
	// Bottle is ペットボトル
	Bottle
	// MixedPaper is ミックスペーパー
	MixedPaper
	// Plastic is プラスチック
	Plastic
	// Cardboard ダンボール
	Cardboard
)

// GetGarbageType return garbage type
func GetGarbageType(weekday time.Weekday) []GarbageType {
	switch weekday {
	case time.Monday:
		return []GarbageType{Plastic}
	case time.Wednesday, time.Saturday:
		return []GarbageType{Normal}
	case time.Thursday:
		return []GarbageType{MixedPaper, Cardboard}
	case time.Friday:
		return []GarbageType{Bottle}
	}
	return []GarbageType{}
}

func (t GarbageType) String() string {
	switch t {
	case Normal:
		return "普通ゴミ"
	case Bottle:
		return "ペットボトル"
	case MixedPaper:
		return "ミックスペーパー"
	case Plastic:
		return "プラスチック"
	case Cardboard:
		return "ダンボール"
	default:
		return "なし"
	}
}
