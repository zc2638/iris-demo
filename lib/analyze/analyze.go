package analyze

import (
	"encoding/json"
	"github.com/kataras/iris/core/errors"
	"sop/lib/curl"
)

const domain = "http://106.75.241.81"

// 推送数据
func Create(sop, models string) error {

	const url  = domain + "/match/update"

	h := curl.HttpReq{
		Url: url,
		Params: map[string]interface{}{
			"sops": sop,
			"sop_models": models,
		},
	}

	b, err := h.Post()
	if err != nil {
		return err
	}

	var res struct{
		Detail string `json:"details"`
		Err int `json:"err"`
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	if res.Err != 0 {
		return errors.New(res.Detail)
	}
	return nil
}

// 数据匹配
func Match(jsonStr string) (result MatchResult, err error) {

	const url = domain + "/match"

	h := curl.HttpReq{
		Url: url,
		Params: map[string]interface{}{
			"aps": jsonStr,
		},
	}

	b, err := h.Post()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &result)
	return
}

type MatchResult []struct {
	Aps uint `json:"aps"`
	Sop uint `json:"sop"`
}