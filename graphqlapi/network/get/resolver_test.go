/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
package network_get

import (
	"context"
	"testing"

	"github.com/creativesoftwarefdn/weaviate/models"
	"github.com/go-openapi/strfmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/source"
)

func TestNetworkGetInstanceQueryWithoutFilters(t *testing.T) {
	t.Parallel()
	resolver := &fakeNetworkResolver{}

	query := []byte(
		`{ Network { Get { weaviateA { Things { City { name } } } } } } `,
	)

	principal := &models.KeyTokenGetResponse{
		Token: strfmt.UUID("stand-in-for-token-uuid"),
		KeyGetResponse: models.KeyGetResponse{
			KeyID: strfmt.UUID("stand-in-for-key-id-uuid"),
		},
	}
	expectedSubQuery := `Get { Things { City { name } } }`
	expectedTarget := "weaviateA"
	expectedResultString := "placeholder for result from Local.Get"

	// in a real life scenario graphql will set the start and end
	// correctly. We just need to manually specify them in the test
	params := paramsFromQueryWithStartAndEnd(query, 18, 56, "weaviateA", resolver, principal)
	result, err := NetworkGetInstanceResolve(params)

	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if resolver.Called != true {
		t.Error("expected resolver.ProxyNetworkGetInstance to have been called, but it was never called")
	}

	if actual := resolver.CalledWith.SubQuery; string(actual) != expectedSubQuery {
		t.Errorf("expected subquery to be %s, but got %s", expectedSubQuery, actual)
	}

	if actual := resolver.CalledWith.TargetInstance; string(actual) != expectedTarget {
		t.Errorf("expected targetInstance to be %#v, but got %#v", expectedTarget, actual)
	}

	if actual := resolver.CalledWith.Principal; actual != principal {
		t.Errorf("expected principal to be %#v, but got %#v", principal, actual)
	}

	if _, ok := result.(string); !ok {
		t.Errorf("expected result to be a string, but was %t", result)
	}

	if resultString, ok := result.(string); !ok || resultString != expectedResultString {
		t.Errorf("expected result to be %s, but was %s", expectedResultString, resultString)
	}

}

func paramsFromQueryWithStartAndEnd(query []byte, start int, end int,
	instanceName string, resolver Resolver, principal *models.KeyTokenGetResponse) graphql.ResolveParams {
	return graphql.ResolveParams{
		Source: FiltersAndResolver{
			Resolver: resolver,
		},
		Info: graphql.ResolveInfo{
			FieldName: instanceName,
			FieldASTs: []*ast.Field{
				&ast.Field{
					Loc: &ast.Location{
						Start: start,
						Source: &source.Source{
							Body: []byte(query)},
						End: end,
					},
				},
			},
		},
		Context: context.WithValue(context.Background(), "principal", principal),
	}
}

type fakeNetworkResolver struct {
	Called     bool
	CalledWith ProxyGetInstanceParams
}

func (r *fakeNetworkResolver) ProxyGetInstance(info ProxyGetInstanceParams) (*models.GraphQLResponse, error) {
	r.Called = true
	r.CalledWith = info
	return &models.GraphQLResponse{
		Data: map[string]models.JSONObject{
			"Local": map[string]interface{}{
				"Get": "placeholder for result from Local.Get"},
		},
	}, nil
}