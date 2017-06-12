function main() {
	var a = {};
	var val;

	for (var i = 0; i < 1000000; i++) {
		var stri = i.toString();
		a[stri] = stri;
	}

	for (var i = 0; i < 1000000; i++) {
		var stri = i.toString();
		val = a[stri];
	}
}

main();
