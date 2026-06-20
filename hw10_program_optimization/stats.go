package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (result DomainStat, err error) {
	result = make(DomainStat)
	scanner := bufio.NewScanner(r)
	reg, err := regexp.Compile("\\." + domain)
	if err != nil {
		return
	}

	var p fastjson.Parser
	for scanner.Scan() {
		var v *fastjson.Value
		v, err = p.ParseBytes(scanner.Bytes())
		if err != nil {
			return
		}
		email := string(v.GetStringBytes("Email"))
		if email == "" {
			err = errors.New("invalid email")
			return
		}
		if reg.Match([]byte(email)) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	err = scanner.Err()
	return
}
