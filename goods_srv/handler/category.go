package handler

import (
	"context"
	"encoding/json"
	"qxshop_srvs/goods_srv/global"
	"qxshop_srvs/goods_srv/model"
	"qxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

// GetAllCategorysList 获取所有商品分类
func (s *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	/* 获取分类的时候构造好返回的json对象，供前端使用
	   	所以在CategoryListResponse结构中专门定义了JsonData用于返回给前端使用
	   	为什么在srv层来实现，因为srv层有gorm，而web层没有gorm并不与数据库交互
	   	如果在web层实现，没有gorm处理起来会比较复杂，所以建议放在srv层来实现
	   [
	   	{
	   		"id":xxx,
	   		"name":"",
	   		"level":1,
	   		"is_tab":false,
	   		"parent":13xxx,
	   		"sub_category":[
	   			"id":xxx,
	   			"name":"",
	   			"level":1,
	   			"is_tab":false,
	   			"sub_category":[]
	   		]
	   	}
	   ]
	*/
	var categories []model.Category
	// SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	// 配置指明了外键后，可以使用Preload预加载，来把品牌的子分类也取出来
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categories)
	b, _ := json.Marshal(&categories)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

// GetSubCategory 获取商品子分类
func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		ParentCategory: category.ParentCategoryID,
		Level:          category.Level,
		IsTab:          category.IsTab,
	}
	var subCategories []model.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	preloads := "SubCategory"
	if category.Level == 1 {
		preloads = "SubCategory.SubCategory"
	}
	global.DB.Preload(preloads).Where(&model.Category{ParentCategoryID: req.Id}).Find(&subCategories)
	for _, subCategory := range subCategories {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			ParentCategory: subCategory.ParentCategoryID,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
		})
	}
	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}

// CreateCategory 创建商品分类
func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}
	category.Name = req.Name
	category.Level = req.Level
	if req.Level != 1 {
		//去查询父类目是否存在
		category.ParentCategoryID = req.ParentCategory
	}
	category.IsTab = req.IsTab
	global.DB.Save(&category)
	return &proto.CategoryInfoResponse{Id: category.ID}, nil
}

// DeleteCategory 删除商品分类
func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}

// UpdateCategory 更新商品分类
func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category

	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}

	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &emptypb.Empty{}, nil
}
