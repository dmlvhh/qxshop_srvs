package main

import (
	"context"
	"fmt"
	"qxshop_srvs/goods_srv/proto"

	"github.com/golang/protobuf/ptypes/empty"
)

func TestGetCategoryList() {
	rsp, err := brandClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Total)
	fmt.Println(rsp.JsonData)
}
func TestGetSubCategoryList() {
	rsp, err := brandClient.GetSubCategory(context.Background(), &proto.CategoryListRequest{
		Id: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.SubCategorys)
}
