# Telling to use Docker's golang ready image
FROM golang
# Create app folder 
RUN mkdir /server
# Copy our file in the host contianer to our contianer
ADD . /server
# Set /app to the go folder as workdir
WORKDIR /server
# Generate binary file from our /app
RUN go build
# Expose the port 3000
EXPOSE 8080:8080
# Run the app binarry file 
CMD ["./server"]
