import random
import string

class Robot:
    def __init__(self):
        self.name = ''.join(random.sample(string.ascii_uppercase, k=2)) + ''.join(random.sample(string.digits, k=3))
    def reset(self):
        random.seed(42)
        self.name = ''.join(random.sample(string.ascii_uppercase, k=2)) + ''.join(random.sample(string.digits, k=3))

