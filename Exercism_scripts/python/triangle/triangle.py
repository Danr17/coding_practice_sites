def equilateral(sides):
    if is_triangle(sides):
        return sides[0] == sides[1] == sides [2]
    return False

def isosceles(sides):
    if is_triangle(sides):
        return sides[0] == sides[1] or sides[1] == sides[2] or sides[0] == sides[2]
    return False

def scalene(sides):
    if is_triangle(sides) and not isosceles(sides) and not equilateral(sides):
        return True
    return False

def is_triangle(sides):
    if sides[0] <=0 or sides[1] <=0 or sides[2] <= 0 or \
    sides[0] + sides[1] < sides[2] or sides[1] + sides[2] < sides[0] or sides[0] + sides[2] < sides[1]:
        return False
    return True

