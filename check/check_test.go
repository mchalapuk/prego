/*
   Copyright 2014 Maciej Chałapuk

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package check

import "testing"
import "fmt"
import "strings"

const (
  irrelevant = "irrelevant in this test"
)

func expectError(t *testing.T, test func () error, checkCallTemplate string,
values ...interface{}) {
  err := test()
  if err == nil {
    checkCallStr := fmt.Sprintf(checkCallTemplate, values...)
    t.Errorf("Expected %s to return error, got nil", checkCallStr)
  }
}

func expectNoError(t *testing.T, test func () error, checkCallTemplate string,
values ...interface{}) {
  err := test()
  if err != nil {
    checkCallStr := fmt.Sprintf(checkCallTemplate, values...)
    t.Errorf("Expected %s to return nil, got %#v", checkCallStr,  err)
  }
}

func expectMessage(t *testing.T, test func () error, expectedMessage string,
checkCallTemplate string, values ...interface{}) {
  err := test()
  actualMessage := err.Error()
  if strings.Index(actualMessage, expectedMessage) != 0 {
    checkCallStr := fmt.Sprintf(checkCallTemplate, values...)
    t.Errorf("Expected %s to return error with message %s, got %#v",
    checkCallStr, expectedMessage, actualMessage)
  }
}

func TestTrueReturnsNilWhenPassingTrue(t *testing.T) {
  expectNoError(t, func() error { return True(true, irrelevant) }, "check.True(true)")
}

func TestTrueReturnsErrWhenPassingFalse(t *testing.T) {
  expectError(t, func() error { return True(false, irrelevant) }, "check.True(false)")
}

func TestFlaseReturnsNilWhenPassingFalse(t *testing.T) {
  expectNoError(t,func() error { return False(false,irrelevant) },"check.False(false)")
}

func TestFalseReturnsErrWhenPassingTrue(t *testing.T) {
  expectError(t, func() error { return False(true, irrelevant) }, "check.False(true)")
}

type Test struct {
}

func TestNilReturnsNilWhenPassingNil(t *testing.T) {
  expectNoError(t, func() error { return Nil(nil, irrelevant) }, "check.Nil(nil)")
}

func TestNilReturnsErrWhenPassingNotNil(t *testing.T) {
  val := &Test{}
  expectError(t, func() error { return Nil(val, irrelevant) }, "check.Nil(%#v)", val)
}

func TestNotNilReturnsNilWhenPassingNotNil(t *testing.T) {
  val := &Test{}
  expectNoError(t, func() error { return NotNil(val, irrelevant) }, "check.NotNil(%#v)", val)
}

func TestNotNilReturnsErrWhenPassingNil(t *testing.T) {
  expectError(t, func() error { return NotNil(nil, irrelevant) }, "check.NotNil(nil)" )
}

func TestInRangeEpsilonReturnsNilWhenPassingValueInRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {0, 0, 0 ,.1},
    []float64 {0, 0, 0 ,.0000001},
    []float64 {0, 0, 1 ,.1},
    []float64 {0, -1, 0 ,.1},
    []float64 {0, -1, 1 ,.1},
    []float64 {0, -.000001, .000001, .0000001},
    []float64 {1000, 0, 1000, .1},
    []float64 {-1000, -2000, 0, .1},
    []float64 {.1, 0, 0, .2},
  }
  template := "check.InRangeEpsilon(%#v, %#v, %#v, %#v)"
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    f := func() error { return InRangeEpsilon(value, upper, lower, epsilon, irrelevant) }
    expectNoError(t, f, template, value, upper, lower, epsilon)
  }
}

func TestInRangeEpsilonReturnsErrWhenPassingValueOutOfRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {1, 0, 0 ,.1},
    []float64 {-.1, 0, 1 ,.05},
    []float64 {1.1, 0, 1 ,.05},
    []float64 {-.000001, 0, 1 ,.0000005},
    []float64 {1.000001, 0, 1 ,.0000005},
  }
  template := "check.InRangeEpsilon(%#v, %#v, %#v, %#v)"
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    f := func() error { return InRangeEpsilon(value, upper, lower, epsilon, irrelevant) }
    expectError(t, f, template, value, upper, lower, epsilon)
  }
}

func TestInRangeReturnsNilWhenPassingValueInRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {0, 0, 0},
    []float64 {0, 0, 1},
    []float64 {0, -1, 0},
    []float64 {0, -1, 1},
    []float64 {0, -.000001, .000001},
    []float64 {1000, 0, 1000},
    []float64 {-1000, -2000, 0},
    []float64 {.000000001, 0, 0},
  }
  template := "check.InRange(%#v, %#v, %#v)"
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    f := func() error { return InRange(value, upper, lower, irrelevant) }
    expectNoError(t, f, template, value, upper, lower)
  }
}

func TestInRangeReturnsErrWhenPassingValueOutOfRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {1, 0, 0},
    []float64 {-.1, 0, 1},
    []float64 {1.1, 0, 1},
    []float64 {-.001, 0, 1},
    []float64 {1.001, 0, 1},
  }
  template := "check.InRange(%#v, %#v, %#v)"
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    f := func() error { return InRange(value, upper, lower, irrelevant) }
    expectError(t, f, template, value, upper, lower)
  }
}

func TestTrueReturnsErrWithProperMessage(t *testing.T) {
  expected := "message"
  f := func() error { return True(false, expected) }
  expectMessage(t, f, expected, "True(false)")
}

func TestTrueReturnsErrWithProperMessageWithOneArgument(t *testing.T) {
  message, argument := "message %v", 12
  expected := fmt.Sprintf(message, argument)
  f := func() error { return True(false, message, argument) }
  expectMessage(t, f, expected, "True(false)")
}

func TestTrueReturnsErrWithProperMessageWithTwoArguments(t *testing.T) {
  message, argument0, argument1 := "message %v - %v", 12, 122
  expected := fmt.Sprintf(message, argument0, argument1)
  f := func() error { return True(false, message, argument0, argument1) }
  expectMessage(t, f, expected, "True(false)")
}


