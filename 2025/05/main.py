class Range:
    start: int
    stop: int

def task_1(ranges: list[Range], ids: int):
    counter = 0

    for id in ids:
        for r in ranges:
            if r.start <= id <= r.stop:
                counter += 1
                break

    return counter

def return_start_val(e: Range):
    return e.start

def task_2(ranges: list[Range]):
    counter = 0

    ranges.sort(key=return_start_val)

    shortened = True
    while shortened:
        shortened = False
        for i in range(len(ranges)-1):
            if ranges[i].stop >= ranges[i+1].start:
                if ranges[i+1].stop > ranges[i].stop:
                    ranges[i].stop = ranges[i+1].stop
                ranges.pop(i+1)
                shortened = True
                break
            
    for r in ranges:
        counter += r.stop - r.start + 1
            
    return counter


if __name__ == "__main__":
    file = open("input", "r")
    lines = file.read().split("\n")
    ranges = []
    ids = []

    for line in lines:
        if line == "": continue
        if "-" in line:
            r = Range()
            r.start = int(line[:line.index("-")])
            r.stop = int(line[line.index("-")+1:])
            ranges.append(r)
        else:
            ids.append(int(line))

    print(f"solution for task1: {task_1(ranges, ids)}")
    print(f"solution for task2: {task_2(ranges)}")