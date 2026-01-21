import pytest
from main import akit_kumar, patto_bear


@pytest.mark.parametrize(
    "m, n, h_fences, v_fences, expected",
    [
        (4, 3, [2, 3], [2], 4),
        (6, 7, [2], [4], -1),
    ],
)
class TestMaximumSquareArea:
    def test_akit_kumar(
        self, m: int, n: int, h_fences: list[int], v_fences: list[int], expected: int
    ):
        assert akit_kumar(m, n, h_fences, v_fences) == expected

    def test_patto_bear(
        self, m: int, n: int, h_fences: list[int], v_fences: list[int], expected: int
    ):
        assert patto_bear(m, n, h_fences, v_fences) == expected
