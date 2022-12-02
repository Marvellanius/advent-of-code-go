run: 
	go run src/Y$(YEAR)/Day$(DAY)/main.go -part $(PART)

test:
	cd src/Y$(YEAR)/Day$(DAY)/; go test -v

input:
	go run commands/input/main.go -day $(DAY) -year $(YEAR)

scaffold:
	go run commands/scaffold/main.go -day $(DAY) -year $(YEAR)

debug:
	dlv debug src/Y$(YEAR)/Day$(DAY)/main.go