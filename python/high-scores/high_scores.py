def latest(scores):
    return scores[-1]


def personal_best(scores):
    return max(scores)


def personal_top_three(scores):
    top_three = []
    for _ in range(min(3, len(scores))):
        top_three.append(max(scores))
        scores.remove(max(scores))
    return top_three


