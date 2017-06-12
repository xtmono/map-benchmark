#!/usr/bin/env python3

import threading

a = {}

class putThread (threading.Thread):
	def __init__(self):
		threading.Thread.__init__(self)
	def run(self):
		for i in range(1000000):
			stri = str(i)
			a[stri] = stri

class getThread (threading.Thread):
	def __init__(self):
		threading.Thread.__init__(self)
	def run(self):
		for i in range(1000000):
			stri = str(i)
			val = a[stri]

def main():
	threads = []

	for i in range(8):
		thread = putThread()
		thread.start()
		threads.append(thread)
	for t in threads:
		t.join()

	for i in range(8):
		thread = getThread()
		thread.start()
		threads.append(thread)
	for t in threads:
		t.join()

if __name__ == '__main__':
	main()
