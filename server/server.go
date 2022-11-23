package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Interpreter struct {
	OriginalInput         string
	Cmd                   string
	Key                   string
	Value                 string
	OptionalArgument      string
	OptionalArgumentValue int
}

type GlobalCache struct {
	HashMap    map[string]string
	HashMapTTL map[string]int
}

func InitCache() *GlobalCache {
	var cache GlobalCache
	cache.HashMap = make(map[string]string)
	cache.HashMapTTL = make(map[string]int)
	return &cache
}

var Cache *GlobalCache

func Server() {
	l, err := net.Listen("tcp", ":1210")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Starting cheena key-value server at localhost:1210")
	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// Initialize Global Cache
	Cache = InitCache()
	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		Handler(conn, netData)
	}

}

func Handler(conn net.Conn, input string) {
	fmt.Println("Received Input" + input)
	input = strings.TrimSpace(input)

	iptr := &Interpreter{OriginalInput: input}
	iptr.Parser(input)
	result := iptr.Execute()
	conn.Write([]byte(result))
	conn.Close()
}

func (iptr *Interpreter) Execute() string {
	switch iptr.Cmd {
	case SET:
		if iptr.Key == "" {
			return "Error! SET <key> can not be empty"
		}
		if iptr.Value == "" {
			return "Error! SET's <key> must have a value"
		} else {
			Cache.HashMap[iptr.Key] = iptr.Value
		}
		if iptr.OptionalArgument == "EX" {
			if iptr.OptionalArgumentValue == 0 {
				return "Error! TTL value is missing."
			}
			Cache.HashMapTTL[iptr.Key] = iptr.OptionalArgumentValue
		}
	case GET:
		if iptr.Key == "" {
			return "Error! GET <key> can not be empty"
		}
		if val, ok := Cache.HashMap[iptr.Key]; ok {
			return val
		}
		return "Error! " + iptr.Key + " does not exists"
	}
	return ""
}
