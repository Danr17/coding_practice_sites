def convert(number):
    x = ""
    if number % 3 == 0:
        x += "Pling"
    if number % 5 == 0:
        x += "Plang"
    if number % 7 == 0:
        x += "Plong"
    
    return x if x is not "" else str(number)
