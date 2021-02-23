package adapters

import (
	"testing"
	"time"

	binance "github.com/adshao/go-binance/v2"

	"github.com/cryptellation/models.go"
)

type testCaseKLineToCandleStick struct {
	KLine       binance.Kline
	CandleStick models.CandleStick
}

var testCasesKLineToCandleStick = []testCaseKLineToCandleStick{
	{
		KLine:       binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "1.5"},
		CandleStick: models.CandleStick{Time: time.Unix(0, 0), Open: 1, High: 2, Low: 0.5, Close: 1.5},
	},
	{
		KLine:       binance.Kline{OpenTime: 0, Open: "2.0", High: "4.0", Low: "1", Close: "3"},
		CandleStick: models.CandleStick{Time: time.Unix(0, 0), Open: 2, High: 4, Low: 1, Close: 3},
	},
}

func TestKLineToCandleStick(t *testing.T) {
	for i, test := range testCasesKLineToCandleStick {
		cs, err := KLineToCandleStick(test.KLine)
		if err != nil {
			t.Error("There should be no error on CandleStick", i, ":", err)
		} else if test.CandleStick != cs {
			t.Error("CandleStick", i, "is not transformed correctly:", test.CandleStick, cs)
		}
	}
}

func TestKLineToCandleStick_IncorrectOpen(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "error", High: "2.0", Low: "0.5", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on open")
	}
}

func TestKLineToCandleStick_IncorrectHigh(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "error", Low: "0.5", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on high")
	}
}

func TestKLineToCandleStick_IncorrectLow(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "error", Close: "1.5"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on low")
	}
}

func TestKLineToCandleStick_IncorrectClose(t *testing.T) {
	c := binance.Kline{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "error"}
	if _, err := KLineToCandleStick(c); err == nil {
		t.Error("There should be an error on close")
	}
}

func TestKLinesToCandleSticks(t *testing.T) {
	// Only get klines
	kl := make([]*binance.Kline, len(testCasesKLineToCandleStick))
	for i := range testCasesKLineToCandleStick {
		kl[i] = &testCasesKLineToCandleStick[i].KLine
	}

	// Test function
	cs, err := KLinesToCandleSticks(kl)
	if err != nil {
		t.Error("There should be no error:", err)
	}

	for i, test := range testCasesKLineToCandleStick {
		if test.CandleStick != cs[i] {
			t.Error("CandleStick", i, "is not transformed correctly:", test.CandleStick, cs[i])
		}
	}
}

func TestKLinesToCandleSticks_IncorrectOpen(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "error", High: "2.0", Low: "0.5", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on open")
	}
}

func TestKLinesToCandleSticks_IncorrectHigh(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "error", Low: "0.5", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on high")
	}
}

func TestKLinesToCandleSticks_IncorrectLow(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "2.0", Low: "error", Close: "1.5"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on low")
	}
}

func TestKLinesToCandleSticks_IncorrectClose(t *testing.T) {
	c := []*binance.Kline{{OpenTime: 0, Open: "1.0", High: "2.0", Low: "0.5", Close: "error"}}
	if _, err := KLinesToCandleSticks(c); err == nil {
		t.Error("There should be an error on close")
	}
}
