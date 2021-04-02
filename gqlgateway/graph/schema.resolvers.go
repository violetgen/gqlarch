package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/parikshitg/gqlarch/gqlgateway/graph/generated"
	"github.com/parikshitg/gqlarch/gqlgateway/graph/model"
)

func HttpWrapper(url string, req, res interface{}) error {
	requestBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	return nil
}

func (r *queryResolver) Sum(ctx context.Context, req model.SumRequest) (*model.SumResponse, error) {
	var request = &model.SumRequest{A: req.A, B: req.B}
	var response model.SumResponse
	err := HttpWrapper("http://localhost:8081/add", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (r *queryResolver) Subtract(ctx context.Context, req model.SubtractRequest) (*model.SubtractResponse, error) {
	var request = &model.SubtractRequest{A: req.A, B: req.B}
	var response model.SubtractResponse
	err := HttpWrapper("http://localhost:8082/subtract", request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
