def valid_braces(string):
    for i in range(len(string)):
        string = string.replace("()", "").replace("{}", "").replace("[]", "")
    if string == "":
        return True
    else:
        return False

    
print("1 :", valid_braces("()"))
print("2 :", valid_braces("[(])"))
print("3 :", valid_braces("(){}[]"))
print("4 :", valid_braces("([{}])"))
print("5 :", valid_braces("([)]"))
print("6 :", valid_braces(")"))
print("7 :", valid_braces("("))
print("8 :", valid_braces("[({})](]"))


