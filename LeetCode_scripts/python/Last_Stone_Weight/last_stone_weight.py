from typing import List

class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        len_list = len(stones)
        if len_list == 1:
            return stones[0]

        s = sorted(stones)
        add_elem = 0 if s[-1] - s[-2] == 0 else s[-1] - s[-2]
        if len_list == 2:
            return add_elem

        s.pop(len_list -1)
        s.pop(len_list -2)
        s.append(add_elem)
        return self.lastStoneWeight(s)