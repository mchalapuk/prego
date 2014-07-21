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

import "fmt"

/*
  Returns nil if given predicate is true, otherwise returns error
*/
func CheckTrue(predicate bool, message string, args ...interface{}) error {
  if !predicate {
    return fmt.Errorf(message, args)
  }
  return nil
}

/*
  Returns nil if given predicate is false, otherwise returns error
*/
func CheckFalse(antiPredicate bool, message string, args ...interface{}) error {
  return CheckTrue(!antiPredicate, message, args);
}

/*
  Returns nil if given value is nil, otherwise returns error
*/
func CheckNil(value interface{}, message string, args ...interface{}) error {
  return CheckTrue(value == nil, message, args)
}

/*
  Returns nil if given value is not nil, otherwise returns error
*/
func CheckNotNil(value interface{}, message string, args ...interface{}) error {
  return CheckTrue(value != nil, message, args)
}

/*
  Returns nil if given value is contained in given range &lt;lower, upper&gt;,
  otherwise returns error. Uses given epsilon for float comparison.
*/
func CheckInRangeEpsilon(value float64, lower float64, upper float64,
        epsilon float64, message string, args ...interface{}) error {
  predicate := value - lower + epsilon > 0 && upper - value + epsilon > 0
  return CheckTrue(predicate, message, args)
}

/*
  Returns nil if given value is contained in given range &lt;lower, upper&gt;,
  otherwise returns error. Uses epsilon value of 0.00001.
*/
func CheckInRange(value float64, lower float64, upper float64,
                  message string, args ...interface{}) error {
  return CheckInRangeEpsilon(value, lower, upper, .00001, message, args)
}

