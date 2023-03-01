package productsdto

type ProductResponse struct {
	ID        			int    `json:"id"` 
	Name      			string `json:"name" form:"name" validate:"required"`   
	Price     			int    `json:"price" form:"price" validate:"required"`
	Description  string `json:"description" form:"description" validate:"required"`
	Stock  						int    `json:"stock" form:"stock" validate:"required"`
	Photo  						string `json:"photo" form:"photo" validate:"required"`
}
