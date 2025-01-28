package main

import (
	"github.com/mohanson/godump/pretty"
)

func main() {
	pretty.PrintTable([][]string{
		{"coin", "usable", "freeze"},
		{"bnb", "3.4247", "15.1042"},
		{"btc", "0.0356", "0.2018"},
		{"ckb", "526333.5353", "3160364.7891"},
		{"eth", "0.4864", "2.5437"},
		{"trx", "7498.283690", "33516.101561"},
		{"usdt", "19494.875228", "71498.150403"},
	})
}
