// Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
// Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
// Note: You are not suppose to use the library's sort function for this problem.
// Example:
// Input: [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
// Follow up:
// A rather straight forward solution is a two-pass algorithm using counting sort.
// First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
// Could you come up with a one-pass algorithm using only constant space?

package main

func sortColors(nums []int) {
	red, blue, i := 0, len(nums)-1, 0

	for i <= blue {
		if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			i++
			red++
		} else if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else {
			i++
		}
	}

}

func sortColors_wrong(nums []int) {
	red, blue, i := 0, len(nums)-1, 0

	for i = 0; i < len(nums)/2; i++ {
		if nums[i] == 2 {
			blue--
			nums[i], nums[blue] = nums[blue], nums[i]
			if nums[i] == 0 && i != red {
				red++
				nums[i], nums[red] = nums[red], nums[i]
			}
		} else if nums[i] == 0 {
			red++
			nums[i], nums[red] = nums[red], nums[i]
			if nums[i] == 2 && i != blue {
				blue--
				nums[i], nums[blue] = nums[blue], nums[i]
			}
		}

		if nums[len(nums)-1-i] == 0 {
			red++
			nums[len(nums)-1-i], nums[red] = nums[red], nums[len(nums)-1-i]
			if nums[len(nums)-1-i] == 2 {
				blue--
				nums[len(nums)-1-i], nums[blue] = nums[blue], nums[len(nums)-1-i]
			}
		} else if nums[len(nums)-1-i] == 0 {
			blue--
			nums[len(nums)-1-i], nums[blue] = nums[blue], nums[len(nums)-1-i]
			if nums[len(nums)-1-i] == 2 {
				red++
				nums[len(nums)-1-i], nums[red] = nums[red], nums[len(nums)-1-i]
			}
		}
		// fmt.Println(nums, red, blue)
	}
	if len(nums)/2*2 != len(nums) {
		if nums[i] == 2 && i < blue {
			blue--
			nums[i], nums[blue] = nums[blue], nums[i]
		} else if nums[i] == 0 && i > red {
			red++
			nums[i], nums[red] = nums[red], nums[i]
		} else if nums[i] == 1 {
			if red == -1 && nums[i] < nums[red+1] {
				red++
				nums[i], nums[red] = nums[red], nums[i]
			} else if blue == len(nums) && nums[i] > nums[blue-1] {
				blue--
				nums[i], nums[blue] = nums[blue], nums[i]
			}
		}
		// fmt.Println(nums, red, blue)
	}
}
