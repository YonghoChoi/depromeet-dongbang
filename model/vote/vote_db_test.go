package vote

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"testing"
	"time"
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

	n := New(
		"",
		"test vote title",
		"test content",
		[]Option{Duplicate, Anonymous},
		time.Now(),
	)
	if err := Insert(n); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("insert vote. vote : %v\n", n)
}

func TestUpdate(t *testing.T) {
	setup()
	votes, err := GetVoteAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if votes == nil || len(votes) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. vote : %v\n", votes[0])

	votes[0].Content = "modified content"
	if err := Update(votes[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("updated notce. vote : %v\n", votes[0])
}

func TestDelete(t *testing.T) {
	setup()
	votes, err := GetVoteAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if votes == nil || len(votes) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. vote : %v\n", votes[0])
	if err := Delete(votes[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("deleted notce. vote : %v\n", votes[0])
}
