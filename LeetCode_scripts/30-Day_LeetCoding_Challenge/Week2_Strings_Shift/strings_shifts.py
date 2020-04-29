from typing import List

class Solution:
    def stringShift(self, s: str, shift: List[List[int]]) -> str:
        shiftResult = {"left": 0, "right": 0}
        for n  in shift:
            if n[0] == 0:
                shiftResult["left"] += n[1]
            else:
                shiftResult["right"] += n[1]
        times = shiftResult["left"] - shiftResult["right"]
        if times > 0:
            while times > 0:
                s = s[1:] + s[0]
                times -= 1
            return s
        else:
            while times < 0:
                s = s[-1] + s[:-1]
                times += 1
            return s