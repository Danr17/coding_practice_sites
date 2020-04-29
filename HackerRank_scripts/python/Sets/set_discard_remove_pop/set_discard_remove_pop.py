
if __name__ == "__main__":
    n = int(input())
    s = set(map(int, input().split()))
    N = int(input())
    for _ in range(N):
        func = input()
        if func == "pop":
            s.pop()
            continue

        func, value = func.split()
        if func == "remove":
            s.remove(int(value))
        elif func == "discard":
            s.discard(int(value))
    print(sum(s))
