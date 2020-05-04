class Solution:
    def isHappy(self, n: int) -> bool:
        numbers = {}
        result = 0
        happy = False
        while True :
            print(n)
            result = self.nextSum(n)
            if result == 1:
                happy = True
                break
            if result in numbers.keys():
                break
            numbers[result] = True
            n = result
        return happy 
    
    @staticmethod
    def nextSum(n: int) -> int:
        s = []
        while n > 0 :
            s.append(n % 10)
            n = n // 10
        return sum([i**2 for i in s])
