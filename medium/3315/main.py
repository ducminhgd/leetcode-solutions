# 3315. Construct the Minimum Bitwise Array II
# https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/description


def min_bitwise_array(nums: list[int]) -> list[int]:
    result = []
    for p in nums:
        removable = ((p + 1) & ~p) >> 1
        if removable == 0:
            result.append(-1)
        else:
            result.append(p ^ removable)
    return result
