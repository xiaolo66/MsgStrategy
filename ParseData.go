package MsgStrategy

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/shiguantian/wsex"
)

func ParseData(data string) ([]string, error) {
	reg1 := regexp.MustCompile(`"title":"(币安上市|币安创新区上市).*?",`)
	if reg1 == nil {
		return nil, fmt.Errorf("reg1 MustCompile error")
	}
	reg2 := regexp.MustCompile(`（(.*?)）`)
	if reg2 == nil {
		return nil, fmt.Errorf("reg2 MustCompile error")
	}
	Currency := make([]string, 0)
	tmpMsg := reg1.FindAllString(data, -1)
	for _, v := range tmpMsg {
		result := reg2.FindAllStringSubmatch(v, -1)
		for _, currency := range result {
			Currency = append(Currency, currency[1])
		}
	}
	return Currency, nil
}

func SliceContain(s string, r []string) bool {
	for _, value := range r {
		if s == value {
			return true
		}
	}
	return false
}

var Wg sync.WaitGroup

func RealOrder(symbol string, e wsex.IExchange) {
	markets, err := e.FetchMarkets()
	defer Wg.Done()
	if err != nil {
		Log.Errorf("Fetch  Markets failed err:%v", err)
		return
	}
	market, ok := markets[symbol]
	if !ok {
		Log.Errorf("symbol:%s not exist in %s", symbol)
		//Log.Fatalf("symbol:%s not exist in %s",symbol)
		return
	}
	Log.Infof("Symbol :%s exist in Exchange ", symbol)

	balance, err := e.FetchBalance()
	if err != nil {
		Log.Errorf("Fetch %v Balance failed err:%v", e, err)
		return
	}
	Log.Infof("Account Balcane:%v", balance)
	ticker, err := e.FetchTicker(symbol)
	if err != nil {
		Log.Errorf("Fetch %v ticker failed err:%v", symbol, err)
		return
	}
	Log.Infof("symbol: %s last price:%f", symbol, ticker.Last)
	buyPrice := ticker.Last * 1.05
	rest := balance[market.QuoteID]
	amount := (rest.Available / 1.1) / buyPrice
	fmt.Printf("%s的买入价格是%f,买入数量是%f", market.Symbol, buyPrice, amount)
	order, err := e.CreateOrder(symbol, buyPrice, amount, wsex.Buy, wsex.LIMIT, wsex.Normal, false)
	if err != nil {
		Log.Errorf("CreateOrder %s failed err: ", symbol, err)
		return
	}
	Log.Infof("CreateOrder Success %s", order.ID)

	return
}
