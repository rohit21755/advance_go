package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Server struct{}

func (s *Server) handleRequest(filename string) error {
	newFileName := hashFilename(filename)
	fmt.Println("new filename: ", newFileName)
	return nil
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func main() {
	s := &Server{}
	s.handleRequest("cool_picture.jpg")
}
