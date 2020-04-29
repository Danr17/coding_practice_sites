if __name__ == "__main__":
    set_A = set(map(int, input().split()))
    nr_sets = int(input())
    strict_subset = True
    for i in range(nr_sets):
        set_B = set(map(int, input().split()))
        strict_subset = strict_subset and set_A.issuperset(set_B)
    print(strict_subset)


