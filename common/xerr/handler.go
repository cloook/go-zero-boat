package xerr

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

func Handler(ctx context.Context, err error) (int, interface{}) {
	logx.WithContext(ctx).Errorf("[+++|API-ERR|+++]: %+v ,type:%+v", err, errors.Cause(err))
	switch e := errors.Cause(err).(type) {
	case *CodeError:
		return http.StatusOK, e.Data()
	default:
		if gstatus, ok := status.FromError(err); ok { // grpc err
			grpcCode := uint32(gstatus.Code())
			if IsCodeErr(grpcCode) { // filter custom err
				return http.StatusOK, NewRpcError(grpcCode, gstatus.Message())
			}
			return http.StatusOK, NewRpcError(SERVER_COMMON_ERROR, MapErrMsg(SERVER_COMMON_ERROR))
		}
		return http.StatusInternalServerError, nil
	}
}
