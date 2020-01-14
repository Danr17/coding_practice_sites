class Matrix:
    def __init__(self, matrix_string):
        self.matrix_string = matrix_string
        self.list_matrix = [[int(obj) for obj in line.split()] for line in self.matrix_string.splitlines()]

    def row(self, index):
        return self.list_matrix[index-1]

    def column(self, index):
        i = 0
        while i < len(self.matrix_string[0]):
            return [row[index-1] for row in self.list_matrix]
            i += 1
