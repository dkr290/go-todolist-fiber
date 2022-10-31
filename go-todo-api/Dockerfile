FROM golang:1.19.2-alpine3.15 as builder

ARG dbhost
ARG dbuser
ARG dbpass
ARG dbport = "5432"
ARG dbname = "goTodo"



WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod tidy
COPY . .
RUN go build -o todo-api .


# Input parameters for the Dockerfile expected in os.Getenv
ENV DATABASE_HOST=$dbhost
ENV DATABASE_USER=$dbuser
ENV DATABASE_PASS=$dbpass
ENV DATABASE_PORT=$dbport
ENV DB_NAME=$dbname


FROM golang:1.19.2-alpine3.15
# Add maintainer info
LABEL maintainer="Danail Surudzhiyski"

COPY --from=builder /build/todo-api .

EXPOSE 8001

CMD ["./todo-api"]
