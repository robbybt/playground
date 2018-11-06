package main

import (
	"fmt"
	"github.com/josephburnett/jd/lib"
)

func main() {
	jsonA := `[{"body":{"foo":1,"bar":3},"header":{"qwe":10,"asd":20}}]`
	jsonB := `[{"body":{"foo":1,"bar":3},"header":{"qwe":11,"asd":20}},{"body":{"foo":1,"bar":3},"header":{"qwe":11,"asd":20}}]`
	diff1,_ := GetDifferenceJson(jsonA,jsonB)
	jsonC := `[{"body":{"bar":3},"header":{"qwe":11}},{"body":{"foo":1,"bar":3},"header":{"qwe":11,"asd":20}}]`
	diff2,_ := GetDifferenceJson(jsonB,jsonC)

	//fmt.Println(html.EscapeString(diff1))
	fmt.Println(jsonC)
	comebackJson3,_ := ApplyChangeFromDifferences(jsonC,[]string{diff2,diff1})
	//comebackJson,_ := ApplyChangeFromDifference(jsonC,diff2)
	//fmt.Println(comebackJson)
	//comebackJson2,_ := ApplyChangeFromDifference(comebackJson,diff1)
	fmt.Println(comebackJson3)
}

func GetDifferenceJson(initialJson,changedJson string) (string,error) {
	nodeInitial,err:= jd.ReadJsonString(initialJson)
	if err != nil {
		return "",err
	}
	nodeChanged,err:= jd.ReadJsonString(changedJson)
	if err != nil {
		return "",err
	}
	return nodeChanged.Diff(nodeInitial).Render(),nil
}

func ApplyChangeFromDifference(jsonString,diffString string) (string,error) {
	nodeJson,err := jd.ReadJsonString(jsonString)
	if err != nil {
		return "",err
	}
	diff, err := jd.ReadDiffString(diffString)
	if err != nil {
		return "",err
	}
	appliedJson,err := nodeJson.Patch(diff)
	if err != nil {
		return "",err
	}
	return appliedJson.Json(),nil
}

func ApplyChangeFromDifferences(jsonString string ,diffStrings []string) (string,error) {
	result := jsonString
	var err error
	for _,diffString := range diffStrings {
		result,err = ApplyChangeFromDifference(result,diffString)
		if err != nil {
			return "",err
		}
	}
	return result,nil
}

//func ExampleJsonNode_Diff() string {
//
//	"asd"
//	a, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"campaignID":400,"discount_percentage":20,"discounted_price":"Rp 80.000","original_price":"Rp 100.000","custom_stock":9999,"campaign_status":7,"start_date":"2018-10-17 14:15:00","end_date":"2018-10-24 13:41:00","campaign_type_name":"Flash Sale","discounted_price_fmt":"80000","original_price_fmt":"100000","apps_only":true,"applinks":"https://tokopedia.link/c8eLw8pIWQ","final_price":80000,"seller_price":300,"tokopedia_subsidy":100}}]`)
//
//	b, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"applinks":"https://tokopedia.link/c8eLw8pIWQ"}}]`)
//	stringjson := b.Diff(a).Render()
//	diff, _ := jd.ReadDiffString(stringjson)
//
//	c, _ := b.Patch(diff)
//	fmt.Println(c)
//	return "a"
//	// Output:
//	// @ ["foo"]
//	// - "bar"
//	// + "baz"
//}
//
//func ExampleJsonNode_Patch(s string) {
//	a, _ := jd.ReadJsonString(`[{"product_id":15221806,"campaign":{"campaignID":400,"discount_percentage":20,"discounted_price":"Rp 80.000","original_price":"Rp 100.000","custom_stock":9999,"campaign_status":7,"start_date":"2018-10-17 14:15:00","end_date":"2018-10-24 13:41:00","campaign_type_name":"Flash Sale","discounted_price_fmt":"80000","original_price_fmt":"100000","apps_only":true,"applinks":"https://tokopedia.link/c8eLw8pIWQ"}}]`)
//	diff, _ := jd.ReadDiffString(s)
//	b, _ := a.Patch(diff)
//	fmt.Print(b)
//}
