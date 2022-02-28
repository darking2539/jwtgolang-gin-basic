FROM golang:alpine

# ENV GIN_MODE=release
# ENV PROD=TURE
# ENV GORM_URL="root:root@tcp(192.168.200.31:3306)/root?charset=utf8mb4&parseTime=True&loc=Local"
# ENV JWT_SECRET="123456"
# ENV PORT="8081"

# Creating the `app` directory in which the app will run 
RUN mkdir /app

# Move everything from root to the newly created app directory
ADD . /app

# Specifying app as our work directory in which
# futher instructions should run into
WORKDIR /app

# Download all neededed project dependencies
RUN go mod download

#  Build the project executable binary
RUN go build -o main .

# Run/Starts the app executable binary
CMD ["/app/main"]