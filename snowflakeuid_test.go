package uidgo_test

import (
	"testing"
	"uidgo"
)

func TestSnowflakeSeqGenerator_GenerateId(t *testing.T) {
	var dataCenterId, workId int64 = 1, 1
	generator, err := uidgo.NewSnowflakeSeqGenerator(dataCenterId, workId)
	if err != nil {
		t.Error(err)
		return
	}
	var x, y string
	for i := 0; i < 100; i++ {
		y, err = generator.GenerateId()
		if err != nil {
			t.Error(err)
			continue
		}
		if x == y {
			t.Errorf("x(%s) & y(%s) are the same", x, y)
		}
		t.Logf("generate id: %s", y)
		x = y
	}

	dataCenterId, workId = 1, 1
	generator, err = uidgo.NewSnowflakeSeqGenerator(dataCenterId, workId)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 100; i++ {
		y, err = generator.GenerateId()
		if err != nil {
			t.Error(err)
			continue
		}
		if x == y {
			t.Errorf("x(%s) & y(%s) are the same", x, y)
		}
		t.Logf("generate id: %s", y)
		x = y
	}
}
