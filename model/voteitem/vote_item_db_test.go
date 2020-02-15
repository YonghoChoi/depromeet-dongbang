package voteitem

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"testing"
)

func setup() {
	configPath := flag.String("config", "../../cmd/api/conf/config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)
}

func TestInsert(t *testing.T) {
	setup()
	n := New(
		"test",
		"test",
	)
	if err := Insert(n); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("insert vote. vote : %v\n", n)
}

func TestUpdate(t *testing.T) {
	setup()
	votes, err := GetVoteItemAll()
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
	votes, err := GetVoteItemAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if votes == nil || len(votes) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. vote : %v\n", votes[0])
	if err := Delete(votes[0].Id); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("deleted notce. vote : %v\n", votes[0])
}
