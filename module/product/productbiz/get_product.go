package productbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

type GetProductStore interface {
	FindProductWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type getProductBusiness struct {
	store GetProductStore
}

func NewGetProductBusiness(store GetProductStore) *getProductBusiness {
	return &getProductBusiness{store: store}
}

func (business *getProductBusiness) GetProduct(context context.Context, id int) (*productmodel.Product, error) {
	result, err := business.store.FindProductWithCondition(context, map[string]interface{}{"id": id}, "Ratings", "Ratings.User")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(productmodel.EntityName, err)
	}
	return result, err
}
