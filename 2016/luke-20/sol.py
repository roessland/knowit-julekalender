from itertools import permutations

prev = 0
for perm in permutations("0123456789"):
    num1 = int("".join(perm[:5]))
    num2 = int("".join(perm[5:]))
    if num1*num2 > prev:
        prev = num1*num2
print(num1, num2, prev)
