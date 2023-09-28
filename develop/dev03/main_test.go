package main

import (
	"testing"
)

func TestCheckKeysReverse(t *testing.T) {
	lines := []*Line{
		{StrValue: "ner"},
		{StrValue: "atc"},
		{StrValue: "boc"},
	}

	keys := []string{"-r"}

	expected := []*Line{
		{StrValue: "ner"},
		{StrValue: "boc"},
		{StrValue: "atc"},
	}

	result := checkKeys(lines, keys)

	for i := range expected {
		if expected[i].StrValue != result[i].StrValue {
			t.Errorf("Ожидаемый результат: %+v, полученный результат: %+v", expected[i].StrValue, result[i].StrValue)
		}
	}
}
func TestCheckKeysNumSuffix(t *testing.T) {
	lines := []*Line{
		{StrValue: "2K"},
		{StrValue: "1G"},
		{StrValue: "3M"},
		{StrValue: "5K"},
		{StrValue: "10M"},
	}

	keys := []string{"-h"}

	expected := []*Line{
		{NumValue: 2000},
		{NumValue: 5000},
		{NumValue: 3000000},
		{NumValue: 10000000},
		{NumValue: 1000000000},
	}

	result := checkKeys(lines, keys)

	for i := range expected {
		if expected[i].NumValue != result[i].NumValue {
			t.Errorf("Ожидаемый результат: %+v, полученный результат: %+v", expected[i].NumValue, result[i].NumValue)
		}
	}
}
func TestCheckKeysMonth(t *testing.T) {
	lines := []*Line{
		{StrValue: "September"},
		{StrValue: "August"},
		{StrValue: "January"},
		{StrValue: "October"},
		{StrValue: "December"},
	}

	keys := []string{"-M"}

	expected := []*Line{
		{StrValue: "January"},
		{StrValue: "August"},
		{StrValue: "September"},
		{StrValue: "October"},
		{StrValue: "December"},
	}

	result := checkKeys(lines, keys)

	for i := range expected {
		if expected[i].StrValue != result[i].StrValue {
			t.Errorf("Ожидаемый результат: %+v, полученный результат: %+v", expected[i].StrValue, result[i].StrValue)
		}
	}
}

func TestCheckKeysNum(t *testing.T) {
	lines := []*Line{
		{NumValue: 10011},
		{NumValue: 1245},
		{NumValue: 5423},
		{NumValue: 5},
		{NumValue: 7652},
	}

	keys := []string{"-n"}

	expected := []*Line{
		{NumValue: 5},
		{NumValue: 1245},
		{NumValue: 5423},
		{NumValue: 7652},
		{NumValue: 10011},
	}

	result := checkKeys(lines, keys)

	for i := range expected {
		if expected[i].NumValue != result[i].NumValue {
			t.Errorf("Ожидаемый результат: %+v, полученный результат: %+v", expected[i].NumValue, result[i].NumValue)
		}
	}
}
