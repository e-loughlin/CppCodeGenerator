#!/usr/bin/env python
#
# Copyright 2008 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Author: Evan Loughlin, via copying gmock_gen.py and wrapping additionally functionality around it.
# All rights retained by original author of gmock_gen.py.

"""CppCodeGenerator wrapper of gmock_gen.py to modify the interface prior to having gmock_gen.py use it"""

import os
import sys
import gmock_gen


def removeSignalsFromInterface(interfaceContents):
    updatedContents = []
    withinSignalBlock = False

    for line in interfaceContents:
        if "signals:" in line:
            withinSignalBlock = True
        if ("private:" in line) or ("public:" in line) or ("};" in line):
            withinSignalBlock = False
        if not withinSignalBlock:
            updatedContents.append(line)
    return updatedContents

def main(argv=sys.argv):
    if(len(argv) < 2):
        sys.stderr("Path to an existing interface is required.")
        sys.exit(0)

    filePath = argv[1]
    f = open(filePath)
    interfaceContents = f.readlines()
    f.close()

    interfaceContents = removeSignalsFromInterface(interfaceContents)

    currentScriptPath = os.path.realpath(__file__)
    tempFilePath = os.path.join(os.path.dirname(currentScriptPath), "temporaryFile.txt")

    f = open(tempFilePath, "w")
    f.writelines(interfaceContents)
    f.close()

    sys.argv[1] = tempFilePath

    # Add the directory of this script to the path so we can import gmock_class.
    sys.path.append(os.path.dirname(__file__))

    from cpp import gmock_class
    # Fix the docstring in case they require the usage.
    gmock_class.__doc__ = gmock_class.__doc__.replace('gmock_class.py', __file__)
    gmock_class.main()

if __name__ == '__main__':
    main()