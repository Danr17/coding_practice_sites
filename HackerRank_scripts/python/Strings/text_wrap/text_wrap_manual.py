import textwrap

def wrap(string, max_width):
    index = 0
    char_add = ""
    for char in string:
        if index == max_width:
            char_add = char_add + "\n"
            index = 0
        char_add = char_add + char
        index += 1
    return char_add

if __name__ == '__main__':
    string, max_width = input(), int(input())
    result = wrap(string, max_width)
    print(result)