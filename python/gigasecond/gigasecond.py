import datetime

def add(moment):
    newdate = moment + datetime.timedelta(seconds= 10**9)
    return newdate
