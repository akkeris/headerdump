FROM golang:1.7
ADD . .
CMD './start.sh'
EXPOSE 9000
