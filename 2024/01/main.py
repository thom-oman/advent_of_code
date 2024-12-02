import sys

def read_file():
    f = open("./test_input.txt", "r")
    return f.read()

def part1():
    input = read_file()
    left, right = [], []
    for l in input.split("\n"):
        nums = l.split()
        l, r = nums[0], nums[1]
        left.append(int(l))
        right.append(int(r))

    left.sort()
    right.sort()

    sum = 0
    for l, r in zip(left, right):
        sum += abs(l - r)

    print(sum)

def part2():
    input = read_file()
    left, right = [], []
    for l in input.split("\n"):
        nums = l.split()
        l, r = nums[0], nums[1]
        left.append(int(l))
        right.append(int(r))

    tallys = {}
    for i in right:
        if i not in tallys:
            tallys[i] = 0

        tallys[i] += 1

    sum = 0
    for i in left:
        mul = 0
        if i in tallys:
            mul = tallys[i]
        sum += i * mul

    print(sum)


if __name__ == "__main__":
    part = sys.argv[1]
    if part == "1":
        part1()
    elif part == "2":
        part2()
    else:
        raise(Exception("Must be part 1 or 2"))
