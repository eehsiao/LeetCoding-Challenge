// Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
// Examples:
// s = "leetcode"
// return 0.
// s = "loveleetcode",
// return 2.
// Note: You may assume the string contain only lowercase letters.
// its like https://github.com/eehsiao/30-Day-LeetCoding-Challenge/blob/master/week4-firstUniqueNumber.go

package main

import (
	"strings"
)

type linked struct {
	prv  *linked
	next *linked
	char string
	idx  int
}

type FirstUnique struct {
	head *linked
	last *linked
	m    map[string]*linked
}

func firstUniqChar(s string) int {
	sSlice := strings.Split(s, "")
	firstUnique := FirstUnique{
		head: nil,
		last: nil,
		m:    make(map[string]*linked),
	}

	for i, v := range sSlice {
		firstUnique.Add(i, v)
	}

	return firstUnique.ShowFirstUnique()
}

func (this *FirstUnique) ShowFirstUnique() int {
	num := -1
	if this.head != nil {
		num = this.head.idx
	}

	return num
}

func (this *FirstUnique) Add(idx int, value string) {
	if this.head == nil {
		if l, ok := this.m[value]; ok {
			if l != nil {
				this.remove(l)
				this.m[value] = nil
			}
		} else {
			this.head = &linked{
				prv:  nil,
				next: nil,
				char: value,
				idx:  idx,
			}
			this.last = this.head
			this.m[value] = this.head
		}
	} else {
		if l, ok := this.m[value]; ok {
			if l != nil {
				this.remove(l)
				this.m[value] = nil
			}
		} else {
			this.last.next = &linked{
				prv:  this.last,
				next: nil,
				char: value,
				idx:  idx,
			}
			this.last = this.last.next
			this.m[value] = this.last
		}

	}
}

func (this *FirstUnique) remove(l *linked) {
	if this.head == l && this.last == l {
		this.head, this.last, l = nil, nil, nil
	} else {
		if l.prv != nil {
			l.prv.next = l.next
		} else {
			l.next.prv = nil
			this.head = l.next
		}

		if l.next != nil {
			l.next.prv = l.prv
		} else {
			l.prv.next = nil
			this.last = l.prv
		}
		l = nil
	}
}
