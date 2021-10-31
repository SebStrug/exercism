package erratum

import (
	"errors"
)

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(opener, input)
		}
		return errors.New(err.Error())
	}
	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case FrobError:
				resource.Defrob(x.defrobTag)
				err = x.inner
			case error:
				err = x
			}
		}
	}()
	resource.Frob(input)

	return err
}
