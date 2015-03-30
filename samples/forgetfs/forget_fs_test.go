// Copyright 2015 Google Inc. All Rights Reserved.
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

package forgetfs_test

import (
	"testing"

	"github.com/jacobsa/fuse/samples"
	"github.com/jacobsa/fuse/samples/forgetfs"
	. "github.com/jacobsa/ogletest"
)

func TestForgetFS(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Boilerplate
////////////////////////////////////////////////////////////////////////

type ForgetFSTest struct {
	samples.SampleTest
	fs *forgetfs.ForgetFS
}

func init() { RegisterTestSuite(&ForgetFSTest{}) }

func (t *ForgetFSTest) SetUp(ti *TestInfo) {
	panic("TODO")
}

func (t *ForgetFSTest) TearDown() {
	panic("TODO: Unmount then call Check")
}

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *ForgetFSTest) Open_Foo() {
	AssertTrue(false, "TODO")
}

func (t *ForgetFSTest) Open_Dir() {
	AssertTrue(false, "TODO")
}

func (t *ForgetFSTest) Stat_Foo() {
	AssertTrue(false, "TODO")
}

func (t *ForgetFSTest) Stat_Dir() {
	AssertTrue(false, "TODO")
}

func (t *ForgetFSTest) CreateFile() {
	AssertTrue(false, "TODO")
}

func (t *ForgetFSTest) MkDir() {
	AssertTrue(false, "TODO")
}
