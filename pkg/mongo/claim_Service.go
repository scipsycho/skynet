package mongo

import (
	"context"
	"log"
	root "skynet/pkg"
	"skynet/pkg/crypto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClaimService struct {
	claimCollection     *mongo.Collection
	claimDefnCollection *mongo.Collection
}

func NewClaimService(session *Session, config *root.MongoConfig) *ClaimService {
	claimCollection := session.client.Database(config.DbName).Collection("Claim")
	claimDefnCollection := session.client.Database(config.DbName).Collection("ClaimDefn")

	return &ClaimService{claimCollection, claimDefnCollection}
}

func (claimServ *ClaimService) CreateClaimDefn(attributesToTypes map[string]string, username string, commonName string) error {

	claimDefn := root.ClaimDefn{UserIdentifier: username, CommonName: commonName}

	var c crypto.Crypto
	var err error

	claimDefn.ClaimDefnIdentifier, err = c.GenerateRandomASCIIString(20)
	if err != nil {
		log.Fatal(err)
	}
	claimDefn.AttributesToType = attributesToTypes

	_, err = claimServ.claimDefnCollection.InsertOne(context.TODO(), claimDefn)

	return nil
}

func (claimServ *ClaimService) CreateClaim(claim *root.Claim) error {

	// claim := root.Claim{UserIdentifier: userID, CommonName: commonName, ClaimDefnIdentifier: claimDefnID}
	_, err := claimServ.claimCollection.InsertOne(context.TODO(), *claim)
	return err
}

func (claimServ *ClaimService) GetClaimDefnByCommonName(IName string, CName string) (root.ClaimDefn, error) {

	var claimDefn root.ClaimDefn
	filter := bson.D{{"commonname", CName}}

	err := claimServ.claimDefnCollection.FindOne(context.TODO(), filter).Decode(&claimDefn)

	if err != nil {
		log.Fatal(err)
	}

	return claimDefn, nil
}

func (claimServ *ClaimService) GetAllClaimDefns() ([]root.ClaimDefn, error) {
	var claimDefns []root.ClaimDefn

	findOptions := options.Find()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := claimServ.claimDefnCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var claimDefn root.ClaimDefn
		err := cur.Decode(&claimDefn)
		if err != nil {
			return nil, err
		}

		claimDefns = append(claimDefns, claimDefn)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claimDefns, nil
}

/*
func (claimServ *ClaimService) GetClaimByUserID(identifier string) ([]root.Claim, error) {

	var claims []root.Claim
	claim := root.Claim{}
	filter := bson.D{{"useridentifier", identifier}}

	cur, _ := claimServ.claimCollection.Find(context.TODO(), filter, options.Find())

	for cur.Next(context.TODO()) {
		err := cur.Decode(&claim)
		if err != nil {
			return nil, err
		}

		claims = append(claims, claim)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claims, nil
}

func (claimServ *ClaimService) GetClaimByCommonName(identifier string, commonName string) (root.Claim, error) {

	var claims []root.Claim
	claim := root.Claim{}
	filter := bson.D{{"useridentifier", identifier}}

	cur, _ := claimServ.claimCollection.Find(context.TODO(), filter, options.Find())

	for cur.Next(context.TODO()) {
		err := cur.Decode(&claim)
		if err != nil {
			return root.Claim{}, err
		}

		claims = append(claims, claim)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	resultIndex := 0
	for index, i := range claims {
		if i.CommonName == commonName {
			resultIndex = index
		}
	}

	return claims[resultIndex], nil
}

func (claimServ *ClaimService) GetClaimDefnByClaimDefnID(identifier string) ([]root.ClaimDefn, error) {

	var claimDefns []root.ClaimDefn
	claimDefn := root.ClaimDefn{}
	filter := bson.D{{"claimdefnidentifier", identifier}}

	cur, _ := claimServ.claimDefnCollection.Find(context.TODO(), filter, options.Find())

	for cur.Next(context.TODO()) {
		err := cur.Decode(&claimDefn)
		if err != nil {
			return nil, err
		}

		claimDefns = append(claimDefns, claimDefn)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claimDefns, nil
}
*/
func (claimServ *ClaimService) GetAllClaims() ([]root.Claim, error) {
	var claims []root.Claim

	findOptions := options.Find()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := claimServ.claimCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {

		var claim root.Claim
		err := cur.Decode(&claim)
		if err != nil {
			return nil, err
		}

		claims = append(claims, claim)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	return claims, nil
}
