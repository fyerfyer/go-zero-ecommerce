// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type ListOrdersRequest struct {
	UID      int64 `form:"uid"`
	Status   int32 `form:"status,optional"`
	Cursor   int64 `form:"cursor,optional"`
	PageSize int64 `form:"ps,default=20"`
}

type ListOrdersResponse struct {
	Orders    []*Order `json:"orders"`
	IsEnd     bool     `json:"is_end"`
	OrderTime int64    `json:"order_time"`
}

type Order struct {
	OrderID            string  `json:"order_id"`
	Status             int32   `json:"status"`
	Quantity           int64   `json:"quantity"`
	Payment            float64 `json:"payment"`
	TotalPrice         float64 `json:"total_price"`
	CreateTime         int64   `json:"create_time"`
	ProductID          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductImage       string  `json:"product_image"`
	ProductDescription string  `json:"product_description"`
}
