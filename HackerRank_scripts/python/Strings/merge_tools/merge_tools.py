def merge_the_tools(string, k):
    i = 0
    while i < len(string):
        baseChars = set()
        substring = ""
        for s in string[int(i):int(i+k)]:
            if s in baseChars:
                continue
            baseChars.add(s)
            substring = substring + s
        print(substring)
        i = i + k

if __name__ == '__main__':
    string, k = input(), int(input())
    merge_the_tools(string, k)