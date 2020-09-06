package response

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseFromEmptyString(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := middleware.ResponseJsonBody{}.Handle(request, func(request inter.Request) inter.Response {
		return outcome.NewResponse(outcome.Options{})
	})

	assert.Equal(t, "", response.Content())
}

func TestJsonResponseFromJsonString(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := middleware.ResponseJsonBody{}.Handle(request, func(request inter.Request) inter.Response {
		return outcome.Json("{\"height\": 12}")
	})

	assert.Equal(t, "{\"height\": 12}", response.Content())
}
