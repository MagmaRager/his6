package service

import (
	"testing"
)

func TestLogin(t *testing.T) {
	var datas = []struct {
		code string
		password string
		id int
	}{
		{"1003", "1", 100003},
		{"1002", "123456", 100002},
		{"1001", "123456", 100001},
	}

	var login LoginService
	for _, d := range datas {
		actual, err := login.Login(d.code, d.password)
		if err != nil {
			t.Errorf("Error!" + err.Error())
			continue
		}
		if actual.Id != d.id {
			t.Errorf("Expected the sum of %v to be %d but instead got %d!", d, d.id, actual.Id)
		}
	}
}


func BenchmarkLogin(b *testing.B) {
	var login LoginService
	for n := 0; n < b.N; n++ {
		login.Login("1001", "123456")
	}
}

func BenchmarkMenu(b *testing.B) {
	var login LoginService
	for n := 0; n < b.N; n++ {
		login.QueryMenus()
	}
}

func BenchmarkCacheMenu(b *testing.B) {
	var login LoginService
	for n := 0; n < b.N; n++ {
		login.QueryCacheMenus()
	}
}
