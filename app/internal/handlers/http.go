package handlers

import (
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/domain"
	"github.com/JuanGQCadavid/coding-interview-backend-level-3/app/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	service ports.Service
}

func NewHttpHandler(service ports.Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (hdl *HttpHandler) SetRouter(router *gin.Engine) {
	router.GET("/ping", hdl.Ping)
	router.GET("/items", hdl.GetItems)
	router.GET("/items/:id", hdl.GetItem)
	router.POST("/items", hdl.CreateItem)
	router.PUT("/items/:id", hdl.UpdateItem)
	router.DELETE("/items/:id", hdl.DeleteItem)

}

// GET ping

// method: 'GET',
//
//	url: '/ping' - 200
func (hdl *HttpHandler) Ping(context *gin.Context) {}

// method: 'GET',
// url: '/items' - 200
// [] Empty if not items
// [
//
//	{
//		id: expect.any(Number),
//		name: 'Item 1',
//		price: 10
//	}
//
// ]
func (hdl *HttpHandler) GetItems(context *gin.Context) {}

//	 method: 'GET',
//	   url: `/items/${response.result!.id}`
//	200
//
// 404 - Not found
//
//	{
//		id: expect.any(Number),
//		name: 'Item 1',
//		price: 10
//	}
func (hdl *HttpHandler) GetItem(context *gin.Context) {}

// Validations: 400 status code : POST | PUT
// errors: [
// 	{
// 		field: 'price',
// 		message: 'Field "price" cannot be negative'
// 	}
// ]

// errors: [
//
//	{
//		field: 'price',
//		message: 'Field "price" is required'
//	}
//
// ]
func (hdl *HttpHandler) validateRequest(item *domain.Item) error {
	return nil
}

// method: 'POST',
// url: '/items',
//
//	payload: {
//		name: 'Item 1',
//		price: 10
//	}
func (hdl *HttpHandler) CreateItem(context *gin.Context) {
	// validateRequest(item *domain.Item)
}

// {
// 	method: 'PUT',
// 	url: `/items/${createdItem!.id}`,
// 	payload: {
// 		name: 'Item 1 updated',
// 		price: 20
// 	}
// }

// 200 -> Propose 204
func (hdl *HttpHandler) UpdateItem(context *gin.Context) {
	// validateRequest(item *domain.Item)
}

//	{
//		method: 'DELETE',
//		url: `/items/${createdItem!.id}`
//	}
//
// 204
func (hdl *HttpHandler) DeleteItem(context *gin.Context) {}
