package config

import (
	"math/rand"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	data := Get()

	if data.FileName != "config" {
		t.Errorf("dta not valid expected : %v got : %v", "config", data.FileName)
	}

	randomString := []string{"fjhgkj", "ndnv", "wii", "hee", "wrw", "qwq", "xv", "amnd", "wiru", "wrtfd", "wye", "vbbb", "drr"}

	sourceRand := rand.NewSource(time.Now().UnixNano())
	random := rand.New(sourceRand)

	index := random.Intn(len(randomString))
	keyExpected := randomString[index]

	index = random.Intn(len(randomString))
	secretExpected := randomString[index]

	data.Config.API.Key = keyExpected
	data.Config.API.Secret = secretExpected

	err := Save(data)
	if err != nil {
		t.Error("error test saving data err : " + err.Error())
	}

	Refresh()

	data = Get()

	if data.Config.API.Key != keyExpected && data.Config.API.Secret != secretExpected {
		t.Errorf("error test refreshing data got : %#v", data)
	}

}
