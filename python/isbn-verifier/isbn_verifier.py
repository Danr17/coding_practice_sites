def is_valid(isbn):
    valid = ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "X"]
    isbn_repl = [x for x in isbn if x in valid]

    if len(isbn_repl) != 10 :
        return False

    index = 10
    sum = 0 
    for num in isbn_repl[:-1]:
        if num not in valid[:-1]:
            return False
        sum += int(num) * index
        index -= 1

    if isbn_repl[-1] in valid:
        if isbn_repl[-1] == "X":
            num = 10
            sum = sum + num * 1
        else:
            sum = sum + int(isbn_repl[-1]) * 1

    if int(sum) % 11 == 0:
        return True
    
    return False


