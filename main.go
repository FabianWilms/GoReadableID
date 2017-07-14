package main

import (
	"github.com/FabianWilms/GoReadableID/readableId"
	"log"
	"time"
)

func main(){
	tStart := time.Now();
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	log.Printf("Generated ID: %v", readableId.GetRandomID())
	since := time.Since(tStart)
	log.Printf("Took time: %v", since)
}