import pytest
from main import max_side_length


@pytest.mark.parametrize(
    "mat, threshold, expected",
    [
        (
            [[1, 1, 3, 2, 4, 3, 2], [1, 1, 3, 2, 4, 3, 2], [1, 1, 3, 2, 4, 3, 2]],
            4,
            2,
        ),
        (
            [[1, 1, 1, 1], [1, 0, 0, 0], [1, 0, 0, 0], [1, 0, 0, 0]],
            6,
            3,
        ),
    ],
)
def test_max_side_length(mat: list[list[int]], threshold: int, expected: int):
    assert max_side_length(mat, threshold) == expected
