package gin

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type NewAlbumWrap struct {
	Data album `json:"data"`
}

func TestJsonBind(t *testing.T) {
	ctx := gin.Context{}
	case1 := NewAlbumWrap{
		Data: album{
			ID:     "5",
			Title:  "title5",
			Artist: "artist5",
			Price:  4.5,
		},
	}
	bytes, _ := json.Marshal(case1)
	postreq, err := http.NewRequest(http.MethodPost, "http://localhost:8080", strings.NewReader(string(bytes)))
	if err != nil {
		t.Errorf("new request failed, error: %v", err)
		panic(err)
	}
	ctx.Request = postreq

	var req NewAlbumWrap
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		t.Errorf("bind json failed, error: %v", err)
		panic(err)
	}

	t.Logf("request: %v", req)
}
