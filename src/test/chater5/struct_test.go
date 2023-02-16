package chater5

import (
	"github.com/ve-weiyi/go-examplse/src/utils/apierr"
	"github.com/ve-weiyi/go-examplse/src/utils/jsonconv"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	apiError := apierr.ApiError{}
	log.Printf("%+v", jsonconv.ObjectToJson(apiError))
	log.Printf("%+v", apiError.ToJson())
}
