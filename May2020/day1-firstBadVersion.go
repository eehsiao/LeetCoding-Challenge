/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

var badVer int

// simulator isBadVersion API
func isBadVersion(version int) bool {
	return version >= badVer
}

func firstBadVersion(n int) int {
	i := 1
	for i < n {
		m := (i + n) / 2
		if isBadVersion(m) {
			if i == m-1 {
				if !isBadVersion(m - 1) {
					i = m
				}
				break
			}
			n = m - 1
		} else {
			if n == m+1 {
				if isBadVersion(n) {
					i = n
				} else {
					i = n + 1
				}
				break
			}
			i = m + 1
		}
	}

	return i
}
