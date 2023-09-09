package files

func read[T any](file *File, fn func() (*T, error)) (*T, error) {
	err := file.openRead()
	if err != nil {
		return nil, err
	}
	defer file.close()
	return fn()
}

func write[T any](file *File, trunc bool, fn func() (*T, error)) (*T, error) {
	if err := file.mkparent(); err != nil {
		return nil, err
	}
	if err := file.openWrite(trunc); err != nil {
		return nil, err
	}
	defer file.close()

	if trunc {
		if err := file.clear(); err != nil {
			return nil, err
		}
	}
	return fn()
}
