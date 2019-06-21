"""
Included for example/comparison
"""

def do():
    return 1

def dodo(f):
    return f()

def dododo(f, g):
    return f(g)

if __name__ == '__main__':
    print(do())
    print(dodo(do))
    print(dododo(dodo, do))
