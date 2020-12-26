APPS     :=	edge

$(APPS):
	go build -o bin/$@ cmd/$@/*.go

clean:
	rm -rf bin/*
