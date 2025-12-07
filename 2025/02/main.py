class IDRange:
    start: int
    stop: int

def task_1(id_ranges: list[IDRange]) -> int:
    counter = 0
    for id_range in id_ranges:
        for id in range(id_range.start, id_range.stop+1):
            id_str = str(id)
            if len(id_str) % 2 == 0:
                if id_str[:int(len(id_str) / 2)] == id_str[int(len(id_str) / 2):]:
                    counter += id

    return counter

def task_2(id_ranges: list[IDRange]) -> int:
    counter = 0
    for id_range in id_ranges:
        for id in range(id_range.start, id_range.stop+1):
            id_str = str(id)
            for i in range(1, int(len(id_str) / 2)+1):
                if len(id_str) % i == 0:
                    items = []
                    for j in range(int(len(id_str) / i)):
                        items.append(id_str[j*i:j*i + i])

                    if len(set(items)) == 1:
                        counter += id
                        break

    return counter

if __name__ == "__main__":
    file = open("input", "r")
    content = file.read()
    ids = content.split(",")
    id_ranges = []
    for id in ids:
        id_range = IDRange()
        id_range.start = int(id[:id.index("-")])
        id_range.stop = int(id[id.index("-")+1:])
        id_ranges.append(id_range)

    print(f"solution task1: {task_1(id_ranges)}")
    print(f"solution task2: {task_2(id_ranges)}")
