package wechat

import (
	"context"
	"encoding/json"
	local_http "go_third_auth/local-http"
	"go_third_auth/tools"
)

func (w Wechat) Do(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	var atErr WechatError
	respRaw, err := local_http.Get(ctx, baseUri+path,
		tools.MapMerge(tools.StructToMap(payload), tools.StructToMap(w.opts)))

	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(respRaw, &atErr)
	if atErr.ErrCode != 0 {
		return nil, atErr
	}
	return respRaw, nil
}
