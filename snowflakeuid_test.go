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
		y, err = generator.GenerateId1()
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

	for i := 0; i < 100; i++ {
		id, err := generator.GenerateId2()
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("generate id: %v", id)
	}

	for i := 0; i < 100; i++ {
		id, y, err := generator.GenerateId3()
		if err != nil {
			t.Error(err)
			continue
		}
		t.Logf("generate id: %v, %s", id, y)
	}
}
