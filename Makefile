
category=all
from75=false


build:
	go build -o exec ./main/main.go

run:
	./exec -category=$(category) -from75=$(from75)
