#!/usr/bin/python3
import os

from formula import formula

n1 = os.environ.get("RIT_NUMBER_1")
n2 = os.environ.get("RIT_NUMBER_2")
divider = os.environ.get("RIT_DIVIDER")
multiselect = os.environ.get("RIT_INPUT_MULTISELECT")
formula.Run(n1, n2, divider, multiselect)
