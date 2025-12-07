def task_1(lines: list[str]) -> int:
    ptr = 50
    counter = 0

    for line in lines:
        inc = int(line[1:])
        if line[0] == "L":
            ptr -= inc
        else:
            ptr += inc
        ptr = ptr % 100
        
        if ptr == 0:
            counter += 1
        
    return counter

def task_2(lines: list[str]) -> int:
    ptr = 50
    counter = 0

    for line in lines:
        inc = int(line[1:])
        counter += int(inc / 100)
        inc %= 100
        if line[0] == "L":
            ptr -= inc
            if ptr <= 0 and ptr != -inc:
                counter += 1
            ptr %= 100
        else:
            ptr += inc
            if ptr >= 100:
                counter += 1
            ptr %= 100

    return counter

if __name__ == "__main__":
    file = open("input", "r")
    content = file.read()
    lines = content.split("\n")

    print(f"solution task1: {task_1(lines)}")
    print(f"solution task2: {task_2(lines)}")