def square(number):
    if number <= 0 or number > 64:
        raise ValueError("the number should be in (0,64) range")
    return 2 ** (number-1)


def total():
    return sum([2 ** x for x in range(0, 64)])
