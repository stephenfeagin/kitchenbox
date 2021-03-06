package ldjson

import (
	"testing"
	"time"
)

func TestParseDatePublishedJSON(t *testing.T) {
	want := time.Date(2019, 4, 2, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name, input string
		want        time.Time
	}{

		{"Full Datetime", "2019-04-02T00:00:00.000Z", want},
		{"Date Only", "2019-04-02", want},
		{"Invalid Input", "xxx", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			datePublished, _ := unmarshalDatePublished(test.input)
			if datePublished != test.want {
				t.Fatal("incorrectly unmarshalled DatePublished")
			}
		})
	}
}
