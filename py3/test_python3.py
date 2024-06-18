import os

from test_import_py import f
from test_import_py import x

# import cython
def test_print_name(name):
    print(os.getcwd())
    y = f(1.1)
    print(x)
    return y
