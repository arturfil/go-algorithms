
category=all
exclude=false


build: 
	go build -o exec ./main/main.go

run:
	./exec -category=$(category) -exclude=$(exclude)
