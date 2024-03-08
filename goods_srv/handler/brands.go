package handler

import (
	"context"
	"qxshop_srvs/goods_srv/global"
	"qxshop_srvs/goods_srv/model"
	"qxshop_srvs/goods_srv/proto"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 品牌列表
func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	brandListResponse := proto.BrandListResponse{}
	var brands []model.Brands
	result := global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	var total int64
	global.DB.Model(&model.Brands{}).Count(&total)
	brandListResponse.Total = int32(total)

	var brandResponses []*proto.BrandInfoResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	brandListResponse.Data = brandResponses
	return &brandListResponse, nil
}

// 新建品牌
func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	//新建品牌 需要先查询品牌是否已存在
	if result := global.DB.Where("name=?", req.Name).First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	brand := &model.Brands{
		Name: req.Name,
		Logo: req.Logo,
	}
	global.DB.Save(brand)
	return &proto.BrandInfoResponse{Id: brand.ID}, nil
}

// 品牌删除
func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

// 品牌更新
func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	brands := model.Brands{}
	if result := global.DB.First(&brands); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	global.DB.Save(&brands)
	return &emptypb.Empty{}, nil
}
