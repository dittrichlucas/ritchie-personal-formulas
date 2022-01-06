#!/usr/bin/python3

from formula.output import readOutput, removeFile, runFormula, writeOutput


def Run(input1, input2, input3, inputx):
    cmd = f'rit math sum --number_1={input1} --number_2={input2}'
    runFormula(cmd)

    data = readOutput()
    divisor = int(data['output']['sum']) / int(input3)

    print(f'Result: {str(divisor)}')

    removeFile()

    print(inputx)
