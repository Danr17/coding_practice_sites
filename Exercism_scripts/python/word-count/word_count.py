import string

def count_words(sentence):
    words = {}
    last_index = 0
    valid_letters = string.ascii_letters + string.digits + "'"

    for idx, char in enumerate(sentence):
        if char in valid_letters:
            if idx != len(sentence) -1:
                continue
            else:
                idx +=1
        word = sentence[last_index:idx].lower()
        
        while word.startswith("'"):
            word =  word[1:-1]      
        
        if word not in string.whitespace:
            if word not in words.keys():
                words[word] = 1
            else:
                words[word] += 1
        last_index = idx + 1
    
    return words


