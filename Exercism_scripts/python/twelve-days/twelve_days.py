
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
    
    for day in range(start_verse, end_verse + 1):
        lyric = f'On the {numbers[day - 1]} day of Christmas my true love gave to me: '

        if day == 1:
            lyric += what[day -1] + "."
        else:
            lyric += get_lyric(day)
        lyrics.append(lyric) 
        
    return lyrics

def get_lyric(day: int) -> str:
    line = ""
    while day > 0: 
        if day == 1:
            line += "and " + what[day - 1] + "."
            break
        line += what[day - 1] + ", "
        day -= 1
    return line
