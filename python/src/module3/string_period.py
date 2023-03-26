
def z_function(string: str):
    result = [0] * len(string)
    l, r = 0, 0
    for i in range(1, len(string)):
        result[i] = max(0, min(result[i - l], r - i))
        while i + result[i] < len(string) and string[result[i]] == string[i + result[i]]:
            result[i] += 1
        if i + result[i] > r:
            l, r = i, i + result[i]
    return result


def string_period(string: str):
    z = z_function(string)
    for i in range(len(z)):
        if i + z[i] == len(string) and len(string) % i == 0:
            return len(string) // i
    return -1


if __name__ == "__main__":
    print(z_function("abcabcabc"))
