def is_armstrong_number(number):
    power = len(str(number))
    li = list(str(number))
    sum = 0
    for l in li:
        t =int(l)**power
        sum += t

    return sum == number 
    
