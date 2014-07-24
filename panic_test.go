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

// vim: sw=2 ts=2 expandtab

package precond

import "testing"
import "fmt"
import "strings"

const (
  irrelevant = "irrelevant in this test"
)

func expectSerenity(t *testing.T, precondCheckStr string, test func()) {
  defer func() {
    if err := recover(); err != nil {
      t.Errorf("Unexpected panic in %s, got error %v ", precondCheckStr, err)
    }
  }()
  test()
}

func expectPanic(t *testing.T, precondCheckStr string, test func()) {
  defer func() {
    if err := recover(); err == nil {
      t.Errorf("Expected panic in %s not experienced", precondCheckStr)
    }
  }()
  test()
}

func expectPanicMessage(t *testing.T,
                        precondCheckStr string,
                        expectedMessage string,
                        test func()) {
  defer func() {
    err := recover()
    if err == nil {
      t.Errorf("Expected panic from %s not experienced", precondCheckStr)
      return
    }
    actualMessage := fmt.Sprintf("%s", err)
    if strings.Index(actualMessage, expectedMessage) != 0 {
      t.Errorf("Expected %s to panic error with message '%v',"+
               " got '%v'", precondCheckStr, expectedMessage, actualMessage)
    }
  }()
  test()
}

func TestPrecondTrueDoesNotPanicWhenPassingTrue(t *testing.T) {
  expectSerenity(t, "precond.True(true)", func() { True(true, irrelevant) })
}

func TestPrecondTruePanicsWhenPassingFalse(t *testing.T) {
  expectPanic(t, "precond.True(false)", func() { True(false, irrelevant) })
}

func TestPrecondFalseDoesNotPanicWhenPassingFalse(t *testing.T) {
  expectSerenity(t, "precond.False(false)", func() { False(false, irrelevant) })
}

func TestPrecondFalsePanicsWhenPassingTrue(t *testing.T) {
  expectPanic(t, "precond.False(true)", func() { False(true, irrelevant) })
}

func TestPrecondNilDoenNotPanicWhenPassingNil(t *testing.T) {
  expectSerenity(t, "precond.Nil(nil)", func() { Nil(nil, irrelevant) })
}

type Test struct {
}

func TestPrecondNilPanicsWhenPassingNotNil(t *testing.T) {
  value := &Test{}
  checkStr := fmt.Sprintf("precond.Nil(%v)", value)
  expectPanic(t, checkStr, func() { Nil(value, irrelevant) })
}

func TestPrecondNotNilDoesNotPanicWhenPassingNotNil(t *testing.T) {
  value := &Test{}
  checkStr := fmt.Sprintf("precond.NotNil(%v)", value)
  expectSerenity(t, checkStr, func() { NotNil(value, irrelevant) })
}

func TestPrecondNotNilPanicWhenPassingNil(t *testing.T) {
  expectPanic(t, "precond.NotNil(nil)", func() { NotNil(nil, irrelevant) })
}

func TestPrecondInRangeEpsilonDoesNotPanicWhenPassingValueInRange(t *testing.T){
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
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    precondCheckStr := fmt.Sprintf("precond.InRangeEpsilon(%v, %f, %f, %f)",
                                  value, upper, lower, epsilon)
    expectSerenity(t, precondCheckStr,
        func() { InRangeEpsilon(value, upper, lower, epsilon, irrelevant) })
  }
}

func TestPrecondInRangeEpsilonPanicsWhenPassingValueOutOfRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {1, 0, 0 ,.1},
    []float64 {-.1, 0, 1 ,.05},
    []float64 {1.1, 0, 1 ,.05},
    []float64 {-.000001, 0, 1 ,.0000005},
    []float64 {1.000001, 0, 1 ,.0000005},
  }
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    precondCheckStr := fmt.Sprintf("precond.InRangeEpsilon(%v, %f, %f, %f)",
                                  value, upper, lower, epsilon)
    expectPanic(t, precondCheckStr,
        func() { InRangeEpsilon(value, upper, lower, epsilon, irrelevant) })
  }
}

func TestPrecondInRangeReturnsNilWhenPassingValueInRange(t *testing.T) {
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
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    precondCheckStr := fmt.Sprintf("precond.InRange(%v, %f, %f)",
                                  value, upper, lower)
    expectSerenity(t, precondCheckStr,
        func() { InRange(value, upper, lower, irrelevant) })
  }
}

func TestPrecondInRangePanicsWhenPassingValueOutOfRange(t *testing.T) {
  tests := [][]float64 {
    []float64 {1, 0, 0},
    []float64 {-.1, 0, 1},
    []float64 {1.1, 0, 1},
    []float64 {-.001, 0, 1},
    []float64 {1.001, 0, 1},
  }
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    precondCheckStr := fmt.Sprintf("precond.InRange(%v, %f, %f)",
                                  value, upper, lower)
    expectPanic(t, precondCheckStr,
        func() { InRange(value, upper, lower, irrelevant) })
  }
}

func TestPrecondTruePanicsErrWithProperMessage(t *testing.T) {
  expected := "message"
  expectPanicMessage(t, "precond.True(false)", expected,
                     func() { True(false, expected) })
}

func TestPrecondTruePanicsErrWithProperMessageWithOneArgument(t *testing.T) {
  message, arg := "message %v", 12
  expected := fmt.Sprintf(message, arg)
  expectPanicMessage(t, "precond.True(false)", expected,
                     func() { True(false, message, arg) })
}

func TestPrecondTruePanicsErrWithProperMessageWithTwoArguments(t *testing.T) {
  message, arg0, arg1 := "message %v -%v", 12, 122
  expected := fmt.Sprintf(message, arg0, arg1)
  expectPanicMessage(t, "precond.True(false)", expected,
                     func() { True(false, message, arg0, arg1) })
}

