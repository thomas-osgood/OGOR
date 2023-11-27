package generics

// function designed to create and initialize
// a new searchable int.
func NewSearchableInt() *SearchableInt {
	return new(SearchableInt)
}

// function designed to create and initialize
// a new searchable int32.
func NewSearchableInt32() *SearchableInt32 {
	return new(SearchableInt32)
}

// function designed to create and initialize
// a new searchable int32 slice.
func NewSearchableInt32Slice() *SearchableInt32Slice {
	return new(SearchableInt32Slice)
}

// function designed to create and initialize
// a new searchable int64.
func NewSearchableInt64() *SearchableInt64 {
	return new(SearchableInt64)
}

// function designed to create and initialize
// a new searchable int64 slice.
func NewSearchableInt64Slice() *SearchableInt64Slice {
	return new(SearchableInt64Slice)
}

// function designed to create and initialize
// a new searchable string.
func NewSearchableString() *SearchableString {
	return new(SearchableString)
}

// function designed to create and initialize
// a new searchable string slice.
func NewSearchableStringSlice() *SearchableStringSlice {
	return new(SearchableStringSlice)
}
