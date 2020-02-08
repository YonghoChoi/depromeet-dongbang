package notice

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"testing"
)

func setup() {
	configPath := flag.String("config", "../../cmd/api/conf/config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)
}

func TestInsert(t *testing.T) {
	setup()
	u := user.New("test user", "test token", "")
	if err := user.Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("insert user. user : %v\n", u)

	n := New(u, "test notice", "test content")
	if err := Insert(n); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("insert notice. notice : %v\n", n)
}

func TestUpdate(t *testing.T) {
	setup()
	notices, err := GetNoticeAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if notices == nil || len(notices) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. notice : %v\n", notices[0])

	notices[0].Content = "modified content"
	if err := Update(notices[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("updated notce. notice : %v\n", notices[0])
}

func TestDelete(t *testing.T) {
	setup()
	notices, err := GetNoticeAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if notices == nil || len(notices) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. notice : %v\n", notices[0])
	if err := Delete(notices[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("deleted notce. notice : %v\n", notices[0])
}
