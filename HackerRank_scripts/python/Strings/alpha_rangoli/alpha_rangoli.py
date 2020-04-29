def print_rangoli(size):
    ascii_letters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
    letters = list(ascii_letters[:size])
    
    middle = "a"
    for letter in letters[1:]:
        middle = letter + "-" + middle + "-" + letter

    cone = [] 
    for i in range(size-1):
        top = letters[-i-1]
        for j in reversed(range(i)):
            top = letters[-j-1] + "-" + top + "-" + letters[-j-1]
        cone.append(top.center(len(middle), "-"))
    
    [print(i) for i in cone]
    print(middle)
    [print(i) for i in reversed(cone)]

if __name__ == '__main__':
    n = int(input())
    print_rangoli(n)