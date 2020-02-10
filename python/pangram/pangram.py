import string

def is_pangram(sentence):
    dictWord = dict.fromkeys(string.ascii_lowercase, 0)
    for _, w in enumerate(sentence):
        dictWord[w.lower()] = 1
    for w in dictWord.values():
        if w == 0:
            return False 
    return True




