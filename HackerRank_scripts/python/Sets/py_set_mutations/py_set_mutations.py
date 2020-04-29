if __name__ == "__main__":
    n = int(input())
    firstSet = set(map(int, input().split()))
    ops = int(input())
    for _ in range(ops):
        function, nr = input().split()
        secondSet = set(map(int, input().split()))
        if function == "intersection_update":
            firstSet.intersection_update(secondSet)
        elif function == "update":
            firstSet.update(secondSet)
        elif function == "symmetric_difference_update":
            firstSet.symmetric_difference_update(secondSet)
        elif function == "difference_update":
            firstSet.difference_update(secondSet)  
    print(sum(firstSet))