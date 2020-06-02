// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
// You may assume that the array is non-empty and the majority element always exist in the array.
// Example 1:
// Input: [3,2,3]
// Output: 3
// Example 2:
// Input: [2,2,1,1,1,2,2]
// Output: 2

package main

type elementLinked struct {
	prv  *elementLinked
	next *elementLinked
	num  int
	cnt  int
}

type MajorityElement struct {
	head *elementLinked
	last *elementLinked
	m    map[int]*elementLinked
	mLen int
}

func majorityElement(nums []int) int {
	mElement := MajorityElement{
		head: nil,
		last: nil,
		m:    make(map[int]*elementLinked),
		mLen: len(nums) / 2,
	}

	for _, v := range nums {
		if mElement.Add(v) {
			break
		}
	}

	return mElement.last.num
}

func (this *MajorityElement) Add(value int) bool {
	if this.head == nil {
		this.head = &elementLinked{
			prv:  nil,
			next: nil,
			num:  value,
			cnt:  1,
		}
		this.last = this.head
		this.m[value] = this.head
	} else {
		if l, ok := this.m[value]; ok {
			if l != nil {
				l.cnt++
				//if cnt> last.cnt the move to last
				if l != this.last && l.cnt > this.last.cnt {
					if l == this.head {
						this.head = l.next
						l.next.prv = nil
						l.next = nil
						l.prv = this.last
						this.last.next = l
						this.last = l
					} else {
						l.prv.next = l.next
						l.next.prv = l.prv
						l.prv = this.last
						this.last.next = l
						l.next = nil
						this.last = l
					}
				}
			}
		} else {
			this.head = &elementLinked{
				prv:  nil,
				next: this.head,
				num:  value,
				cnt:  1,
			}
			this.head.next.prv = this.head
			this.m[value] = this.head
		}
	}

	if this.last.cnt > this.mLen {
		return true
	}

	return false
}
