
def cycle_shift(string_one: str, string_two: str) -> int:
    if string_one == string_two or len(string_one) != len(string_two):
        return 0

    for i in range(len(string_one) - 1):
        string_one = string_one[len(string_one) - 1] + string_one
        string_one = string_one[:len(string_one) - 1]
        if string_one == string_two:
            return i + 1
    return -1


if __name__ == "__main__":
    print(cycle_shift("abc", "bca"))
