package generators

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/thomas-osgood/OGOR/networking/proxyscrape"
)

// function designed to generate a random string
// with a length that falls within a designated range. this
// can be useful for generating random file names or unique
// usernames and passwords.
func GenRandomName(minlen int, maxlen int) (randstr string, err error) {
	var baseval *big.Int
	var biglen *big.Int
	var bigmin *big.Int
	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var length int
	var randidx *big.Int

	// validate min/max parameters
	if (minlen <= 0) || (maxlen <= 0) {
		return "", errors.New("min and max lengths mus be greater than zero")
	} else if minlen > maxlen {
		return "", errors.New("min length must be less than or equal to max length")
	}

	// this is the number that will be used to generate
	// the random number. this is the difference of the
	// max value and min value because the final random
	// number will be calculated by adding the min value
	// so the number falls within the range MIN <= x <= MAX.
	baseval = big.NewInt(int64(maxlen - minlen))

	// convert the minimum value to a big.Int so it can be
	// used to adjust the randomly generated length.
	bigmin = big.NewInt(int64(minlen))

	// use the crypto/rand library to generate a length
	// for the string.
	biglen, err = rand.Int(rand.Reader, big.NewInt(baseval.Int64()))
	if err != nil {
		return "", fmt.Errorf("error generating the length: %s", err)
	}

	// adjust the generated number to fit within the range.
	biglen = biglen.Add(biglen, bigmin)

	length = int(biglen.Int64())

	randstr = ""

	for i := 0; i < length; i++ {
		// calculate the random index to choose. if
		// there is an error, choose index 0.
		randidx, err = rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			randidx = big.NewInt(0)
		}
		// append the char at the randomly generated
		// index to the randomly generated string.
		randstr = fmt.Sprintf("%s%c", randstr, charset[randidx.Int64()])
	}

	return randstr, nil
}

// function designed to continually feed proxies into a channel
// so the end-user can use them to make requests. this can allow
// a user to have a different IP address each request.
func ProxyGenerator(ps *proxyscrape.ProxyScraper, c chan string, wg *sync.WaitGroup, exit *bool) (err error) {
	if wg != nil {
		wg.Add(1)
		defer wg.Done()
	}

	defer close(c)

	var idx int = 0
	var max int = 0

	// if no proxyscraper object was passed in, create a new one
	// with a default configuration.
	if ps == nil {
		ps, err = proxyscrape.NewProxyScraper()
		if err != nil {
			return err
		}
	}

	// check the lenght of the proxies slice to make sure it
	// has been populated. if it is empty, call the GetProxies
	// function to populate it.
	if len(ps.Proxies.Proxies) < 1 {
		err = ps.GetProxies()
		if err != nil {
			return err
		}
	}

	// set length to use for mod division.
	max = len(ps.Proxies.Proxies)

	// continually loop through the pulled down proxies and
	// feed them to the channel. if the "exit" flag is set
	// to true, break the loop and exit.
	for {
		if *exit {
			break
		}
		c <- ps.Proxies.Proxies[idx]
		idx = (idx + 1) % max
	}

	return nil
}

// function designed to continually loop through a given slice and
// feed the current item into the channel.
func SliceIterator[T any](slice []T, commschan chan T, exit *bool) (err error) {
	defer close(commschan)

	var curitem T
	var idx int = 0

	// return error if slice is empty.
	if len(slice) < 1 {
		return errors.New("cannot iterate over empty slice")
	}

	// continuall loop through target slice and feed the current item
	// into the channel. if the exit flag is set, break the loop, close
	// the channel and return without error.
	for {
		if *exit {
			break
		}

		curitem = slice[idx]
		commschan <- curitem

		// keep idx within the bounds of the slice by taking the
		// modulus result of idx+1.
		idx = (idx + 1) % len(slice)
	}

	return nil
}
