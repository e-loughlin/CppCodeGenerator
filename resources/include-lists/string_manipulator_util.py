

import sys
import os
import ntpath

def readFile(filePath):
    with open(filePath, "r") as file:
        return file.read()

def writeToDisk(filePath, stringToSave):
    with open(filePath, "w+") as newFile:
        newFile.write(stringToSave)

def main():
    filePath = os.path.abspath("./qt-includes.txt")
    stuff = readFile(filePath)
    stuff = stuff.replace(" ", "\n")

    lines = stuff.split("\n")
    newLines = []

    for line in lines:
        line = line.replace("(", "").replace(")", "")
        if line[0] != "Q":
            line = line[1:]
        newLines.append(line)
    
    stuff = "\n".join(newLines)

    writeToDisk("qt-includes-new.txt", stuff)
    

if __name__ == "__main__":
    main()
    