package result

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/xu756/imlogic/internal/xerr"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HttpResult http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errCode := xerr.SystemError
		errMsg := "系统错误,请联系管理员"

		re := regexp.MustCompile(`ErrCode:(\d+).*，ErrMsg:(.*)`)
		// 查找匹配项
		matches := re.FindStringSubmatch(err.Error())

		// 提取 ErrCode 和 ErrMsg
		if len(matches) >= 3 {
			Code, _ := strconv.ParseUint(matches[1], 10, 32)
			errCode = uint32(Code)
			errMsg = matches[2]
		} else {
			fmt.Println("未找到 ErrCode 和 ErrMsg")
		}
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
}
