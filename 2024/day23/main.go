package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
)

////////

type SetItem string

type Set []SetItem

func (s *Set) contains(i SetItem) bool {
	for _, si := range *s {
		if si == i {
			return true
		}
	}
	return false
}

func (s *Set) toString() string {
	x := "{"
	for _, v := range *s {
		x = fmt.Sprintf("%s%v,", x, v)
	}
	x = fmt.Sprintf("%s}", x)
	return x
}

func (s *Set) add(i SetItem) {
	if !s.contains(i) {
		*s = append(*s, i)
	}
}
func (s *Set) remove(i SetItem) {
	if s.contains(i) {
		ss := Set{}
		for _, si := range *s {
			if si != i {
				ss = append(ss, si)
			}
		}
		*s = ss
	}
}

func (s *Set) addAndReturn(i SetItem) *Set {
	ss := Set{}
	for _, si := range *s {
		ss.add(si)
	}
	if !ss.contains(i) {
		ss.add(i)
	}
	return &ss
}
func (s *Set) union(sBis Set) {
	for _, si := range sBis {
		s.add(si)
	}
}

func cross(sA Set, sB Set) *Set {
	s := Set{}
	for _, iA := range sA {
		if sB.contains(iA) {
			s.add(iA)
		}
	}
	for _, iB := range sB {
		if sA.contains(iB) {
			s.add(iB)
		}
	}
	return &s
}

////////

func readFile(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ls []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		ls = append(ls, sc.Text())
	}
	return ls, sc.Err()
}

func main() {
	inputFilePtr := flag.String("i", "", "Input file")

	flag.Parse()

	fmt.Printf("Input file: %s\n", *inputFilePtr)
	ls, err := readFile(*inputFilePtr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Lines:", len(ls))
	out := part1(ls)
	fmt.Printf("Output P1: %d\n", out)
	out = part2(ls)
	fmt.Printf("Output P2: %d\n", out)
}

type Connections map[string]*Set
type UniqueConnections map[string]bool

func (u *UniqueConnections) add(comp []string) {
	slices.Sort(comp)
	conn := strings.Join(comp, ",")
	if _, ok := (*u)[conn]; ok {
		return
	}
	(*u)[conn] = true
}

func parseInput(ls []string) Connections {
	c := make(Connections)
	for _, l := range ls {
		computers := strings.Split(l, "-")
		compA, compB := computers[0], computers[1]
		if _, ok := c[compA]; !ok {
			c[compA] = &Set{}
		}
		if _, ok := c[compB]; !ok {
			c[compB] = &Set{}
		}
		c[compA].add(SetItem(compB))
		c[compB].add(SetItem(compA))
	}
	return c
}

func (c *Connections) mutualConnections(a string, b string) []SetItem {
	mutual := []SetItem{}
	for _, connA := range *(*c)[a] {
		if string(connA) == b {
			continue
		}
		for _, connB := range *(*c)[b] {
			if string(connB) == a {
				continue
			}
			if connA == connB {
				mutual = append(mutual, connA)
			}
		}
	}
	return mutual
}

func part1(ls []string) int {
	c := parseInput(ls)

	u := make(UniqueConnections)
	for compA, connsA := range c {
		if !strings.HasPrefix(compA, "t") {
			continue
		}
		for _, compB := range *connsA {
			if compA == string(compB) {
				continue
			}
			for _, mutual := range c.mutualConnections(compA, string(compB)) {
				u.add([]string{compA, string(compB), string(mutual)})
			}
		}
	}
	//for conn := range u {
	//	fmt.Println(conn)
	//}
	return len(u)
}

func (c *Connections) toString() string {
	x := ""
	for k, v := range *c {
		x = fmt.Sprintf("%s%v:%s  ", x, k, v.toString())
	}
	return x
}

func (c *Connections) bn(r Set, p Set, x Set) Set {
	if len(p) == 0 && len(x) == 0 {
		return r
	}

	clique := Set{}
	slices.Sort(p)
	for _, v := range p {
		r_v := r.addAndReturn(v)
		p_v := cross(p, *(*c)[string(v)])
		x_v := cross(x, *(*c)[string(v)])
		res := c.bn(*r_v, *p_v, *x_v)
		if len(res) > len(clique) {
			clique = res
		}
		p_v.remove(v)
		x_v.add(v)
	}
	//	p_v.remove(v)
	//	x_v.add(v)
	//	res := c.bn(*r_v, *p_v, *x_v)
	//	if len(res) > len(s) {
	//		s = res
	//	}
	//}
	return clique
}

func part2(ls []string) int {
	c := parseInput(ls)
	//c := Connections{
	//	"1": {"2", "3"},
	//	"2": {"1", "3"},
	//	"3": {"1", "2", "4"},
	//	"4": {"3"},
	//	"5": {},
	//}

	p := Set{}
	for item := range c {
		p.add(SetItem(item))
	}
	x := c.bn(Set{}, p, Set{})
	fmt.Printf("MAX CLIQUE %s\n", x)
	return len(x)
}
