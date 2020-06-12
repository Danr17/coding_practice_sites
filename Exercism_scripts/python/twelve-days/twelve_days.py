
numbers = ["first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"]
what = ["a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming"
]

def recite(start_verse: int, end_verse: int) -> list:
    if start_verse > end_verse or end_verse > 12:
        raise Exception("start_verse should be smaller than end_verse and end_verse should be smaller or equal to length of the song")
    
    lyrics = []
    
    for i in range(start_verse, end_verse + 1):
        line = f'On the {numbers[i - 1]} day of Christmas my true love gave to me: '

        if i == 1:
            line += what[i -1] + "."
        else:
            line += get_lyric(i)
        print(line)
        lyrics.append(line) 
        
    return lyrics

def get_lyric(line: int) -> str:
    lyric = ""
    while line > 0: 
        if line == 1:
            lyric += "and " + what[line - 1] + "."
            break
        lyric += what[line - 1] + ", "
        line -= 1
    return lyric
