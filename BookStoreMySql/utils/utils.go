package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	// ReadAll이 어떤 역할을 하는지 확인

	if err != nil {
		panic(err)
	} else {
		marshalErr := json.Unmarshal([]byte(body), x)

		if marshalErr != nil {
			return
		}
	}
}
