package postgres

import (
	ConfigMain "github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/postgres_gorm"
	"testing"
)

func TestFillMassTable(t *testing.T) {
	ConfigMain.LoadEnv()
	postgres_gorm.Connect()
	defer postgres_gorm.CloseConnection()

	Otvet, err := FillMapTable()
	if err != nil {
		t.Error("TestFillMassTable() error: ", err)
	}

	if len(Otvet) == 0 {
		t.Error("TestFillMassTable() error: len =0")
	}
}
