class PhoneNumber:
    def __init__(self, number):
        self.cleanNList = [n for n in number if str.isdigit(n)]
        if len(self.cleanNList) != 10  or int(self.cleanNList[0]) < 2 or int(self.cleanNList[4]) < 2:
            raise ValueError("wrong number")
        self.number = "".join(self.cleanNList)
