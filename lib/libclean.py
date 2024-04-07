class LibClean:
    _collection = set()

    def collector(self, val):
        self._collection.add(val)

    def removeduplicates(self, collection):
        cleanedcollection = set()

        for x in collection:
            for y in self.getcollection():
                if x != y:
                    cleanedcollection.add(x)

        return cleanedcollection

    def getcollection(self):
        return list(self._collection)
