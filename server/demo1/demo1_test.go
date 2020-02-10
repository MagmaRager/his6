package demo1

import (
	"fmt"
	"testing"

	"github.com/kataras/iris/httptest"
)

func TestDemo1(t *testing.T) {
	app := base.App
	base.RegisterGetHandler("/login", loginHandler)
	base.RegisterGetHandler("/menus", queryMenuHandle)
	e := httptest.New(t, app)

	var datas = []struct {
		code     string
		password string
		id       int
	}{
		{"1003", "1", 100003},
		{"1002", "123456", 100002},
		{"1001", "123456", 100001},
	}

	for _, d := range datas {
		url := "?code=" + d.code + "&password=" + d.password

		resp := e.GET("/login").WithURL(url).Expect()
		resp.Status(httptest.StatusOK)
		fmt.Println(resp.Body())
	}

	e.GET("/menus").Expect().Status(httptest.StatusOK)
}
