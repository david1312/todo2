FROM golang:1.14 
# ini ambil aplikasi golnag dari dockerhub
WORKDIR /Users/david.bernardi/ba-it/go/src/github.com/david1312/todo2git
COPY . .