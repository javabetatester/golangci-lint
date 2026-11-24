package nolint_filter

func RetErr2() error {
	return nil
}

func _() {
	RetErr2() //nolint:allfoo
}

func _() {
	RetErr2() //nolint:all // explanation
}

func _() {
	RetErr2() //nolint:bad3,errcheck
}

func _() {
	RetErr2() //nolint:bad4,all
}
