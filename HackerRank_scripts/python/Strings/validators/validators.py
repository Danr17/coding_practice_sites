if __name__ == '__main__':
    s = input()
    alphanumeric = False
    alphabetical = False
    digits = False
    lowercase = False
    uppercase = False
    for i, char in enumerate(s):
        if char.isalnum():
            alphanumeric = True
        if char.isalpha():
            alphabetical = True
        if char.isdigit():
            digits = True
        if char.islower():
            lowercase = True
        if char.isupper():
            uppercase = True
    print(alphanumeric, alphabetical, digits, lowercase, uppercase, sep='\n')