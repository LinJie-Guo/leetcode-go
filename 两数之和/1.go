package main

/**
 * 两数之和
 */
func twoSum(nums []int, target int) []int {
    count := len(nums)

    for i := 0; i < count; i++ {
        for j := i+1; j < count; j++ {
            if target == nums[i] + nums[j] {
                return []int{i,j}
            }
        }
    }
    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    twoSum(nums, target)
}