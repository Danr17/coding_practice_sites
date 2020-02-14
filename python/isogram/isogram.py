def is_isogram(string: str) -> bool:
    duplicates = {}
    for _, s in enumerate(string):
        if s.isalpha() and s.lower() in duplicates:
            return False
        duplicates[s.lower()] = 1
    return True

