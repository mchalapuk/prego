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

package check

import "testing"
import "fmt"
import "strings"

const (
  irrelevant = "irrelevant in this test"
)

func TestTrueReturnsNilWhenPassingTrue(t *testing.T) {
  err := True(true, irrelevant)
  if err != nil {
    t.Errorf("Expected True(%+v) to return nil,got %+v", true, err);
  }
}

func TestTrueReturnsErrWhenPassingFalse(t *testing.T) {
  err := True(false, irrelevant)
  if err == nil {
    t.Errorf("Expected True(%+v) to return error, got nil", false);
  }
}

func TestFlaseReturnsNilWhenPassingFalse(t *testing.T) {
  err := False(false, irrelevant)
  if err != nil {
    t.Errorf("Expected False(%+v) to return nil,got %+v", false, err);
  }
}

func TestFalseReturnsErrWhenPassingTrue(t *testing.T) {
  err := False(true, irrelevant)
  if err == nil {
    t.Errorf("Expected False(%+v) to return error, got nil", false);
  }
}

type Test struct {
}

func TestNilReturnsNilWhenPassingNil(t *testing.T) {
  err := Nil(nil, irrelevant)
  if err != nil {
    t.Errorf("Expected Nil(%+v) to return nil, got %+v", nil, err);
  }
}

func TestNilReturnsErrWhenPassingNotNil(t *testing.T) {
  value := &Test{}
  err := Nil(value, irrelevant)
  if err == nil {
    t.Errorf("Expected NotNil(%+v) to return error, got nil", value);
  }
}

func TestNotNilReturnsNilWhenPassingNotNil(t *testing.T) {
  value := &Test{}
  err := NotNil(value, irrelevant)
  if err != nil {
    t.Errorf("Expected NotNil(%+v) to return nil, got %+v", value, err);
  }
}

func TestNotNilReturnsErrWhenPassingNil(t *testing.T) {
  err := NotNil(nil, irrelevant)
  if err == nil {
    t.Errorf("Expected NotNil(%+v) to return error, got nil", nil);
  }
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
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    err := InRangeEpsilon(value, upper, lower, epsilon, irrelevant)
    if err != nil {
      t.Errorf("Expected InRangeEpsilon(%v, %v, %v, %v) to return nil, got %+v",
               value, upper, lower, epsilon, err);
    }
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
  for _, test := range(tests) {
    value, upper, lower, epsilon := test[0], test[1], test[2], test[3]
    err := InRangeEpsilon(value, upper, lower, epsilon, irrelevant)
    if err == nil {
      t.Errorf("Expected InRangeEpsilon(%v, %v, %v, %v) to return error, got nil",
               value, upper, lower, epsilon);
    }
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
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    err := InRange(value, upper, lower, irrelevant)
    if err != nil {
      t.Errorf("Expected InRange(%v, %v, %v) to return nil, got %+v",
               value, upper, lower, err);
    }
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
  for _, test := range(tests) {
    value, upper, lower := test[0], test[1], test[2]
    err := InRange(value, upper, lower, irrelevant)
    if err == nil {
      t.Errorf("Expected InRange(%v, %v, %v) to return error, got nil",
               value, upper, lower);
    }
  }
}

func TestTrueReturnsErrWithProperMessage(t *testing.T) {
  expected := "message"
  err := True(false, expected)
  actual := fmt.Sprintf("%s", err)
  if strings.Index(actual, expected) != 0 {
    t.Errorf("Expected True(%+v) to return error with message '%v', got '%v'",
            false, "message", actual)
  }
}

func TestTrueReturnsErrWithProperMessageWithOneArgument(t *testing.T) {
  message := "message %v"
  argument := 12
  expected := fmt.Sprintf(message, argument)
  err := True(false, expected, message, argument)
  actual := fmt.Sprintf("%s", err)
  if strings.Index(actual, expected) != 0 {
    t.Errorf("Expected True(%+v) to return error with message '%v', got '%v'",
            false, "message", actual)
  }
}

func TestTrueReturnsErrWithProperMessageWithTwoArguments(t *testing.T) {
  message := "message %v - %v"
  argument0 := 12
  argument1 := 122
  expected := fmt.Sprintf(message, argument0, argument1)
  err := True(false, expected, message, argument0, argument1)
  actual := fmt.Sprintf("%s", err)
  if strings.Index(actual, expected) != 0 {
    t.Errorf("Expected True(%+v) to return error with message '%v', got '%v'",
            false, "message", actual)
  }
}

