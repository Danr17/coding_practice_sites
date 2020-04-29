
def solve(s):
    capitalizedName = ""
    spaces = 0
    index = 0
    for i, n in enumerate(s):
        if n is " ":
            spaces +=1
            if spaces == 1:
                capitalizedName = capitalizedName + s[index:i].capitalize()
                continue
            
        if n != " " and spaces > 0:
            capitalizedName = capitalizedName + (" " * spaces)
            spaces = 0
            index = i
        
        if i == len(s) - 1:
            capitalizedName = capitalizedName + s[index:i+1].capitalize()

    return capitalizedName

if __name__ == '__main__':
    #fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = solve(s)
    print(result)

    #fptr.write(result + '\n')

    #fptr.close()