package response

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/transport/http"
	hs "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
	stdhttp "net/http"
)

type httpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func EncoderResponse() http.EncodeResponseFunc {
	return func(w stdhttp.ResponseWriter, request *stdhttp.Request, i interface{}) error {
		if i == nil {
			return nil
		}
		resp := &httpResponse{
			Code:    stdhttp.StatusOK,
			Message: "操作成功",
			Data:    i,
		}
		data, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

func EncoderError() http.EncodeErrorFunc {
	return func(w stdhttp.ResponseWriter, r *stdhttp.Request, err error) {
		if err == nil {
			return
		}
		se := &httpResponse{}
		gs, ok := status.FromError(err)
		if !ok {
			se = &httpResponse{Code: stdhttp.StatusInternalServerError}
		}
		se = &httpResponse{
			Code:    hs.FromGRPCCode(gs.Code()),
			Message: gs.Message(),
			Data:    nil,
		}
		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(stdhttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/"+codec.Name())
		w.WriteHeader(se.Code)
		_, _ = w.Write(body)
	}
}
