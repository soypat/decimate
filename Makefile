

bin:
	go build github.com/soypat/decimate

test:
	./decimate.exe -x x -y y testdata/t.csv
	./decimate.exe -x a,b -y ,, testdata/a.tsv