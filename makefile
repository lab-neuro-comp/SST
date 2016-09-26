main:
	go build -o sst.exe github.com/lab-neuro-comp/SST/src/main

build: main
	rcedit sst.exe --set-icon data/icon.ico
