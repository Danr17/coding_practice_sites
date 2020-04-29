package twosum

func twoSum(nums []int, target int) []int {
	indices := []int{}

	for i := 0; i < len(nums); i++ {
		for j := range nums {
			if i == j {
				continue
			} else if nums[i]+nums[j] == target {
				indices = append(indices, i, j)
				return indices

			}

		}
	}
	return nil

}
