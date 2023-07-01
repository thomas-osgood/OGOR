package dnsenum

// type alias for a function that can set
// options in an EnumOpts object. this is
// used in the NewEnumerator function.
type EnumOptsFunc func(*EnumOpts) error
