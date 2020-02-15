package voter

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

	n := New(
		"vote id",
		u,
	)
	if err := Insert(n); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("insert vote. vote : %v\n", n)
}

func TestUpdate(t *testing.T) {
	setup()
	votes, err := GetVoterAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if votes == nil || len(votes) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected voter. vote : %v\n", votes[0])

	votes[0].VoteItemId = "test"
	if err := Update(votes[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("updated voter. vote : %v\n", votes[0])
}

func TestDelete(t *testing.T) {
	setup()
	votes, err := GetVoterAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if votes == nil || len(votes) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected voter. vote : %v\n", votes[0])
	if err := Delete(votes[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("deleted voter. vote : %v\n", votes[0])
}
