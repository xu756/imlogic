package middleware

import (
	"context"
	"errors"
	"imlogic/internal/xerr"
	"regexp"
	"strconv"
)

func ClientErrorHandler(ctx context.Context, err error) error {
	re := regexp.MustCompile(`code:(\d+),msg:(.+)`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) == 3 {
		code, _ := strconv.Atoi(match[1])
		return xerr.ErrMsg(int32(code))
	}
	return err
}

func ServerErrorHandler(ctx context.Context, err error) error {
	return errors.Unwrap(err)

}
