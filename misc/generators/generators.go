package generators

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/thomas-osgood/OGOR/networking/proxyscrape"
)

// function designed to generate a random string
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

	rand.Seed(time.Now().UnixMicro())

	length = minlen + rand.Intn(maxlen-minlen)

	randstr = ""

	for i := 0; i < length; i++ {
		randstr = fmt.Sprintf("%s%s", randstr, string(charset[rand.Intn(len(charset))]))
	}

	return randstr, nil
}

// function designed to continually feed proxies into a channel
// so the end-user can use them to make requests. this can allow
// a user to have a different IP address each request.
func ProxyGenerator(ps *proxyscrape.ProxyScraper, c chan string, wg *sync.WaitGroup, exit *bool) (err error) {
	wg.Add(1)
	defer wg.Done()
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
