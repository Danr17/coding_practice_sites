def is_isogram(string: str) -> bool:
    duplicates = set()
    for _, s in enumerate(string):
        s_lower = s.lower()
        if s.isalpha() and s_lower in duplicates:
            return False
        duplicates.add(s_lower)
    return True

