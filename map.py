#!/usr/bin/env python3

def main():
	a = {}

	for i in range(1000000):
		a[i] = i

	for i in range(1000000):
		val = a[i]

if __name__ == '__main__':
	main()
