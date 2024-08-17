package middleware

import (
	"context"
	"errors"
	"imlogic/internal/xerr"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote"
)

func ClientErrorHandler(ctx context.Context, err error) error {
	var e *remote.TransError
	if errors.As(err, &e) {
		return xerr.CodeError{
			Code: e.TypeID(),
			Msg:  e.Error(),
		}

	}
	return err
}

func ServerErrorHandler(ctx context.Context, err error) error {
	if errors.Is(err, kerrors.ErrBiz) {
		err = errors.Unwrap(err)
	}
	var xErr xerr.CodeError
	if errors.As(err, &xErr) {
		return remote.NewTransError(xErr.GetCode(), xErr.RpcErr())
	}
	return err

}
