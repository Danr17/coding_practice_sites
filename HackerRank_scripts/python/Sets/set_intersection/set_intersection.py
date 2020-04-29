if __name__ == "__main__":
    first = int(input())
    firstlist = set(map(int, input().split()))
    second = int(input())
    secondlist = set(map(int, input().split()))
    print(len(firstlist.intersection(secondlist)))