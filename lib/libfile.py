from platform import system
from pathlib import PureWindowsPath, PurePosixPath
from lib.libvar import LibVar
from lib.libstring import LibString
from lib.libclean import LibClean
from os import path, walk


class LibFile:
    _option = ""
    _directory = ""
    _foldercontents = []

    _libvar = LibVar()
    _libstring = LibString()
    _libclean = LibClean()

    def __init__(self, directory, option):
        self._option = option
        self._directory = directory

    def _convertpath(self):
        if system() == "Linux" or system() == "Darwin":
            return PurePosixPath(*PureWindowsPath(self._directory).parts)
        else:
            return PureWindowsPath(self._directory)

    def findfiles(self):
        for root, dir, files in walk(self._convertpath()):
            for file in files:
                if file.endswith('.osc'):
                    self._foldercontents.append(path.join(root, file))

    def _findduplicates(self, lst):
        uniques = list(set(lst))

        for item in uniques:
            self._libclean.collector(item)
        return uniques

    def processfile(self, file):
        if system() == "Linux" or system() == "Darwin":
            with open(file, "r", encoding="cp1252") as f:
                filecontent = f.read()
        else:
            with open(file, "r") as f:
                filecontent = f.read()

        if self._option == "varlist":
            return self._libclean.removeduplicates(list(self._findduplicates(self._libvar.extractvarname(filecontent))))
        elif self._option == "stringvarlist":
            return list(self._findduplicates(self._libstring.extractstringname(filecontent)))

    def _modifyfilename(self, file):
        if self._option == "varlist":
            return file.replace(".osc", "_varlist.txt")
        elif self._option == "stringvarlist":
            return file.replace(".osc", "_stringvarlist.txt")

    def writeoutput(self, file):
        i = 0
        with open(self._modifyfilename(file), 'w') as o:
            for content in self.processfile(file):
                o.write(content)
                if (i + 1) < len(self.processfile(file)):
                    o.write("\n")
                i += 1
        o.close()

    def getdirectorycontents(self):
        return self._foldercontents
