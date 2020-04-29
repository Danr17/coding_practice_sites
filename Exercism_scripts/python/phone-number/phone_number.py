class PhoneNumber:
    def __init__(self, number):
        self.number = self.clean(number)
        self.area_code = self.number[0:3]
    
    def clean(self, number):
        self.cleanNList = [n for n in number if str.isdigit(n)]
        if len(self.cleanNList) == 11 and self.cleanNList[0] == '1':
            del self.cleanNList[0]
        if len(self.cleanNList) != 10  or int(self.cleanNList[0]) < 2 or int(self.cleanNList[3]) < 2:
            raise ValueError("wrong number")
        return "".join(self.cleanNList)
        
    def pretty(self):
        return "(" + ''.join(self.number[0:3]) + ")" + " " + ''.join(self.number[3:6]) + "-" + ''.join(self.number[6:])


    
