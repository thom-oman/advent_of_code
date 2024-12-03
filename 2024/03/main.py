import sys
import re


def read_file():
    f = open("./test_input.txt", "r")
    return f.read()


def part1():
    input = read_file()
    i = 0
    sum = 0
    while i < len(input):
        if input[i] == 'm':
            to_search = input[i:i+12]
            match = re.search(
                    r"^mul\((?P<first>\d{1,3}),(?P<second>\d{1,3})\)",
                    to_search
            )

            if match:
                sum += int(match["first"]) * int(match["second"])
        i += 1
    print(sum)


def part2():
    input = read_file()
    i = 0
    sum = 0
    enabled = True
    while i < len(input):
        if input[i] == 'm' and enabled:
            to_search = input[i:i+12]
            match = re.search(
                    r"^mul\((?P<first>\d{1,3}),(?P<second>\d{1,3})\)",
                    to_search
            )

            if match:
                sum += int(match["first"]) * int(match["second"])

        if input[i:i+2] == 'do':
            if input[i:i+4] == 'do()':
                enabled = True
            if input[i:i+7] == "don't()":
                enabled = False

        i += 1
    print(sum)


if __name__ == "__main__":
    part = sys.argv[1]
    if part == "1":
        part1()
    elif part == "2":
        part2()
    else:
        raise Exception("Must be part 1 or 2")
