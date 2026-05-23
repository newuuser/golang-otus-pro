package hw10programoptimization

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (result DomainStat, err error) {
	result = make(DomainStat)
	r = bufio.NewReader(r)
	scanner := bufio.NewScanner(r)
	reg, err := regexp.Compile("\\." + domain)
	if err != nil {
		return
	}

	var p fastjson.Parser
	for scanner.Scan() {
		v, err := p.ParseBytes(scanner.Bytes())
		if err != nil {
			return nil, err
		}
		email := string(v.GetStringBytes("Email"))
		if reg.Match([]byte(email)) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}
	return
}
