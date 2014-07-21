// vim: sw=2 ts=2 expandtab

package precond

import "fmt"

func CheckTrue(predicate bool,  message string, args ...interface{}) error {
  if !predicate {
    return fmt.Errorf(message, args)
  }
  return nil
}

func CheckFalse(antiPredicate bool, message string, args ...interface{}) error {
  return CheckTrue(!antiPredicate, message, args);
}

func CheckNotNil(value interface{}, message string, args ...interface{}) error {
  return CheckTrue(value != nil, message, args)
}

func CheckInRangeEpsilon(value float64, lower float64, upper float64,
        epsilon float64, message string, args ...interface{}) error {
  predicate := value - lower + epsilon > 0 && upper - value + epsilon > 0
  return CheckTrue(predicate, message, args)
}

func CheckInRange(value float64, lower float64, upper float64,
                  message string, args ...interface{}) error {
  return CheckInRangeEpsilon(value, lower, upper, .00001, message, args)
}

