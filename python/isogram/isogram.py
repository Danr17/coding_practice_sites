def is_isogram(string: str) -> bool:
    duplicates = set()
    for _, s in enumerate(string.lower()):
        if s.isalpha() and s in duplicates:
            return False
        duplicates.add(s)
    return True
