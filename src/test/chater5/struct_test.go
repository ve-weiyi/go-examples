package chater5

import (
	"github.com/ve-weiyi/go-examplse/src/utils/jsonconv"
	"github.com/ve-weiyi/go-examplse/src/utils/response"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	apiError := response.SuccessResponse{}
	log.Printf("%+v", jsonconv.ObjectToJson(apiError))
	log.Printf("%+v", apiError.ToJson())
}
