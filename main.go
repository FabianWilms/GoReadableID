package main

import (
	"github.com/FabianWilms/GoReadableID/readableId"
	"log"
	time2 "time"
)

func main(){
	tStart := time2.Now();
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
	since := time2.Since(tStart)
	log.Printf("Took time: %v", since)
}