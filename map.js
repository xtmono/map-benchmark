function main() {
	var a = {};
	var val;

	for (var i = 0; i < 1000000; i++) {
		a[i] = i;
	}

	for (var i = 0; i < 1000000; i++) {
		val = a[i];
	}
}

main();
