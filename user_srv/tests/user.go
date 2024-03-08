package main

import (
	"context"
	"fmt"
	"qxshop_srvs/user_srv/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	//创建与给定目标（服务端）的连接句柄。 //grpc.WithInsecure已弃用,需要使用grpc.WithTransportCredentials(insecure.NewCredentials()
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}
func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.PassWord)
		checkRsp, err := userClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password:          "test123",
			EncryptedPassword: user.PassWord,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}

func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("qingxin%d", i),
			Mobile:   fmt.Sprintf("131456783%d", i),
			PassWord: "test123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}
func main() {
	Init()
	TestGetUserList()
	//TestCreateUser()
	conn.Close()
}
