package product

import (
	"context"

	productv1 "github.com/GosMachine/protos/gen/go/product"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) GetCategory(ctx context.Context, name string) (*productv1.GetGategoryResponse, error) {
	resp, err := c.api.GetCategory(ctx, &productv1.GetCategoryRequest{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetCategoryNames(ctx context.Context) ([]string, error) {
	resp, err := c.api.GetCategoryNames(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return resp.Names, nil
}
