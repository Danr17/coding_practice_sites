class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        
        def cleanup(L: str) -> str:
            result = ""
            i = 0
            for c in reversed(L):
                if c == "#":
                    i +=1
                    continue
                if i > 0:
                    i -=1
                    continue
                result = result + c
            return result

        return cleanup(S) == cleanup(T)