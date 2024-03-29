package generics

type CustomSlice interface {
	// interface restricting the type to the custom slices
	// that are defined within this document.
	SearchableStringSlice | SearchableInt32Slice | SearchableInt64Slice | SearchableIntSlice
}

type SearchObjectTypes interface {
	// interface outlining the various types of SearchObjects
	// that exist. this is used in the definition of the
	// SearchObject interface to make it generic.
	StringTypes | Int32Object | Int64Object | IntObject
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
	// function designed to clear/reset the object
	// to the value it is when initialized. for
	// slices, it will remove all elements from
	// the slice, it will zero out a string, and
	// set numbers to 0.
	//
	// example call:
	//
	//	searchableint32 = searchableint32.Clear()
	Clear() O
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
	// function designed to get the length of the object.
	Length() int
	// function designed to return the string
	// representstion of the given object.
	String() string
}

// generic definition of a slice that has certain, custom
// functions attached to it.
type SpecialSlice[T CustomSlice] interface {
	// function designed to combine two custom slices.
	// this will append the target object to the object
	// calling the Combine function.
	Combine(T) T
}

type IntObject interface {
	// generic definition of an int object. this
	// interface includes int and int slices.
	SearchableInt | SearchableIntSlice
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

// definition of a custom type of int
// that fits a SearchableObject definition.
type SearchableInt int

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
