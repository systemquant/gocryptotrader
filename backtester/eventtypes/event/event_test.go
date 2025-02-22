package event

import (
	"strings"
	"testing"
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	gctkline "github.com/thrasher-corp/gocryptotrader/exchanges/kline"
)

func TestEvent_AppendWhy(t *testing.T) {
	t.Parallel()
	e := &Base{}
	e.AppendReason("test")
	y := e.GetReason()
	if !strings.Contains(y, "test") {
		t.Error("expected test")
	}
}

func TestEvent_GetAssetType(t *testing.T) {
	t.Parallel()
	e := &Base{
		AssetType: asset.Spot,
	}
	if y := e.GetAssetType(); y != asset.Spot {
		t.Error("expected spot")
	}
}

func TestEvent_GetExchange(t *testing.T) {
	t.Parallel()
	e := &Base{
		Exchange: "test",
	}
	if y := e.GetExchange(); y != "test" {
		t.Error("expected test")
	}
}

func TestEvent_GetInterval(t *testing.T) {
	t.Parallel()
	e := &Base{
		Interval: gctkline.OneMin,
	}
	if y := e.GetInterval(); y != gctkline.OneMin {
		t.Error("expected one minute")
	}
}

func TestEvent_GetTime(t *testing.T) {
	t.Parallel()
	tt := time.Now()
	e := &Base{
		Time: tt,
	}
	y := e.GetTime()
	if !y.Equal(tt) {
		t.Errorf("expected %v", tt)
	}
}

func TestEvent_IsEvent(t *testing.T) {
	t.Parallel()
	e := &Base{}
	if y := e.IsEvent(); !y {
		t.Error("it is an event")
	}
}

func TestEvent_Pair(t *testing.T) {
	t.Parallel()
	e := &Base{
		CurrencyPair: currency.NewPair(currency.BTC, currency.USDT),
	}
	y := e.Pair()
	if y.IsEmpty() {
		t.Error("expected currency")
	}
}
