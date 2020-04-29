if __name__ == "__main__":
    n = int(input())
    tourists = list(map(int, input().split()))
    myset = set(tourists)
    print(((sum(myset)*n)-(sum(tourists)))//(n-1))

