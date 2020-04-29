def symetric_diff(list1, list2):
    set1 = set(list1)
    set2 = set(list2)
    diff1 = set1.difference(set2)
    diff2 = set2.difference(set1)
    return diff1.union(diff2)


if __name__ == "__main__":
    m = int(input())
    m_list = list(map(int, input().split()))
    n = int(input())
    n_list = list(map(int, input().split()))
    diff = symetric_diff(m_list, n_list)
    for element in sorted(diff):
        print(element)
