#!/usr/bin/env python3

def main():
	a = {}

	for i in range(1000000):
		stri = str(i)
		a[stri] = stri

	for i in range(1000000):
		stri = str(i)
		val = a[stri]

if __name__ == '__main__':
	main()
