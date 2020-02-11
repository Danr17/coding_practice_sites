
def palindrome(n):
    rev = 0
    temp = n
    while(temp > 0):
        dig=temp % 10
        rev=rev*10+dig
        temp=temp//10
    if rev == n:
        return True
    return False

def palinProducts(min_factor, max_factor):
    dictProducts = {x * y : [x, y] for y in range(min_factor, max_factor +1) for x in range(min_factor, max_factor +1)}
    palindroms = [x for x in dictProducts.keys() if palindrome(x)]
    if len(palindroms) == 0:
        return None, [], None, []
    return palindroms[0], [dictProducts[palindroms[0]]], palindroms[-1], [dictProducts[palindroms[-1]]]


def largest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("The max_factor should be gretter then min_factor")
    _, _, maxVal, dictMax = palinProducts(min_factor, max_factor)
    return maxVal, dictMax
    

def smallest(min_factor, max_factor):
    if min_factor > max_factor:
        raise ValueError("The max_factor should be gretter then min_factor")
    minVal, dictMin, _, _ = palinProducts(min_factor, max_factor)
    return minVal, dictMin

if __name__ == "__main__":
    print(largest(10, 99))
    
