if __name__ == "__main__":
    test_cases = int(input())
    for i in range(test_cases):
        nr_A = int(input())
        set_A = set(map(int, input().split()))
        nr_B = int(input())
        set_B = set(map(int, input().split()))
        print(set_A.issubset(set_B))


