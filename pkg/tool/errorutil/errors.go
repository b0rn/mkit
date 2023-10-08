package errorutil

import "fmt"

// Formats two errors using fmt.Errorf("%w : %w")
func AddParentError(err error, parent error) error {
	if err == nil {
		return nil
	}
	if parent == nil {
		return err
	}
	return fmt.Errorf("%w : %w", parent, err)
}

// Adds `parent` error to each error in `errs`
func AddParentErrorToErrors(errs []error, parent error) []error {
	ret := []error{}
	for _, e := range errs {
		if e == nil {
			continue
		}
		ret = append(ret, AddParentError(e, parent))
	}
	return ret
}
