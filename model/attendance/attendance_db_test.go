package attendance

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"github.com/google/uuid"
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

	n := New(
		uuid.New().String(),
		time.Now(),
	)
	if err := Insert(n); err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("insert attendance. attendance : %v\n", n)
}

func TestUpdate(t *testing.T) {
	setup()
	attendances, err := GetAttendanceAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if attendances == nil || len(attendances) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. attendance : %v\n", attendances[0])

	attendances[0].Token = "modified token"
	if err := Update(attendances[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("updated notce. attendance : %v\n", attendances[0])
}

func TestDelete(t *testing.T) {
	setup()
	attendances, err := GetAttendanceAll()
	if err != nil {
		t.Fatal(err.Error())
	}

	if attendances == nil || len(attendances) == 0 {
		t.Fatal("not exist")
	}

	t.Logf("selected notce. attendance : %v\n", attendances[0])
	if err := Delete(attendances[0]); err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("deleted notce. attendance : %v\n", attendances[0])
}
