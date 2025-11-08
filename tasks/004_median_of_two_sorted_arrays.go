package tasks

/*
Median of Two Sorted Arrays
Difficulty: Hard

Given two sorted arrays nums1 and nums2 of size m and n respectively,
return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

Example 1:
  Input: nums1 = [1,3], nums2 = [2]
  Output: 2.00000
  Explanation: merged array = [1,2,3], median is 2.

Example 2:
  Input: nums1 = [1,2], nums2 = [3,4]
  Output: 2.50000
  Explanation: merged array = [1,2,3,4], median is (2 + 3) / 2 = 2.5.

Constraints:
  - nums1.length == m
  - nums2.length == n
  - 0 <= m <= 1000
  - 0 <= n <= 1000
  - 1 <= m + n <= 2000
  - -10^6 <= nums1[i], nums2[i] <= 10^6
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	lengthFirst, lengthSecond := len(nums1), len(nums2)
	totalLeftElements := (lengthFirst + lengthSecond + 1) / 2

	searchStart, searchEnd := 0, lengthFirst

	for searchStart <= searchEnd {
		partitionFirst := (searchStart + searchEnd) / 2
		partitionSecond := totalLeftElements - partitionFirst

		maxLeftFirst := -1 << 31
		if partitionFirst > 0 {
			maxLeftFirst = nums1[partitionFirst-1]
		}

		minRightFirst := 1<<31 - 1
		if partitionFirst < lengthFirst {
			minRightFirst = nums1[partitionFirst]
		}

		maxLeftSecond := -1 << 31
		if partitionSecond > 0 {
			maxLeftSecond = nums2[partitionSecond-1]
		}

		minRightSecond := 1<<31 - 1
		if partitionSecond < lengthSecond {
			minRightSecond = nums2[partitionSecond]
		}

		if maxLeftFirst <= minRightSecond && maxLeftSecond <= minRightFirst {

			if (lengthFirst+lengthSecond)%2 == 0 {

				leftMax := max(maxLeftFirst, maxLeftSecond)
				rightMin := min(minRightFirst, minRightSecond)
				return float64(leftMax+rightMin) / 2.0
			}

			return float64(max(maxLeftFirst, maxLeftSecond))
		} else if maxLeftFirst > minRightSecond {
			searchEnd = partitionFirst - 1
		} else {
			searchStart = partitionFirst + 1
		}
	}

	return 0.0
}
