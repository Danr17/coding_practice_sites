def recite(start_verse, end_verse):
    numbers = {1:"first", 2:"second", 3:"third", 4:"fourth", 5: "fifth", 6:"sixth", 7:"seventh", 8:"eighth", 9:"ninth", 10:"tenth", 11:"eleventh", 12:"twelfth"}
    what = {1: "a Partridge in a Pear Tree",
    2: "two Turtle Doves",
    3: "three French Hens",
    4: "four Calling Birds",
    5: "five Gold Rings",
    6: "six Geese-a-Laying",
    7: "seven Swans-a-Swimming",
    8: "eight Maids-a-Milking",
    9: "nine Ladies Dancing",
    10: "ten Lords-a-Leaping",
    11: "eleven Pipers Piping",
    12: "twelve Drummers Drumming"
    }

    if start_verse > end_verse or end_verse > 12:
        raise Exception("start_verse should be smaller than end_verse and end_verse should be smaller or equal to length of the song")
    lyrics = []  
    while start_verse <= end_verse:
        line = f'On the {numbers[start_verse]} day of Christmas my true love gave to me: '
        i = start_verse
        if i == 1 :
            line += what[i] + "."
        else:
            while i > 0:
                if i == 1:
                    line += "and " + what[i] + "."
                    i -= 1
                else:
                    line +=  what[i] + ", "
                    i -= 1
        lyrics.append(line)
        start_verse += 1
    return lyrics


