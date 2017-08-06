package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (ret error) {
	r, e := o()
	for ; e != nil; r, e = o() {
		switch e.(type) {
		default:
			return e
		case TransientError:
		}
	}

	defer r.Close()

	defer func() {
		if x := recover(); x != nil {
			if frobErr, ok := x.(FrobError); ok {
				r.Defrob(frobErr.defrobTag)
			}
			ret, _ = x.(error)
		}
	}()

	r.Frob(input)

	return nil
}
