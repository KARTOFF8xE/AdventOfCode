def task_1(batteries: list[str]) -> int:
    counter = 0

    for battery in batteries:
        battery = list(battery)
        battery = [int(item) for item in battery]
    
        first = max(battery[:len(battery)-1])
        second = max(battery[battery.index(first)+1:])
        nr = int(f"{first}{second}")
        counter += nr
  
    return counter

def task_2(batteries: list[str]) -> int:
    counter = 0

    for battery in batteries:
        battery = list(battery)
        battery = [int(item) for item in battery]
        active = []
        for i in range(12):
            nr = max(battery[:len(battery)-(12-i)+1])
            active.append(str(nr))
            battery = battery[battery.index(nr)+1:]

        counter += int("".join(active))
  
    return counter


if __name__ == "__main__":
    file = open("input", "r")
    lines = file.read().split("\n")

    print(f"solution task1: {task_1(lines)}")
    print(f"solution task2: {task_2(lines)}")