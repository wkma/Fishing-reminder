package main

import (
	"fmt"
	"time"
)

var time_now string
var time_sat string
var time_year string
var time_spring string

//判断上下午
func judgeMorningOrAfternoon(timeHour int) string {
	if timeHour >= 12 {
		return "下午好"
	} else {
		return "上午好"
	}
}

//获取相差时间
func getDayDiffer(start_time, end_time string) int64 {
	var day int64
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		day = diff/3600/24 + 1
		return day
	} else {
		return day
	}
}

//打印信息
func fishingReminder(time_year, time_spring string) {
	/*传入周六，元旦，春节时间*/
	time_now := time.Now().Format("2006-01-02 15:04:05")
	hour := time.Now().Hour()
	today := time.Now().Weekday()
	day := time.Saturday - today
	fmt.Print("摸鱼办提醒您：\n")
	fmt.Printf("%s%s，摸鱼人！\n", time_now, judgeMorningOrAfternoon(hour))
	fmt.Println("工作再累 一定不要忘记摸鱼哦！")
	fmt.Println("有事没事起身去茶水间去厕所去廊道走走")
	fmt.Println("别老在工位上坐着钱是老板的,但命是自己的")
	fmt.Printf("距离周末假期还有%d天\n", day)
	fmt.Printf("距离元旦假期还有%d天\n", getDayDiffer(time_now, time_year))
	fmt.Printf("距离春节假期还有%d天\n", getDayDiffer(time_now, time_spring))
	fmt.Println("上班是帮老板赚钱，摸鱼是赚老板的钱！\n最后，祝愿天下所有摸鱼人，都能愉快的渡过每一天.")
}
func main() {
	time_year = "2022-01-01 00:00:00"
	time_spring = "2022-02-01 00:00:00"
	fishingReminder(time_year, time_spring)
	time.Sleep(time.Second*30)
}
