package util

import "sync"

type UT struct {
	F0   func(string) map[string]string
	M    map[string]string
	Name string
	sync.Mutex
}

func NewUT() *UT {
	ut := &UT{}
	ut.M = map[string]string{}
	ut.Name = "Initial"
	ut.F0 = func(s string) map[string]string {
		m := map[string]string{}
		m[s] = "one"
		return m
	}

	return ut
}

func (ut *UT) DoSomething(s string) {
	ut.Lock()
	defer ut.Unlock()

	ut.M[s] = "two"

}
