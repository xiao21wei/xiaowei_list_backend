package response

type CommonA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type RegisterQ struct {
	Username string `json:"username" binding:"min=5,max=30,required"`
	Password string `json:"password" binding:"min=8,max=30,required"`
	Nickname string `json:"nickname" binding:"min=8,max=30,required"`
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"len=6,required"`
}

type LoginQ struct {
	Username string `json:"username" binding:"min=5,max=30,required"`
	Password string `json:"password" binding:"gte=6,required"`
}

type LoginA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token"`
	UserId  int    `json:"UserId"`
}

type UploadFileA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Path    string `json:"path"`
}

type GetUserA struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
}
