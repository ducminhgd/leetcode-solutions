# 2975. Maximum Square Area by Removing Fences From a Field
# https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/description

MOD = 10**9 + 7


def akit_kumar(m: int, n: int, h_fences: list[int], v_fences: list[int]) -> int:
    """
    Solution from AkitKumar on leetcode
    https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/solutions/7498260/find-biggest-square-from-fence-gaps-begi-8o9m/
    """

    def prep(cuts: list[int], limit: int) -> list[int]:
        return sorted([1] + cuts + [limit])

    h = prep(h_fences, m)
    v = prep(v_fences, n)

    gap_set = set()
    for i in range(len(h)):
        for j in range(i + 1, len(h)):
            gap_set.add(h[j] - h[i])

    best = 0
    for i in range(len(v)):
        for j in range(i + 1, len(v)):
            d = v[j] - v[i]
            if d > best and d in gap_set:
                best = d

    if best == 0:
        return -1
    return (best * best) % MOD


def patto_bear(m: int, n: int, h_fences: list[int], v_fences: list[int]) -> int:
    """
    Solution from PattoBear on leetcode
    https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/solutions/7498442/golang-good-ass-hints-by-pattobears-mq85/
    """
    h_fences = h_fences + [1, m]
    v_fences = v_fences + [1, n]

    diff_set = set()
    for i in range(len(h_fences) - 1):
        for j in range(i + 1, len(h_fences)):
            diff_set.add(abs(h_fences[i] - h_fences[j]))

    out = -1
    for i in range(len(v_fences) - 1):
        for j in range(i + 1, len(v_fences)):
            side = abs(v_fences[i] - v_fences[j])
            if side in diff_set:
                out = max(out, side * side)

    return out % MOD if out != -1 else -1
