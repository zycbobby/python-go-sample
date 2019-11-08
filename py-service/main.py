# coding=utf-8

import sys
import os

abspath = os.path.abspath(__file__)
dname = os.path.dirname(abspath)
os.chdir(dname)

sys.path.append("./third_party")
sys.path.append(".")

import __future__
from handler.car import Car
import requests

if __name__ == '__main__':
    c = Car("benz")
    print(c.brand)
    print(123)
    print(requests.get("http://www.baidu.com"))
    