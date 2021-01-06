package main

import (
	"fmt"
	"time"
)

func mainTimeOne() {
	nowTime := time.Now()
	fmt.Println("当前系统时间：", nowTime)

	// 年、月、日
	year := nowTime.Year()
	fmt.Println(year)
	month := nowTime.Month()
	fmt.Println(month)
	year1, month1, day := nowTime.Date()
	fmt.Println("年、月、日：", year1, month1, day)

	// 周月年中的第几天
	day = nowTime.Day()
	yearDay := nowTime.YearDay()
	weekday := nowTime.Weekday()
	fmt.Println(day, yearDay, weekday)

	// 时分秒
	fmt.Println(nowTime.Hour())
	fmt.Println(nowTime.Minute())
	fmt.Println(nowTime.Second())

	data := time.Date(2020, time.September, 8, 15, 0, 0, 0, time.Now().Location())
	fmt.Println(data)
}

func mainTimeTwo() {
	now := time.Now()

	// 一天之前
	duration, _ := time.ParseDuration("-24h0m0s")
	fmt.Println(now.Add(duration))

	// 一周之前
	fmt.Println(now.Add(duration * 7))

	// 一月之前
	fmt.Println(now.Add(duration * 30))

	// 计算时间差
	fmt.Println("时间差：", now.Sub(now.Add(duration)))
}
