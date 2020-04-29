def minion_game(string):
    vw = 'aeiou'.upper()
    strl = len(string)
    kevin = sum(strl-i for i in range(strl) if string[i] in vw)
    stuart = strl*(strl + 1)/2 - kevin

    if kevin == stuart:
        print('Draw')
    elif kevin > stuart:
        print(f'Kevin {kevin:.0f}')
    else:
        print(f'Stuart {stuart:.0f}')

if __name__ == '__main__':
    s = input()
    minion_game(s)