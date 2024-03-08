package main

import (
	"context"
	"fmt"
	"qxshop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient

func TestGetBrandList() {
	rsp, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{})
	fmt.Println(rsp)
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	for _, brand := range rsp.Data {
		fmt.Println(brand.Name)
	}
}
