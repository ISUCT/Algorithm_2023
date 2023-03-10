P = 359
X = 13


def get_hash(string: str) -> int:
    hashS = 0
    for i in range(len(string)):
        hashS = (X * hashS + ord(string[i])) % P
    return hashS


def get_sliding_hash(old_hash: int, old_char: str, new_char: str, mult: int) -> int:
    return ((old_hash - (ord(old_char) * mult)) * X + ord(new_char)) % P


def get_collision_index(string: str, substrng: str) -> list:
    len_sub = len(substrng)
    len_str = len(string)
    result = []
    mult = X**(len_sub - 1)
    hashSub = get_hash(substrng)
    hashS = get_hash(string[:len_sub])
    for i in range(len_str - len_sub + 1):
        if hashS == hashSub and substrng == string[i:len_sub + i]:
            result.append(i)
        if i < len_str - len_sub:
            hashS = get_sliding_hash(hashS, string[i], string[i + len_sub], mult)
    return result


if __name__ == "__main__":
    file = open('input.txt')
    string = file.readline()
    substring = file.readline()
    file.close()
    get_collision_index(string, substring)
