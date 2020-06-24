

bin:
	go build github.com/soypat/decimate

test:
	./decimate.exe -x x -y y csvtools/testdata/t.csv
# 	./decimate.exe -x x -y y csvtools/testdata/t.csv