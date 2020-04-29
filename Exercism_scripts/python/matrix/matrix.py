class Matrix:
    def __init__(self, matrix_string):
        self.list_matrix = [[int(obj) for obj in line.split()] for line in matrix_string.splitlines()]

    def row(self, index):
        return self.list_matrix[index-1]

    def column(self, index):
        return [row[index-1] for row in self.list_matrix]
    
