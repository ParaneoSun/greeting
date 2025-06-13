package greeting

import "time"

func Do() string {
	//現在時刻を取得
	current_time := time.Now()
	var greet_str string

	switch {
	case current_time.Hour() >= 4 && current_time.Hour() < 10:
		greet_str = "おはよう" + current_time.String()
	case current_time.Hour() >= 10 && current_time.Hour() < 17:
		greet_str = "こんにちは" + current_time.String()
	case current_time.Hour() >= 17:
		greet_str = "こんばんは" + current_time.String()
	}
	return greet_str
}
