FROM golang

WORKDIR /Bot-Discord

COPY Cache ./Cache
COPY Config ./Config
COPY Controllers ./Controllers
COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
