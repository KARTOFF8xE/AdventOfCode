from collections import defaultdict

class Coords:
    x: int
    y: int
    
    def __init__(self, x: int, y: int):
        self.x = x
        self.y = y

    def __eq__(self, other):
        return isinstance(other, Coords) and self.x == other.x and self.y == other.y

    def __hash__(self):
        return hash((self.x, self.y))

def task_1(lines: list[str]):
    start = Coords(0,0)

    for i in range(len(lines)):
        lines[i] = list(lines[i].replace(".", " "))
        if "S" in lines[i]:
            start.x = lines[i].index("S")
            start.y = i

    counter = 0
    current_beams = {start}
    for _ in range(start.y, len(lines)-1):
        new_beams = []

        for beam in current_beams:
            if lines[beam.y+1][beam.x] == " ":
                new_beams.append(Coords(beam.x, beam.y+1))
            else:
                counter += 1
                new_beams.append(Coords(beam.x-1, beam.y+1))
                new_beams.append(Coords(beam.x+1, beam.y+1))

        current_beams = set(new_beams)

    return counter

def task_2(lines: list[str]):
    start = Coords(0,0)

    for i in range(len(lines)):
        lines[i] = list(lines[i].replace(".", " ").replace("^", "#"))
        if "S" in lines[i]:
            start.x = lines[i].index("S")
            start.y = i

    map = defaultdict(int)
    map[start] = 1

    current_beams = {start}
    for _ in range(start.y, len(lines)-1):
        new_beams = []

        for beam in current_beams:
            if lines[beam.y+1][beam.x] == " ":
                new_beams.append(Coords(beam.x, beam.y+1))
                map[Coords(beam.x, beam.y+1)] += map[beam]
            else:
                new_beams.append(Coords(beam.x-1, beam.y+1))
                map[Coords(beam.x-1, beam.y+1)] += map[beam]
                new_beams.append(Coords(beam.x+1, beam.y+1))
                map[Coords(beam.x+1, beam.y+1)] += map[beam]

        current_beams = set(new_beams)

    counter = 0
    for x in range(len(lines[0])):
        counter += map[Coords(x,len(lines)-1)]
    return counter

if __name__ == "__main__":
    file = open("input", "r")
    lines = file.read().split("\n")

    print(f"solution for task1: {task_1(lines.copy())}")
    print(f"solution for task2: {task_2(lines)}")