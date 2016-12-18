#!/usr/bin/python
import redis
import sys
import mmh3

# class representing an object that writes html to cache in redis
class htmlImporter:
  def __init__(self, filePath):
    self.filePath = filePath
    with open(filePath) as fh:
      htmlList = fh.readlines()
    self.htmlString = "".join(htmlList)
    self.redisClient = redis.StrictRedis(host='localhost', port=6379, db=0)

  #method to write a new url and its html to redis
  def writeFile(self, url):
    self.urlHash = str(mmh3.hash64(url)[0])
    print "writing contents of {0} as cached html for {1}".format(
          self.filePath,
          url)
    self.redisClient.set(self.urlHash, self.htmlString)


# usage: python testCache.py url.tld/path file.html
if __name__ == "__main__":
  filePath = sys.argv[2]
  url = sys.argv[1]
  h = htmlImporter(filePath)
  h.writeFile(url)
