myapp: myapp.go
	go build -tags netgo -installsuffix netgo -o myapp myapp.go
