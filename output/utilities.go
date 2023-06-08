package output

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// function designed to create and return a new
// outputter object that can be used later on.
func NewOutputter() (obj *Outputter, err error) {
	obj = &Outputter{}
	return obj, nil
}

// function designed to generate a random string between
// with a length that falls within a designated range. this
// can be useful for generating random file names or unique
// usernames and passwords.
func GenRandomName(minlen int, maxlen int) (randstr string, err error) {
	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var length int

	// validate min/max parameters
	if (minlen <= 0) || (maxlen <= 0) {
		return "", errors.New("min and max lengths mus be greater than zero")
	} else if minlen > maxlen {
		return "", errors.New("min length must be less than or equal to max length")
	}

	rand.Seed(time.Now().UnixMilli())

	length = rand.Intn(minlen + (maxlen - minlen))

	randstr = ""

	for i := 0; i < length; i++ {
		randstr = fmt.Sprintf("%s%s", randstr, string(charset[rand.Intn(len(charset))]))
	}

	return randstr, nil
}
