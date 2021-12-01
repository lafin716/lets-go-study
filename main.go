package main

import (
	"exam/calendar"
	"exam/crawler"
	"exam/datafile"
	"exam/magazine"
	"exam/prose"
	"fmt"
	"log"
)

func main() {
	result := prose.JoinWithCommas([]string{"가", "나", "다", "라", "마"})

	fmt.Println(result)
}

func goRoutine() {
	channel := make(chan crawler.Page)
	go crawler.GetContentLength("https://firstmall.kr", channel)
	go crawler.GetContentLength("https://google.com", channel)
	go crawler.GetContentLength("https://naver.com", channel)
	go crawler.GetContentLength("https://notion.so", channel)
	go crawler.GetContentLength("https://github.com", channel)

	fmt.Println(1, <-channel)
	fmt.Println(2, <-channel)
	fmt.Println(3, <-channel)
	fmt.Println(4, <-channel)
	fmt.Println(5, <-channel)
	fmt.Println(6, <-channel)
}

func calender() {
	date := calendar.Date{}
	err := date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}

func address() {
	var employee magazine.Employee
	employee.Name = "조일현"
	employee.Salary = 60000
	employee.Address = magazine.Address{
		Street:     "신곡동",
		State:      "경기도",
		PostalCode: "01235",
		City:       "의정부",
	}
	fmt.Println(employee)

	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}

func filePrint() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Print with slice")
	printWithSlice(lines)

	fmt.Println("Print with map")
	printWithMap(lines)
}

func printWithMap(lines []string) {
	var ranks map[string]int
	ranks = make(map[string]int)
	var names []string

	for _, line := range lines {
		ranks[line]++
		names = append(names, line)
	}

	//fmt.Println(ranks)
	for key, value := range ranks {
		fmt.Printf("%s : %d\n", key, value)
	}
}

func printWithSlice(lines []string) {
	var names []string
	var counts []int

	for _, line := range lines {
		matched := false

		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}
