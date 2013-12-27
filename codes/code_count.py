# -*- coding: cp936 -*-


import sys
import os


def isSLComment(str):
  comments = ('#', '//')
  for c in comments:
    if str.startswith(c):
      return True

def countFileLine(fname):  
  lines = open(fname).readlines()

  durMultiLineComment = False
  lineCnt = 0
  for l in lines:
    l = l.strip()

    # 跳过空行
    if len(l) == 0:
      continue

    # 跳过单行注释
    if isSLComment(l):
      continue

    #跳过多行注释
    if durMultiLineComment:
      if l.endswith("'''") or l.endswith('"""') or l.endswith('*/'):
        durMultiLineComment = False
      continue
    else:
      if l.startswith("'''") or l.startswith('"""') or l.startswith('/*'):
          # 多行注释只有一行？
          if not l.endswith("'''") and not l.endswith('"""') and not l.endswith('*/'):
            durMultiLineComment = True
          continue
    
    lineCnt += 1
  return lineCnt

def codeCount():
  args = sys.argv
  if len(args) == 1:
    print 'Please input the path'
    return
  if len(args) == 2:
    print 'Please input the suffix'
    return



  rootPath = sys.argv[1]
  suffix = sys.argv[2]
  print 'path:' + rootPath + ' suffix:' + rootPath

  fileCnt = 0
  lineCnt = 0
  for parent, dirnames, filenames in os.walk(rootPath):
    for name in filenames:
      if not name.startswith('_') and not name.startswith('test_') and name.endswith(suffix):
        fileCnt += 1
        fullpath = os.path.abspath(os.path.join(parent, name))
        #print fullpath
        curCnt = countFileLine(fullpath)
        print name, curCnt
        lineCnt += curCnt

  #dump the result
  print 'File count:', fileCnt
  print 'Line count:', lineCnt

codeCount()

