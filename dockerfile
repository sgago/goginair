# Our base image is golang
FROM golang:latest as base

WORKDIR /app

# Install air for live reloading
RUN go install github.com/cosmtrek/air@latest

# Copy the files we need from our local project to the work directory (WORKDIR)
COPY go.mod go.sum .air.toml ./

# Copy our all our go src files over, air will call build and run on this for us
COPY ./src/main ./src/main

# Copy our static HTML go templates over
COPY ./src/templates ./../templates

# Download our go packages
RUN go mod download

# See the .air.toml file's build bin and cmd properties for the specific cmds

# Expose ports to the outside world
EXPOSE 8080

CMD ["air", "-c", ".air.toml"]