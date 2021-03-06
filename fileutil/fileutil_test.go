// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileutil

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestIsDirWriteable(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unexpected ioutil.TempDir error: %v", err)
	}
	defer os.RemoveAll(tmpdir)
	if err = IsDirWriteable(tmpdir); err != nil {
		t.Fatalf("unexpected IsDirWriteable error: %v", err)
	}
	if err = os.Chmod(tmpdir, 0444); err != nil {
		t.Fatalf("unexpected os.Chmod error: %v", err)
	}
	me, err := user.Current()
	if err != nil {
		// err can be non-nil when cross compiled
		// http://stackoverflow.com/questions/20609415/cross-compiling-user-current-not-implemented-on-linux-amd64
		t.Skipf("failed to get current user: %v", err)
	}
	if me.Name == "root" || runtime.GOOS == "windows" {
		// ideally we should check CAP_DAC_OVERRIDE.
		// but it does not matter for tests.
		// Chmod is not supported under windows.
		t.Skipf("running as a superuser or in windows")
	}
	if err := IsDirWriteable(tmpdir); err == nil {
		t.Fatalf("expected IsDirWriteable to error")
	}
}

func TestCreateDirAll(t *testing.T) {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "foo")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	tmpdir2 := filepath.Join(tmpdir, "testdir")
	if err = CreateDirAll(tmpdir2); err != nil {
		t.Fatal(err)
	}

	if err = ioutil.WriteFile(filepath.Join(tmpdir2, "text.txt"), []byte("test text"), PrivateFileMode); err != nil {
		t.Fatal(err)
	}

	if err = CreateDirAll(tmpdir2); err == nil || !strings.Contains(err.Error(), "to be empty, got") {
		t.Fatalf("unexpected error %v", err)
	}
}

func TestExist(t *testing.T) {
	fdir := filepath.Join(os.TempDir(), fmt.Sprint(time.Now().UnixNano()+rand.Int63n(1000)))
	os.RemoveAll(fdir)
	if err := os.Mkdir(fdir, 0666); err != nil {
		t.Skip(err)
	}
	defer os.RemoveAll(fdir)
	if !Exist(fdir) {
		t.Fatalf("expected Exist true, got %v", Exist(fdir))
	}

	f, err := ioutil.TempFile(os.TempDir(), "fileutil")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	if g := Exist(f.Name()); !g {
		t.Errorf("exist = %v, want true", g)
	}

	os.Remove(f.Name())
	if g := Exist(f.Name()); g {
		t.Errorf("exist = %v, want false", g)
	}
}
