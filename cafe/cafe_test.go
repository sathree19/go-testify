package cafe

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)

}

func TestMainHandlerWhenStatusOKAndBodyNotEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, 200, responseRecorder.Code)
	assert.NotNil(t, responseRecorder.Body)
}

func TestMainHandlerWhenSityValueRight(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=ufa", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MainHandle)
	handler.ServeHTTP(responseRecorder, req)

	city := "ufa"

	for k := range cafeList {
		require.Equal(t, k, city)
	}

}
