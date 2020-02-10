def saddle_points(matrix):
    thePoints = []
    for i, numList in enumerate(matrix):
        if len(numList) != len(matrix[0]):
            raise ValueError("irregular matrix")
        for j in range(len(numList)):
            if matrix[i][j] == max(numList) and matrix[i][j] == min([x[j] for x in matrix]):
                thePoints.append({"row": i+1, "column": j+1})
    return thePoints