import math

def factors(value):
    addFactors = []

    while(value%2==0):
        value=value//2
        addFactors.append(2)

    for i in range(3, int(value ** 0.5) + 1):
       while value % i == 0:
            addFactors.append(i)
            value = value // i
    
    if(value>2):
        addFactors.append(value)

    return addFactors

