package generics

type SearchObjectTypes interface {
	// interface outlining the various types of SearchObjects
	// that exist. this is used in the definition of the
	// SearchObject interface to make it generic.
	StringTypes | Int32Object | Int64Object
}

type SearchType interface {
	// type definition outlining the type of search
	// available for the given object. this will apply
	// to SearchableStrings, SearchableStringSlices,
	// SearchableInt32s, SearchableInt32Slices, SearchableInt64s,
	// and SearchableInt64Slices.
	//
	// this allows all custom types listed above to share
	// the same basic functionality and interfaces.
	int | int32 | int64 | string
}

type StringTypes interface {
	// generic used to restrict the types that are attached
	// to this type to string and string slice.
	SearchableString | SearchableStringSlice
}

// generic definition of a SearchObject. this outlines
// the functions a struct must have for it to be considered
// this type of generic.
type SearchObject[T SearchType, O SearchObjectTypes] interface {
	// function designed to append to a
	// SearchObjectType.
	Append(T) O
	// function that determines whether the
	// value passed in is within the searcher
	// object. if it does not exist, an error
	// will be returned.
	In(T) error
	// function that determines the index of
	// the passed in value. if it is not found
	// an error will be returned along with a
	// negative number.
	IndexOf(T) (int, error)
	// function designed to return the string
	// representstion of the given object.
	String() string
}

type Int32Object interface {
	// generic definition of an int32 object. this
	// interface includes int32 and int32 slices.
	SearchableInt32 | SearchableInt32Slice
}

type Int64Object interface {
	// generic definition of an int64 object. this
	// interface includes int64 and int64 slices.
	SearchableInt64 | SearchableInt64Slice
}

// definition of a custom type of int slice
// that fits a SearchableObject definition.
type SearchableIntSlice []int

// definition of a custom type of int32
// that fits a SearchableObject definition.
type SearchableInt32 int32

// definition of a custom type of int32 slice
// that fits a SearchableObject definition.
type SearchableInt32Slice []int32

// definition of a custom type of int64
// that fits a SearchableObject definition.
type SearchableInt64 int64

// definition of a custom type of int32 slice
// that fits a SearchableObject definition.
type SearchableInt64Slice []int64

// definition of a custom type of string that
// fits a SearchableObject definition.
type SearchableString string

// definition of a custom type of string slice
// that fits a SearchableObject definition.
type SearchableStringSlice []string
