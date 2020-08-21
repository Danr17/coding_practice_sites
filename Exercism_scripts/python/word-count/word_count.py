import re
from collections import Counter

def count_words(sentence):
    sentence = re.sub(r'[^a-z0-9 \']', ' ', sentence.lower())
    return Counter([word.strip("'") for word in sentence.split()])


