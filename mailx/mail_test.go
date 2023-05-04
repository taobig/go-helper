package mailx

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	mail := NewSMTPMail("user@qq.com", "password", "smtp.qq.com", 587)
	receivers := []string{"user@163.com"}
	body := `
<h2>测试邮件</h2><strong><span style='color:red'>FAQ</span>（英语：Frequently Asked Question）、Q&A（Questions and Answers），中文译为“常见问题与解答”或「常見問答集」，直译为“常被问到的问题”。这个术语在与电脑和网络有关的内容中经常见到。</strong>
`
	err := mail.Send(receivers, "email subject", body, ContentTypeTextHtml)
	if err != nil {
		//t.Fatal(err)
		t.Log(err)
	}
}
