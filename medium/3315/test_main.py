import pytest
from main import min_bitwise_array


@pytest.mark.parametrize(
    "nums, expected",
    [
        ([2, 3, 5, 7], [-1, 1, 4, 3]),
        ([11, 13, 31], [9, 12, 15]),
    ],
)
def test_min_bitwise_array(nums: list[int], expected: list[int]):
    assert min_bitwise_array(nums) == expected
