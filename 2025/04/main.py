def task_1(lines: list[str]) -> int:
    counter = 0

    for y, line in enumerate(lines):
        for x, item in enumerate(line):
            if item != "@":
                continue

            nr_of_rolls_surround_item = 0
            for x_i in range(x-1, x+2):
                for y_i in range(y-1, y+2):
                    if x_i < 0 or y_i < 0: continue
                    if x_i >= len(line) or y_i >= len(lines): continue
                    if x_i == x and y_i == y: continue
                    if lines[y_i][x_i] == "@": nr_of_rolls_surround_item += 1
            
            if nr_of_rolls_surround_item < 4:
                counter += 1

    return counter

import time
def task_2(lines: list[str]) -> int:
    counter = 0
    old_counter = None

    while True:
        for y, line in enumerate(lines):
            for x, item in enumerate(line):
                if item != "@":
                    continue

                nr_of_rolls_surround_item = 0
                for x_i in range(x-1, x+2):
                    for y_i in range(y-1, y+2):
                        if x_i < 0 or y_i < 0: continue
                        if x_i >= len(line) or y_i >= len(lines): continue
                        if x_i == x and y_i == y: continue
                        if lines[y_i][x_i] == "@" or lines[y_i][x_i] == "x":
                            nr_of_rolls_surround_item += 1
                
                if nr_of_rolls_surround_item < 4:
                    lines[y][x] = "x"
                    counter += 1

        for y in range(len(lines)):
            for x in range(len(lines[y])):
                if lines[y][x] == "x":
                    lines[y][x] = "."

        if counter == old_counter:
            break

        old_counter = counter

    return counter

if __name__ == "__main__":
    file = open("input", "r")
    lines = file.read().split("\n")
    map = []
    for line in lines:
        map.append(list(line))

    print(f"solution for task1: {task_1(map)}")
    print(f"solution for task2: {task_2(map)}")