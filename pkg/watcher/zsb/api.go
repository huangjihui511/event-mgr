package zsb

import (
	"context"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var headerMap = map[string]string{}

type ZSBExchangeResult struct {
	USDSell float64
	USDBuy  float64
}

func getZSExchangeRatio(ctx context.Context) (r ZSBExchangeResult, err error) {
	// 	curl 'http://fx.cmbchina.com/hq/' \
	// -X 'GET' \
	// -H 'Cookie: browsehistory=%7B%22titlearray%22%3A%5B%22%u4E2A%u4EBA%u7ED3%u6C47%uFF0F%u8D2D%u6C47%u4E1A%u52A1%22%2C%22%u5883%u5916%u6C47%u6B3E%22%5D%7D; ASP.NET_SessionId=m3uyutckrxvacxotex1lrdjq' \
	// -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8' \
	// -H 'Upgrade-Insecure-Requests: 1' \
	// -H 'Host: fx.cmbchina.com' \
	// -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15' \
	// -H 'Accept-Language: en-US,en;q=0.9' \
	// -H 'Accept-Encoding: gzip, deflate' \
	// -H 'Connection: keep-alive'

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://fx.cmbchina.com/hq/", nil)
	// if err != nil {
	// 	return
	// }
	// for k, v := range headerMap {
	// 	req.Header.Set(k, v)
	// }
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	return
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return
	// }
	// parsed := make(map[string]interface{})
	// err = xml.Unmarshal(body, &parsed)
	// if err != nil {
	// 	return
	// }
	// return
	doc, err := goquery.NewDocument("http://fx.cmbchina.com/hq/")
	if err != nil {
		return
	}
	_ = doc.Find("#realRateInfo").Each(func(i int, s *goquery.Selection) {
		_ = s.Find("table").Each(func(it int, st *goquery.Selection) {
			_ = st.Find("tr").Each(func(itr int, str *goquery.Selection) {
				if itr == 0 {
					return
				}
				_ = str.Find("td.fontbold").Each(func(itrf int, strf *goquery.Selection) {
					name := strf.Text()
					name = trim(name)
					if name != "美元" {
						return
					}
					_ = str.Find("td.numberright").Each(func(itrn int, strn *goquery.Selection) {
						price := strn.Text()
						price = trim(price)
						p, _ := strconv.ParseFloat(price, 64)
						switch itrn {
						case 0:
							r.USDBuy = p / 100
						case 2:
							r.USDSell = p / 100
						}
					})
				})
			})
		})
	})
	return
}

func trim(s string) string {
	s = strings.Trim(s, "\n")
	s = strings.Trim(s, " ")
	s = strings.Trim(s, "\n")
	return s
}
