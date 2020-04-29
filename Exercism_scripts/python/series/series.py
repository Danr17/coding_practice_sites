def slices(series, length):
    if series == "" or length <= 0 or len(series) < length:
        raise ValueError("The series, length has to be > 0 and the number of elements from series > length")
    return [series[x:x+length] for x in range(0, len(series)) if len(series)-x >= length]

