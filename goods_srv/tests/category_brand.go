package main

import (
	"context"
	"qxshop_srvs/goods_srv/proto"
)

func TestGetCategoryBrandList() {
	rsp, err := brandClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{})
	if err != nil {
		panic(err)
	}
	println(rsp.Total)

	for _, v := range rsp.Data {
		println(v.Category.Name, v.Brand.Name)
	}
}
