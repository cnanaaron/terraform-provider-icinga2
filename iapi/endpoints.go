package iapi

import (
	"encoding/json"
	"fmt"
)

// https://icinga.com/docs/icinga-2/2.10/doc/09-object-types/#endpoint

// GetEndpoint ...
func (server *Server) GetEndpoint(endpointname string) ([]EndpointStruct, error) {

	var endpoints []EndpointStruct

	results, err := server.NewAPIRequest("GET", "/objects/endpoints/"+endpointname, nil)
	if err != nil {
		return nil, err
	}

	// Contents of the results is an interface object. Need to convert it to json first.
	jsonStr, marshalErr := json.Marshal(results.Results)
	if marshalErr != nil {
		return nil, marshalErr
	}

	// then the JSON can be pushed into the appropriate struct.
	// Note : Results is a slice so much push into a slice.

	if unmarshalErr := json.Unmarshal(jsonStr, &endpoints); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return endpoints, err

}

// CreateEndpoint ...
func (server *Server) CreateEndpoint(endpointname string, host string, port string, log_duration int) ([]EndpointStruct, error) {

	var newAttrs EndpointAttrs
	newAttrs.Host = host
	newAttrs.Port = port
	newAttrs.LogDuration = log_duration

	var newEndpoint EndpointStruct
	newEndpoint.Name = endpointname
	newEndpoint.Type = "Endpoint"
	newEndpoint.Attrs = newAttrs

	// Create JSON from completed struct
	payloadJSON, marshalErr := json.Marshal(newEndpoint)
	if marshalErr != nil {
		return nil, marshalErr
	}

	//fmt.Printf("<payload> %s\n", payloadJSON)

	// Make the API request to create the endpoints.
	results, err := server.NewAPIRequest("PUT", "/objects/endpoints/"+endpointname, []byte(payloadJSON))
	if err != nil {
		return nil, err
	}

	if results.Code == 200 {
		endpoints, err := server.GetEndpoint(endpointname)
		return endpoints, err
	}

	return nil, fmt.Errorf("%s", results.ErrorString)

}

// DeleteEndpoint ...
func (server *Server) DeleteEndpoint(endpointname string) error {
	results, err := server.NewAPIRequest("DELETE", "/objects/endpoints/"+endpointname+"?cascade=1", nil)
	if err != nil {
		return err
	}

	if results.Code == 200 {
		return nil
	}

	return fmt.Errorf("%s", results.ErrorString)
}
