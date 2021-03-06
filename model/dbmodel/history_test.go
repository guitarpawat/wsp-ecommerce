package dbmodel

import (
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"
)

func TestMakeSalesHistory_NoMeat(t *testing.T) {
	user := User{ID: "1"}
	expectedTime := time.Now()
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	sh, _ := MakeSalesHistory(expectedTime, user, map[Meat]int{}, expectedPrice, expectedTrackingNumber)

	if expectedTime != sh.Time {
		t.Errorf("expected time: %s, but get: %s", expectedTime, sh.Time)
	}

	for _ = range sh.Meats {
		t.Fatalf("not expected any meats to be here")
	}

	if expectedPrice != sh.Price {
		t.Errorf("expected price: %f, but get: %f", expectedPrice, sh.Price)
	}

	if expectedTrackingNumber != sh.TrackingNumber {
		t.Errorf("expected time: %s, but get: %s", expectedTrackingNumber, sh.TrackingNumber)
	}
}

func TestMakeSalesHistory_NoUser(t *testing.T) {
	user := User{}
	expectedTime := time.Now()
	mockMeat1 := Meat{ID: "1"}
	mockMeat2 := Meat{ID: "2"}
	expectedMeat := map[Meat]int{mockMeat1: 3, mockMeat2: 1}
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	_, err := MakeSalesHistory(expectedTime, user, expectedMeat, expectedPrice, expectedTrackingNumber)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestMakeSalesHistory_SomeMeat(t *testing.T) {
	user := User{ID: bson.ObjectIdHex("cbccd007f4f2ee7cdb6314c2")}
	mockMeat1 := Meat{ID: bson.ObjectIdHex("97a12c5313341fb31752da22")}
	mockMeat2 := Meat{ID: bson.ObjectIdHex("d6441111b9efd95d06e76774")}
	expectedTime := time.Now()
	expectedMeat := []Meats{{mockMeat1.ID, 1}, {mockMeat2.ID, 3}}
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	sh, _ := MakeSalesHistory(expectedTime, user, map[Meat]int{mockMeat1: expectedMeat[0].Quatity, mockMeat2: expectedMeat[1].Quatity}, expectedPrice, expectedTrackingNumber)

	if expectedTime != sh.Time {
		t.Errorf("expected time: %s, but get: %s", expectedTime, sh.Time)
	}

	if len(sh.Meats) != 2 {
		t.Fatalf("expected to have meat: 2")
	}

	val1 := sh.GetMeat(expectedMeat[0].Meat)
	if val1 != expectedMeat[0] {
		t.Errorf("expected quantity of meat id %s: %d, but get id %s: %d", mockMeat1.ID, expectedMeat[0].Quatity, val1.Meat, val1.Quatity)
	}

	val2 := sh.GetMeat(expectedMeat[1].Meat)
	if val2 != expectedMeat[1] {
		t.Errorf("expected quantity of meat id %s: %d, but getid %s: %d", mockMeat2.ID, expectedMeat[1].Quatity, val2.Meat, val2.Quatity)
	}

	if expectedPrice != sh.Price {
		t.Errorf("expected price: %f, but get: %f", expectedPrice, sh.Price)
	}

	if expectedTrackingNumber != sh.TrackingNumber {
		t.Errorf("expected time: %s, but get: %s", expectedTrackingNumber, sh.TrackingNumber)
	}
}

func TestMakeSalesHistory_ErrorNoID(t *testing.T) {
	user := User{ID: "1"}
	mockMeat := Meat{}
	expectedTime := time.Now()
	expectedMeat := map[Meat]int{mockMeat: 3}
	expectedPrice := 10.25
	expectedTrackingNumber := "EA123456789TH"

	_, err := MakeSalesHistory(expectedTime, user, expectedMeat, expectedPrice, expectedTrackingNumber)

	if err == nil {
		t.Errorf("expected error")
	}
}
