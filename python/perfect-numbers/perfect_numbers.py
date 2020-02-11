def classify(number):
    if number <= 0:
        raise ValueError("number has to be > 0")
    sumFactors = sum([x for x in range(1, int(number/2) + 1) if number%x == 0])
    if sumFactors > number:
        return "abundant"
    elif sumFactors < number:
        return "deficient"
    else:
        return "perfect"


