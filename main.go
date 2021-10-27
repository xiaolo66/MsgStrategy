package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	Log.Infof("==========机器人进程启动==========")
	Res.SetConfig(SetMethod("GET"),
		SetUrl("https://www.binancezh.co/bapi/composite/v1/public/cms/article/catalog/list/query?catalogId=48&pageNo=1&pageSize=15"),
		SetHeader("lang", "zh-CN"),
		SetBody(""))
	//var Once =sync.Once{}

	err := InitLog(".logs", "", time.Duration(7)*time.Hour*24, time.Duration(1)*time.Hour*24)
	if err != nil {
		return
	}
	var RawCurrency = make([]string, 0)
	var TarCurrency = make([]string, 0)
	// go func() {
	// 	time.Sleep(time.Minute * 3)
	// 	fmt.Println(1)
	// 	RawCurrency = append(RawCurrency, "EOS")
	// }()
	for {
		time.Sleep(time.Second * 20)
		announce, err := Res.GetMsg()
		if err != nil {
			Log.Errorf("Response failed err: %v", err)
			Log.Infof("Response failed err: %v", err)
			continue
		}
		Currency, err := ParseData(announce)
		if err != nil {
			Log.Errorf("Parser Currency failed err:%v", err)
			return
		}
		Log.Infof("Parser Currency success,data:%v", Currency)
		if len(RawCurrency) == 0 {
			for _, v := range Currency {
				RawCurrency = append(RawCurrency, v)
			}
			Log.Infof("Get Rawcurrency :%v", RawCurrency)
		}
		for _, value := range Currency {
			if !SliceContain(value, RawCurrency) {
				TarCurrency = append(TarCurrency, value)
			}
		}
		if len(TarCurrency) != 0 {
			Log.Infof("币安新上币：%v", TarCurrency)
			for _, cur := range TarCurrency {
				symbol := fmt.Sprintf("%s/USDT", strings.ToUpper(cur))
				fmt.Println(symbol)
				Wg.Add(2)
				go RealOrder(symbol, Huobi)
				go RealOrder(symbol, Gate)
			}
			Wg.Wait()
			break
		}
	}
}
