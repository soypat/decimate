buildflags = -ldflags="-s -w" -i
binname = decimate
distr:
	go build ${buildflags} -o bin/${binname}.exe
	cp README.md README.txt
	zip ${binname} -j bin/${binname}.exe README.txt
	rm README.txt
mkbin:
	mkdir bin

test:
	./decimate.exe -x x -y y testdata/t.csv
	./decimate.exe -x a,b -y ,, testdata/a.tsv
clean: rm ${binname}.zip