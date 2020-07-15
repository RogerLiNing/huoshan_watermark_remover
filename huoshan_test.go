package huoshan

import (
	"testing"
)

func TestGetAvailableHuoshanLink(t *testing.T) {
	t.Log("测试正常的火山视频链接")
	url := "https://share.huoshan.com/hotsoon/s/OHBF20a9Y98/"
	t.Log(url)
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
		t.Log("返回视频链接，测试通过")
	}

	t.Log(videoLink)
}

func TestGetUnAvailableHuoshanLink(t *testing.T) {
	t.Log("测试无效的火山视频链接")
	url := "https://share.huoshan.com/hotsoon/s/88888888888/"
	t.Log(url)
	videoLink, err := WatermarkRemover(url)

	if err != nil {
		t.Fail()
		t.Log("返回值为空，测试通过")
	}

	if len(videoLink) == 0 {
		t.Log("测试通过")
	}

}
