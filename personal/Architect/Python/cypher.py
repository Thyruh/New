alp = [None, "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
       "T", "U", "V", "W", "X", "Y", "Z"]


def encrypting(indexes, wheel_1):
    encrypted = ""
    for i in indexes:
        encrypted += wheel_1[i]
    return encrypted


def indexing(to_crypt):
    indexes = []
    for string_text in to_crypt:
        if string_text == " " or string_text == ",":
            continue
        ind = alp.index(string_text.capitalize())
        print(string_text.capitalize(), end=" ")
        print(ind)
        indexes.append(ind)
    return indexes


def mixing():
    shift = int(input("Enter the value of a shift between the 2 wheels:  "))
    shifted_list = []
    for i in range(shift, len(alp)):
        shifted_list.append(alp[i])
        print(alp[i], end="      ")
    for j in range(1, shift):
        shifted_list.append(alp[j])
        print(alp[j], end="      ")
    print("")
    return shifted_list


def main():
    wheel_1 = mixing()
    print(f"First:  {wheel_1}")

    print("Enter a string you want to cypher:  ")
    to_crypt = str(input())

    index_count = indexing(to_crypt)
    encrypted = encrypting(index_count, wheel_1)

    print(f"{encrypted} is your final result")


main()



