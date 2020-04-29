def swap_case(s):
    swapString = ""
    for char in s:
        if char.isupper():
            swapString = swapString + char.lower()
        else:
            swapString = swapString + char.upper()
    return swapString

if __name__ == '__main__':
    s = input()
    result = swap_case(s)
    print(result)