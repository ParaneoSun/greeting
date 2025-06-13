package greeting

import "time"

func Do() string {
	//現在時刻を取得
	current_time := time.Now()
	var greet_str string

	switch {
	case current_time.Hour() >= 4 && current_time.Hour() < 10:
		greet_str = "おはよう" + string(current_time.Hour())
	case current_time.Hour() >= 10 && current_time.Hour() < 17:
		greet_str = "こんにちは" + string(current_time.Hour())
	case current_time.Hour() >= 17:
		greet_str = "こんばんは" + string(current_time.Hour())
	}
	return greet_str
}
