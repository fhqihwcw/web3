package main

import (
	"fmt"
)

func main() {
	// sn := singleNumber([]int{4, 1, 2, 1, 2})
	// fmt.Println(sn)

	// ip := isPalindrome(1234)
	// fmt.Println(ip)

	vs := validString("(([))][]{}")
	fmt.Println(vs)
}

/*
*
** 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*
*/
func singleNumber(nums []int) int {
	var ans int
	ma := make(map[int]int) //声明并初始化
	for _, num := range nums {
		ma[num]++
	}
	for k, v := range ma {
		if v == 1 {
			ans = k
		}
	}
	return ans
}

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	res := reverse(x)
	if res == x {
		return true
	} else {
		return false
	}
}

// 数字反转函数
func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		fmt.Println(res)
		x /= 10
	}
	return res
}

/*
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
链接：https://leetcode-cn.com/problems/valid-parentheses/
*/

func validString(s string) bool {

	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		// 如果是右括号
		if left, ok := mapping[ch]; ok {
			// 栈为空或栈顶不是对应的左括号
			if len(stack) == 0 || stack[len(stack)-1] != left {
				return false
			}
			// 匹配则弹出栈顶
			stack = stack[:len(stack)-1]
		} else {
			// 左括号入栈
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}

/**
最长公共前缀
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/

/**
基本值类型
加一
难度：简单
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
链接：https://leetcode-cn.com/problems/plus-one/
**/

/**
引用类型：切片
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，
返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，
将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
**/

/**
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/

/**
基础
两数之和
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
链接：https://leetcode-cn.com/problems/two-sum/
*/
