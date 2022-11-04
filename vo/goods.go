package vo

type RequestCreateGoods struct {
	GoodsName    string `json:"goods_name" binding:"required"`
	FirstPicture string `json:"first_picture" binding:"required"`
	Introduce    string `json:"introduce" binding:"required"`
	Price        uint   `json:"price" binding:"required"`
	Count        uint   `json:"count" binding:"required"`
	CreateUser   uint   `json:"createUser"`
}
