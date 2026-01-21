# 1292. Maximum Side Length of a Square with Sum Less than or Equal to Threshold
# https://leetcode.com/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/description


def max_side_length(mat: list[list[int]], threshold: int) -> int:
    m, n = len(mat), len(mat[0])

    # Build prefix sum matrix
    P = [[0] * (n + 1) for _ in range(m + 1)]

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            P[i][j] = P[i - 1][j] + P[i][j - 1] - P[i - 1][j - 1] + mat[i - 1][j - 1]

    def get_rect(x1: int, y1: int, x2: int, y2: int) -> int:
        return P[x2][y2] - P[x1 - 1][y2] - P[x2][y1 - 1] + P[x1 - 1][y1 - 1]

    r = min(m, n)
    ans = 0

    for i in range(1, m + 1):
        for j in range(1, n + 1):
            for c in range(ans + 1, r + 1):
                if i + c - 1 <= m and j + c - 1 <= n and get_rect(i, j, i + c - 1, j + c - 1) <= threshold:
                    ans = c
                else:
                    break

    return ans
