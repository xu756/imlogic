package middleware

import (
	"context"
	"errors"
	"imlogic/internal/xerr"
	"log"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"google.golang.org/grpc/status"
)

func ClientErrorHandler(ctx context.Context, err error) error {
	if s, ok := status.FromError(err); ok {
		log.Print("【 client error 】", s.Message())
	}
	return err
}

func ServerErrorHandler(ctx context.Context, err error) error {
	if errors.Is(err, kerrors.ErrBiz) {
		err = errors.Unwrap(err)
	}
	if errCode, ok := xerr.GetErrorCode(err); ok {
		return status.Errorf(errCode, err.Error())
	}

	return err

}
