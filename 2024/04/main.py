import sys


def read_file():
    f = open("./test_input.txt", "r")
    return f.read()


def part1():
    input = read_file()
    grid = []
    for line in input.split("\n"):
        stripped = line.strip()
        grid.append(list(stripped))

    directions = [
        move_right, move_down, move_left, move_up,
        move_up_and_right, move_down_and_right,
        move_down_and_left, move_up_and_left,
    ]
    count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] != 'X':
                continue

            for d in directions:
                x, y = d(grid, i, j)

                if x == -1 or y == -1 or grid[x][y] != 'M':
                    continue

                x, y = d(grid, x, y)
                if x == -1 or y == -1 or grid[x][y] != 'A':
                    continue

                x, y = d(grid, x, y)
                if x == -1 or y == -1 or grid[x][y] != 'S':
                    continue

                count += 1

    print(count)


def part2():
    input = read_file()
    grid = []
    for line in input.split("\n"):
        stripped = line.strip()
        grid.append(list(stripped))

    directions = [
        move_up_and_right, move_down_and_right,
        move_down_and_left, move_up_and_left,
    ]
    count = 0
    for i in range(len(grid)):
        for j in range(len(grid[i])):
            if grid[i][j] != 'A':
                continue

            adj = {}
            for idx, d in enumerate(directions):
                x, y = d(grid, i, j)
                if x == -1 or y == -1:
                    continue

                val = grid[x][y]
                if val not in adj:
                    adj[val] = []
                adj[val].append(idx)

            if 'S' not in adj or 'M' not in adj:
                continue

            if len(adj['S']) != 2 or len(adj['M']) != 2:
                continue

            if (adj['S'][0] % 2) == (adj['S'][1] % 2) or (adj['M'][0] % 2) == (adj['M'][1] % 2):
                continue

            count += 1

    print(count)


def move_right(grid, i, j):
    x = [i + 1, j]
    if i + 1 >= len(grid):
        x[0] = -1

    return x


def move_down(grid, i, j):
    x = [i, j + 1]
    if j + 1 >= len(grid[i]):
        x[1] = -1

    return x


def move_left(grid, i, j):
    x = [i - 1, j]
    if i - 1 < 0:
        x[0] = -1

    return x


def move_up(grid, i, j):
    x = [i, j - 1]
    if j - 1 < 0:
        x[1] = -1

    return x


def move_up_and_right(grid, i, j):
    x = [i + 1, j - 1]
    if i + 1 >= len(grid):
        x[0] = -1

    if j - 1 < 0:
        x[1] = -1

    return x


def move_down_and_right(grid, i, j):
    x = [i + 1, j + 1]
    if i + 1 >= len(grid):
        x[0] = -1

    if j + 1 >= len(grid[i]):
        x[1] = -1

    return x


def move_down_and_left(grid, i, j):
    x = [i - 1, j + 1]
    if i - 1 < 0:
        x[0] = -1

    if j + 1 >= len(grid[0]):
        x[1] = -1

    return x


def move_up_and_left(grid, i, j):
    x = [i - 1, j - 1]
    if i - 1 < 0:
        x[0] = -1

    if j - 1 < 0:
        x[1] = -1

    return x


# class Grid():
#     def __init__(self, grid):
#         self.grid = grid


if __name__ == "__main__":
    part = sys.argv[1]
    if part == "1":
        part1()
    elif part == "2":
        part2()
    else:
        raise Exception("Must be part 1 or 2")
