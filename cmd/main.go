package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

func main() {
	fmt.Println("===========================")
	fmt.Println("uuid - Shorten UUID v0.1.0")
	fmt.Println("===========================")

	shortVal := flag.String("s", "", "Shorten UUID Value")
	uuidVal := flag.String("l", "", "Long UUID Value")

	flag.Parse()

	if *shortVal != "" {
		uid, err := ShortToLong(*shortVal)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		fmt.Printf("Short UUID: %v\n", *shortVal)
		fmt.Printf("Long UUID: %v\n", uid)
	}

	if *uuidVal != "" {
		shortUUIDVal := LongStrToShort(*uuidVal)
		fmt.Printf("Long UUID: %s\n", *uuidVal)
		fmt.Printf("Short UUID: %v\n", shortUUIDVal)
	}
}

func ShortToLong(shortUUID string) (uuid.UUID, error) {
	btUid, err := base58.FastBase58DecodingAlphabet(shortUUID, base58.FlickrAlphabet)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return uuid.Nil, err
	}

	uid, err := uuid.FromBytes(btUid)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return uuid.Nil, err
	}

	return uid, nil
}

func LongStrToShort(uidStr string) string {
	uid, err := uuid.Parse(uidStr)
	if err != nil {
		fmt.Printf("Error: %v, %v", err, uidStr)
	}
	return LongToShort(uid)
}

func LongToShort(uid uuid.UUID) string {
	bUid, err := uid.MarshalBinary()
	if err != nil {
		fmt.Printf("Error: %v, %v", err, bUid)
	}

	shortUUID := base58.FastBase58EncodingAlphabet(bUid, base58.FlickrAlphabet)

	return shortUUID
}
