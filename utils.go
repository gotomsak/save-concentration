package main

import (
	"context"
	"encoding/binary"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnect() (database *mongo.Client, Context context.Context) {
	USER := os.Getenv("USERR")
	PASS := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOLMONGO")
	uri := "mongodb://" + USER + ":" + PASS + "@" + PROTOCOL

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	if err != nil {
		panic(err.Error())
	}
	return c, ctx
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
}

// Bytes2uint converts []byte to uint64
func Bytes2uint(bytes []byte) uint64 {
	padding := make([]byte, 8-len(bytes))
	i := binary.BigEndian.Uint64(append(padding, bytes...))
	return i
}

// Bytes2int converts []byte to int64
func Bytes2int(bytes []byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := Bytes2uint(bytes)
		i = (^i + 1) & mask
		return int64(-i)

	} else {
		i := Bytes2uint(bytes)
		return int64(i)
	}
}

// Uint2bytes converts uint64 to []byte
func Uint2bytes(i uint64, size int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, i)
	return bytes[8-size : 8]
}

// Int2bytes converts int to []byte
func Int2bytes(i int, size int) []byte {
	var ui uint64
	if 0 < i {
		ui = uint64(i)
	} else {
		ui = (^uint64(-i) + 1)
	}
	return Uint2bytes(ui, size)
}
