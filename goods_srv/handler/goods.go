package handler

import (
	"qxshop_srvs/goods_srv/proto"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

//// 商品接口
//func (GoodsServer) GoodsList(context.Context, *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
//
//}
//
//// 现在用户提交订单有多个商品，你得批量查询商品的信息吧
//func (GoodsServer) BatchGetGoods(context.Context, *BatchGoodsIdInfo) (*GoodsListResponse, error)
//func (GoodsServer) CreateGoods(context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error)
//func (GoodsServer) DeleteGoods(context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error)
//func (GoodsServer) UpdateGoods(context.Context, *CreateGoodsInfo) (*emptypb.Empty, error)
//func (GoodsServer) GetGoodsDetail(context.Context, *GoodInfoRequest) (*GoodsInfoResponse, error)
