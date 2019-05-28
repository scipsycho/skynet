package mongo

import (
	"context"
	root "skynet/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type RecordService struct {
	collection *mongo.Collection
}

func NewRecordService(session *Session, config *root.MongoConfig) *RecordService {
	collection := session.client.Database(config.DbName).Collection("Record")

	return &RecordService{collection}
}

func (recServ *RecordService) CreateRecord(rec *root.Record) error {
	record, err := newRecordModel(rec)
	if err != nil {
		return err
	}

	_, error := recServ.collection.InsertOne(context.TODO(), record)

	return error
}

func (recServ *RecordService) GetAllRecords() ([]root.Record, error) {
	var results []root.Record

	findOptions := options.Find()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := recServ.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var singleRecord root.Record
		err := cur.Decode(&singleRecord)
		if err != nil {
			return nil, err
		}

		results = append(results, singleRecord)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return results, nil
}
