from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        return sum(prices[a+1] - prices[a] for a in range(len(prices)-1) if prices[a] < prices[a+1])