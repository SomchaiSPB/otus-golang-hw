package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"sync"

	"github.com/mailru/easyjson"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	userCh := make(chan *User)
	errCh := make(chan error)
	wg := sync.WaitGroup{}

	go func() {
		defer close(userCh)
		defer close(errCh)
		wg.Wait()
	}()

	wg.Add(1)
	go getUsers(r, userCh, errCh, &wg)

	res := countDomains(userCh, domain)

	return res, nil
}

func getUsers(r io.Reader, ch chan *User, errCh chan error, w *sync.WaitGroup) {
	defer w.Done()
	reader := bufio.NewReader(r)

	for {
		line, _, er := reader.ReadLine()
		var user User
		if er != nil {
			if errors.Is(er, io.EOF) {
				break
			}
			errCh <- er
			return
		}
		_ = easyjson.Unmarshal(line, &user)

		ch <- &user
	}
}

func countDomains(u chan *User, domain string) DomainStat {
	result := make(DomainStat, len(u))

	for user := range u {
		strDomain := strings.Split(user.Email, ".")

		if strDomain[len(strDomain)-1] == domain {
			key := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[key]++
		}
	}

	return result
}
