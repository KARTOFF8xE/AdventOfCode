from enum import Enum

class Operator(Enum):
    add = "+"
    mul = "*"

class Operation:
    operator: Operator
    nums: list[int]

def task_1(lines: list[str]):
    for i in range(len(lines)):
        lines[i] = lines[i].split(" ")
        lines[i] = list(filter(None, lines[i]))

    ops = []
    for i in range(len(lines[0])):
        op = Operation()
        op.nums = []
        op.operator = Operator.add if lines[len(lines)-1][i] == "+" else Operator.mul
        for j in range(len(lines)-1):
            op.nums.append(int(lines[j][i]))
        ops.append(op)

    counter = 0
    for op in ops:
        if op.operator == Operator.add:
            counter += sum(op.nums)
        else:
            mul = 1
            for num in op.nums:
                mul *= num
            counter += mul

    return counter


def task_2(lines: list[str]):
    for i in range(len(lines)):
        lines[i] = list(lines[i])
    
    ops = []
    nums = []
    for i in range(len(lines[0])-1, -1, -1):
        num = [line[i] for line in lines[:-1]]
        if "".join(num).replace(" ", "") == "":
            continue

        num = int("".join(num))
        nums.append(num)

        if lines[len(lines)-1][i] != " ":
            op = Operation()
            op.nums = nums
            op.operator = Operator.add if lines[len(lines)-1][i] == "+" else Operator.mul
            ops.append(op)
            nums = []

    counter = 0
    for op in ops:
        if op.operator == Operator.add:
            counter += sum(op.nums)
        else:
            mul = 1
            for num in op.nums:
                mul *= num
            counter += mul

    return counter

if __name__ == "__main__":
    file = open("input", "r")
    lines = file.read().split("\n")

    print(f"solution for task1: {task_1(lines.copy())}")
    print(f"solution for task2: {task_2(lines)}")