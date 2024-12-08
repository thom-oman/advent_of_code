import sys


def read_file():
    f = open("./inputs.txt", "r")
    return f.read()


def part1():
    input = read_file()
    antinodes = []
    antennas = {}
    for i, line in enumerate(input.strip().split("\n")):
        for j, chr in enumerate(line):
            if chr != ".":
                if chr not in antennas:
                    antennas[chr] = []
                antennas[chr].append([i,j])

        antinodes.append([False for _ in range(len(line))])

    for freq in antennas:
        nodes = antennas[freq]
        if len(nodes) <= 1:
            continue

        combos = []
        for i, n in enumerate(nodes):
            for other in nodes[i+1:]:
                combos.append([n, other])

        for c in combos:
            a, b = c[0], c[1]
            ax, ay, bx, by = a[0], a[1], b[0], b[1]
            dx, dy = bx - ax, by - ay
            an1x, an1y, an2x, an2y = ax - dx, ay - dy, bx + dx, by + dy
            if in_grid(antinodes, an1x, an1y):
                antinodes[an1x][an1y] = True
            if in_grid(antinodes, an2x, an2y):
                antinodes[an2x][an2y] = True

    count = 0
    for row in antinodes:
        for v in row:
            if v:
                count += 1
    print(count)


def part2():
    input = read_file()
    antinodes = []
    antennas = {}

    for i, line in enumerate(input.strip().split("\n")):
        for j, chr in enumerate(line):
            if chr != ".":
                if chr not in antennas:
                    antennas[chr] = []
                antennas[chr].append([i, j])
        antinodes.append([False for _ in range(len(line))])

    for freq in antennas:
        nodes = antennas[freq]
        if len(nodes) <= 1:
            continue

        combos = []
        for i, n in enumerate(nodes):
            for other in nodes[i+1:]:
                combos.append([n, other])

        for c in combos:
            a, b = c[0], c[1]
            ax, ay, bx, by = a[0], a[1], b[0], b[1]
            antinodes[ax][ay] = True
            antinodes[bx][by] = True
            dx, dy = bx - ax, by - ay
            nx, ny = ax - dx, ay - dy
            while in_grid(antinodes, nx, ny):
                antinodes[nx][ny] = True
                nx -= dx
                ny -= dy

            nx, ny = bx + dx, by + dy
            while in_grid(antinodes, nx, ny):
                antinodes[nx][ny] = True
                nx += dx
                ny += dy

    count = 0
    for row in antinodes:
        for v in row:
            if v:
                count += 1
    print(count)


def in_grid(grid, x, y):
    if x < 0 or y < 0:
        return False

    if x >= len(grid) or y >= len(grid[0]):
        return False

    return True


if __name__ == "__main__":
    part = sys.argv[1]
    if part == "1":
        part1()
    elif part == "2":
        part2()
    else:
        raise Exception("Must be part 1 or 2")
