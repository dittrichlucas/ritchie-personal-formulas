#!/usr/bin/python3
import json
import os
import subprocess

homedir = os.path.expanduser('~')
outputFile = homedir + '/.rit/output.json'

def runFormula(cmd):
    cmd = cmd.split(' ')
    subprocess.call(cmd)

def writeOutput(data):
    with open(outputFile, 'w') as f:
        json.dump(data, f)

def readOutput():
    with open(outputFile) as f:
        data = json.load(f)

    return data

def removeFile():
    os.remove(outputFile)
