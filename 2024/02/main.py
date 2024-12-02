import sys


def read_file():
    f = open("./test_input.txt", "r")
    return f.read()


def part1():
    count = 0
    for line in read_file().split("\n"):
        vals = [int(i) for i in line.split()]
        deltas = deltasFor(vals)

        if isSafe(deltas):
            count += 1

    print(count)


def part2():
    count = 0
    for line in read_file().split("\n"):
        vals = [int(i) for i in line.split()]
        deltas = deltasFor(vals)

        if isSafe(deltas):
            count += 1
            continue

        for i in range(len(vals)):
            deltas = deltasFor(vals[:i] + vals[i+1:])
            if isSafe(deltas):
                count += 1
                break

    print(count)


def deltasFor(vals):
    return [vals[i] - vals[i-1] if i > 0 else None
            for i, _ in enumerate(vals)]


def isSafe(deltas):
    posCount, negCount = 0, 0
    for d in deltas[1:]:
        if d > 0:
            posCount += 1
        if d < 0:
            negCount += 1

        if abs(d) < 1 or abs(d) > 3:
            return False

    return (posCount > 0) ^ (negCount > 0)


if __name__ == "__main__":
    part = sys.argv[1]
    if part == "1":
        part1()
    elif part == "2":
        part2()
    else:
        raise Exception("Must be part 1 or 2")
