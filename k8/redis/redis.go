package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"redis/dbs"
	"time"
)

var ctx = context.Background()

const NumberOfCandidates = 4
const candidatesCodeLength = 3

type CandidateCode [candidatesCodeLength]byte

type VotingPackage struct {
	Codes      [NumberOfCandidates]CandidateCode `bson:"codes" json:"codes"`
	VoteSerial primitive.Binary                  `bson:"voteSerial" json:"voteSerial"`
	Used       bool                              `bson:"used"`
}

func main() {
	res, err := dbs.QueryDb[VotingPackage](context.Background(), bson.D{{"used", true}})
	if err != nil {
		panic(err)
	}
	fmt.Printf("wyniki := %v\n", res)

	res, err = dbs.QueryDb[VotingPackage](context.Background(), bson.D{{"used", true}})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 6)
	res, err = dbs.QueryDb[VotingPackage](context.Background(), bson.D{{"used", true}})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pageviews: %v\n", res)
}

func getFromDb() {

}
