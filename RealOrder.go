package MsgStrategy

import (
	"github.com/shiguantian/wsex"
	"github.com/shiguantian/wsex/factory"
)

var Gate = factory.NewExchange(wsex.GateIo, wsex.Options{
	ExchangeName: wsex.GateIo,
	SecretKey:    "205c43b9b12545300b6c33eea9bcb4ccefba57468a8da4cbbbc11578a0a4bed3",
	AccessKey:    "9ac1fd4b82b1a69511bbed99a984f435",
	ProxyUrl: "http://127.0.0.1:4780",
})

var Huobi = factory.NewExchange(wsex.Huobi, wsex.Options{
	ExchangeName: wsex.Huobi,
	SecretKey:    "7fb1dffb-e29f3c17-0abaaa3b-ad2b4",
	AccessKey:    "d76da907-6a9b54b2-307d66d7-b1rkuf4drg",
	ProxyUrl: "http://127.0.0.1:4780",
})
