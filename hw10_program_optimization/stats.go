package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

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
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(&u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	reader := bufio.NewReader(r)
	i := 0

	for {
		line, _, er := reader.ReadLine()
		var user User
		if er != nil {
			if errors.Is(er, io.EOF) {
				break
			}
			err = er
			return
		}
		_ = easyjson.Unmarshal(line, &user)

		result[i] = user
		i++
	}

	return
}

func countDomains(u *users, domain string) (DomainStat, error) {
	result := make(DomainStat, len(u))

	for _, user := range u {
		strDomain := strings.Split(user.Email, ".")

		if strDomain[len(strDomain)-1] == domain {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
