// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structreflectexample_test

import (
	"encoding/json"
	"testing"

	"github.com/danil/structreflectexample"
)

type T struct {
	Foo string
	Bar int
	xyz bool
}

func TestReflect(t *testing.T) {
	v := T{Foo: "Hello, World!", Bar: 42, xyz: true}

	got, err := structreflectexample.Reflect(v)
	if err != nil {
		t.Fatal(err)
	}

	want := "structreflectexample_test.T{Foo:Hello, World!,Bar:42}"

	if got != want {
		t.Errorf("\nwant: %s\n got: %s", want, got)
	}
}

func BenchmarkReflect(b *testing.B) {
	b.ReportAllocs()

	v := T{Foo: "Hello, World!", Bar: 42, xyz: true}

	b.Run("reflect", func(b *testing.B) {
		_, _ = structreflectexample.Reflect(v)
	})
}

func BenchmarkMarshal(b *testing.B) {
	b.ReportAllocs()

	v := T{Foo: "Hello, World!", Bar: 42, xyz: true}

	b.Run("marshal", func(b *testing.B) {
		_, _ = json.Marshal(v)
	})
}
