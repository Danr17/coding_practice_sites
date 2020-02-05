def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("The strands have different lengths, should be equal.")
    
    count = 0
    for i, _ in enumerate(strand_a):
        if strand_a[i] != strand_b[i]:
            count +=1
    
    return count




