# Base image
FROM golang:1.15.6

# Disable go cache as it can generate conflicts on cached tests results
RUN GOCACHE=OFF
RUN go env -w GOPRIVATE=github.com/aipetto

# Configure the repo url so we can configure our work directory.
ENV REPO_URL=github.com/aipetto/go-aipetto-users-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/username/go-aipetto-users-api/src
# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY . $WORKPATH
WORKDIR $WORKPATH/src

## Enable Token Access to our private repos
RUN git config --global url."https://aipetto:91d8c820da5be370cb03a065c394cf5c27ddbaa4@github.com".insteadOf "https://github.com"

RUN go mod tidy

RUN go build -o go-users-api .

# Expose port 8081 to the service
EXPOSE 8081

CMD ["./go-users-api"]