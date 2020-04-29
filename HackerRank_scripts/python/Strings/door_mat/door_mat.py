
if __name__ == '__main__':
    n, m = input().split()
    c = ".|."

    for i in range(1, int(n), 2):
        print((c*i).center(int(m), "-"))
    
    print(("WELCOME").center(int(m), "-"))

    for i in reversed(range(1, int(n), 2)):
        print((c*i).center(int(m), "-"))
