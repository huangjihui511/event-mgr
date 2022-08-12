package scb

import (
	"context"
	"encoding/json"
	"huangjihui511/event-mgr/pkg/logs"
	"io/ioutil"
	"net/http"
	"strconv"
)

var headerMap = map[string]string{
	"correlationId": "BgGvcY789-VP6WneLnh",
	"auth-token":    "eyJ0eXAiOiJ3ZngiLCJhbGciOiJIUzI1NiJ9.NWIwMTM2KjMxMSoreTJlMQ.dZqxa4ePed3LXkOXYR6W7NNrLebud_8HSP3INAPFf-0",
}

type ResponseSCBank struct {
	USDCNY USDCNY `json:"USDCNY,omitempty"`
}

type USDCNY struct {
	BidSpotRate string `json:"bidSpotRate,omitempty"`
	AskSpotRate string `json:"askSpotRate,omitempty"`
}

type ExchangeRatio struct {
	BuyRatio  float64
	SellRatio float64
}

func getSCExchangeRatio(ctx context.Context) (exchangeRatio ExchangeRatio, err error) {
	// curl -k -L -s --compressed 'https://wealth.sc.com/wm/api/cn/wfx/wealthfx/orders/rate/USDCNY?cacheBuster=16601973478&class=030201' \
	// 	-X 'GET' \
	// 	-H 'Accept: */*' \
	// 	-H 'Cookie: _ga_YHRGW6LC0H=GS1.1.1660188585.18.1.1660188599.0; _hid=zu4ZCv7t22JiRG0uR1VUjgA; _ga=GA1.2.222790583.1658580478; _gid=GA1.2.480098922.1660115597; RT="z=1&dm=sc.com&si=fa45e157-61a1-4a38-813b-eb5fa8dee3eb&ss=l6na7hn4&sl=1&tt=1lc&bcn=%2F%2F684d0d45.akstat.io%2F&ld=1ld&ul=1m12&hd=1m5p"; _gcl_au=1.1.345278299.1658580478' \
	// 	-H 'Referer: https://wealth.sc.com/wm/full/cn/wfx/wealth-fx-portal/?trade=CNY&mobile=' \
	// 	-H 'Host: wealth.sc.com' \
	// 	-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15' \
	// 	-H 'Accept-Language: en-US,en;q=0.9' \
	// 	-H 'Accept-Encoding: gzip, deflate, br' \
	// 	-H 'Connection: keep-alive' \
	// 	-H 'correlationId: BgGvcY789-VP6WneLnh' \
	// 	-H 'auth-token: eyJ0eXAiOiJ3ZngiLCJhbGciOiJIUzI1NiJ9.NWIwMTM2KjMxMSoreTJlMQ.dZqxa4ePed3LXkOXYR6W7NNrLebud_8HSP3INAPFf-0'

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://wealth.sc.com/wm/api/cn/wfx/wealthfx/orders/rate/USDCNY?cacheBuster=16601973478&class=030201", nil)
	if err != nil {
		return
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	respUnmarshal := ResponseSCBank{}
	err = json.Unmarshal(body, &respUnmarshal)
	if err != nil {
		return
	}
	exchangeRatio.BuyRatio, err = strconv.ParseFloat(respUnmarshal.USDCNY.AskSpotRate, 64)
	if err != nil {
		return
	}
	exchangeRatio.SellRatio, err = strconv.ParseFloat(respUnmarshal.USDCNY.BidSpotRate, 64)
	logs.Logger.Printf("Get rate from SC Bank: %+v", exchangeRatio)
	return
}
