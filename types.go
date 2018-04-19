package golambda_helper

type Header struct {
	ContentType              string `json:"Content-Type"`
	AccessControlAllowOrigin string `json:"Access-Control-Allow-Origin"`
	Location                 string `json:"Location"`
}

type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
	Header     Header `json:"headers"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ReturnObjectShopName struct {
	ShopName ShopName `json:"shopname"`
}

type ReturnObjectShopNames struct {
	ShopName []ShopName `json:"shopnames"`
}

type ShopName struct {
	Id           string `json:"id"`
	FriendlyName string `json:"friendly_name,omitempty"`
	ShopName     string `json:"shop_name,omitempty"`
	CreateDate   string `json:"create_date,omitempty"`
	Deleted      string `json:"deleted,omitempty"`
}
