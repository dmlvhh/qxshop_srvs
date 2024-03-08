package main

import (
	"qxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn

func Init() {
	var err error
	//创建与给定目标（服务端）的连接句柄。 //grpc.WithInsecure已弃用,需要使用grpc.WithTransportCredentials(insecure.NewCredentials()
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	brandClient = proto.NewGoodsClient(conn)
}
func main() {
	Init()
	//TestGetBrandList()
	//TestGetCategoryList()
	TestGetSubCategoryList()
	conn.Close()
}
