from functools import lru_cache


@lru_cache(50)
def F(n):
    if n == 0 or n == 1:
        return n
    return F(n-1)+F(n-2)


print(F(50))
