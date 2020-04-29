def count_substring(string, sub_string):
    number = 0
    for i in range(0, len(string)):
        if len(string[i:]) >= len(sub_string) and string[i:i+len(sub_string)] == sub_string:
            number += 1
    return number

if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()
    
    count = count_substring(string, sub_string)
    print(count)