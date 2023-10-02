package errorutil

import "fmt"

func AddParentError(err error, parent error) error {
	if err == nil {
		return nil
	}
	if parent == nil {
		return err
	}
	return fmt.Errorf("%w : %w", parent, err)
}

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
