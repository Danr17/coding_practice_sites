
if __name__ == "__main__":
    samples = int(input())
    stamps = set()
    for _ in range(samples):
        stamp = input()
        stamps.add(stamp)
    print(len(stamps))
    

