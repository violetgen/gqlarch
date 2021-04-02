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

func (r *queryResolver) Sum(ctx context.Context, req model.SumRequest) (*model.SumResponse, error) {
	var request = model.SumRequest{A: req.A, B: req.B}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8081/add", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.SumResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &model.SumResponse{Result: response.Result}, nil
}

func (r *queryResolver) Subtract(ctx context.Context, req model.SubtractRequest) (*model.SubtractResponse, error) {
	var request = model.SubtractRequest{A: req.A, B: req.B}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:8082/subtract", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response model.SubtractResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &model.SubtractResponse{Result: response.Result}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
