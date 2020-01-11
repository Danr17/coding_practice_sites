def latest(scores):
    return scores[-1]


def personal_best(scores):
    return max(scores)

def personal_top_three(scores):
    index = min(3, len(scores))
    scores.sort(reverse=True)
    return scores[:index]




