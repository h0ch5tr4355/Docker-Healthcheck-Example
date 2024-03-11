# Stage 1: Build-Stage
FROM golang:latest AS builder

WORKDIR /

# Kopiere den Go-Quellcode in das Arbeitsverzeichnis des Containers
COPY . .


# Kompiliere das Go-Programm
RUN go build -o sample_app cmd/main.go

# Stage 2: Finaler Stage
FROM debian:latest

WORKDIR /app

# Kopiere das kompilierte Go-Programm aus der Build-Stage in den finalen Container
COPY --from=builder sample_app .

# Exponiere den Port 7777, auf dem das Go-Programm läuft
EXPOSE 7777

# Führe das kompilierte Go-Programm aus
CMD ["./sample_app"]
