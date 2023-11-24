package msg

import "testing"

func TestSendEmail(t *testing.T) {

	err := SendEmail("756334744@qq.com", "测试邮件")
	if err != nil {
		t.Fatal(err)
	}
}
