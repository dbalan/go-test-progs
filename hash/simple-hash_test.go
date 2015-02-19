package main

import (
	"testing"
)

type testcase struct {
	s string
	h string
}

var tests = []testcase{
	{"helloworld", "fc5e038d38a57032085441e7fe7010b0"},
	{"", "d41d8cd98f00b204e9800998ecf8427e"},
	{"djkdfgjhdfgjhdfjghdfg", "e46cf1b94c2c687a244a6aa52a75fffb"},
}

func TestMD5Hasing(t *testing.T) {
	for _, pair := range tests {
		v := md5Hash(pair.s)
		if v != pair.h {
			t.Error(
				"For ", pair.s,
				"got ", v,
				"for ", pair.h,
			)
		}

	}
}
