package main

import "fmt"

type Income interface {
	calculate() int
	source() string
}

type FixedBilling struct {
	projectName  string
	biddedAcount int
}

type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

type Advertisement struct {
	adName    string
	CPC       int
	noOfClick int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAcount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func (ad Advertisement) calculate() int {
	return ad.noOfClick * ad.CPC
}

func (ad Advertisement) source() string {
	return ad.adName
}

func calculateNetIncome(ic []Income) {
	var netincome int = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netincome += income.calculate()
	}
	fmt.Printf("Net income of oranization = $%d\n", netincome)
}

func main() {
	project1 := FixedBilling{projectName: "卖水果", biddedAcount: 1000}
	project2 := FixedBilling{projectName: "自媒体", biddedAcount: 2000}
	project3 := TimeAndMaterial{projectName: "外包应用开发", noOfHours: 18, hourlyRate: 200}
	bannerAd := Advertisement{adName: "豆瓣咖啡 banner", noOfClick: 1000, CPC: 5}
	screenAd := Advertisement{adName: "知乎开屏广告", noOfClick: 500, CPC: 3}
	incomeStreams := []Income{project1, project2, project3, bannerAd, screenAd}
	calculateNetIncome(incomeStreams)
}
