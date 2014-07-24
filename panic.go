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

import "./check"

func panicIfError(check func() error) {
  if err := check(); err != nil {
    panic(err)
  }
}

/*
  Panics if given predicate is not true.
*/
func True(predicate bool, message string, args ...interface{}) {
  panicIfError(func() error { return check.True(predicate, message, args...) })
}

/*
  Panics if given predicate is not false.
*/
func False(antiPredicate bool, message string, args ...interface{}) {
  panicIfError(func() error {
      return check.False(antiPredicate, message, args...)
      })
}

/*
  Panics if given value is not nil.
*/
func Nil(value interface{}, message string, args ...interface{}) {
  panicIfError(func() error { return check.Nil(value, message, args...) })
}

/*
  Panics if given value is nil.
*/
func NotNil(value interface{}, message string, args ...interface{}) {
  panicIfError(func() error { return check.NotNil(value, message, args...) })
}

/*
  Panics if given value is not contained in given range &lt;lower, upper&gt;
  Uses given epsilon for float comparison.
*/
func InRangeEpsilon(value float64, lower float64, upper float64,
        epsilon float64, message string, args ...interface{}) {
  panicIfError(func() error {
      return check.InRangeEpsilon(value,lower,upper,epsilon,message,args...)
      })
}

/*
  Throws if given value is not contained in given range &lt;lower, upper&gt;.
  Uses epsilon value of 0.00001 for float comparison.
*/
func InRange(value float64, lower float64, upper float64,
             message string, args ...interface{}) {
  panicIfError(func() error {
      return check.InRange(value, lower, upper, message, args...)
      })
}

