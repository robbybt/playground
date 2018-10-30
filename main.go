package main

import (
	"fmt"

	jd "github.com/josephburnett/jd/lib"
)

func main() {
	_ = ExampleJsonNode_Diff()
	//fmt.Println(s)
	//ExampleJsonNode_Patch(s)

}

func ExampleJsonNode_Diff() string {
	a, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"campaignID":400,"discount_percentage":20,"discounted_price":"Rp 80.000","original_price":"Rp 100.000","custom_stock":9999,"campaign_status":7,"start_date":"2018-10-17 14:15:00","end_date":"2018-10-24 13:41:00","campaign_type_name":"Flash Sale","discounted_price_fmt":"80000","original_price_fmt":"100000","apps_only":true,"applinks":"https://tokopedia.link/c8eLw8pIWQ","final_price":80000,"seller_price":300,"tokopedia_subsidy":100}}]`)

	b, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"applinks":"https://tokopedia.link/c8eLw8pIWQ"}}]`)
	stringjson := b.Diff(a).Render()
	diff, _ := jd.ReadDiffString(stringjson)

	c, _ := b.Patch(diff)
	fmt.Println(c)
	return "a"
	// Output:
	// @ ["foo"]
	// - "bar"
	// + "baz"
}

func ExampleJsonNode_Patch(s string) {
	a, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"campaignID":400,"discount_percentage":20,"discounted_price":"Rp 80.000","original_price":"Rp 100.000","custom_stock":9999,"campaign_status":7,"start_date":"2018-10-17 14:15:00","end_date":"2018-10-24 13:41:00","campaign_type_name":"Flash Sale","discounted_price_fmt":"80000","original_price_fmt":"100000","apps_only":true,"applinks":"https://tokopedia.link/c8eLw8pIWQ"}}]`)
	diff, _ := jd.ReadDiffString(s)
	b, _ := a.Patch(diff)
	fmt.Print(b)
}
