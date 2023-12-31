/*
 * ToDo API
 *
 * A simple ToDo API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// ToDoApiService is a service that implents the logic for the ToDoApiServicer
// This service should implement the business logic for every endpoint for the ToDoApi API. 
// Include any external packages or services that will be required by this service.
type ToDoApiService struct {
}

// NewToDoApiService creates a default api service
func NewToDoApiService() ToDoApiServicer {
	return &ToDoApiService{}
}

// TodosGet - Get all ToDos
func (s *ToDoApiService) TodosGet() (interface{}, error) {
	// TODO - update TodosGet with the required logic for this service method.
	// Add api_to_do_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'TodosGet' not implemented")
}

// TodosIdDelete - Delete a ToDo by ID
func (s *ToDoApiService) TodosIdDelete(id int32) (interface{}, error) {
	// TODO - update TodosIdDelete with the required logic for this service method.
	// Add api_to_do_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'TodosIdDelete' not implemented")
}

// TodosIdGet - Get a ToDo by ID
func (s *ToDoApiService) TodosIdGet(id int32) (interface{}, error) {
	// TODO - update TodosIdGet with the required logic for this service method.
	// Add api_to_do_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'TodosIdGet' not implemented")
}

// TodosIdPut - Update a ToDo by ID
func (s *ToDoApiService) TodosIdPut(id int32, todo Todo) (interface{}, error) {
	// TODO - update TodosIdPut with the required logic for this service method.
	// Add api_to_do_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'TodosIdPut' not implemented")
}

// TodosPost - Create a ToDo
func (s *ToDoApiService) TodosPost(todo Todo) (interface{}, error) {
	// TODO - update TodosPost with the required logic for this service method.
	// Add api_to_do_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'TodosPost' not implemented")
}
