package mongo

import (
	"context"
	root "skynet/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(session *Session, config *root.MongoConfig) *UserService {
	collection := session.client.Database(config.DbName).Collection("User")

	return &UserService{collection}
}

func (userServ *UserService) CreateUser(u *root.User) error {
	user, err := newUserModel(u)
	if err != nil {
		return err
	}

	_, err = userServ.collection.InsertOne(context.TODO(), user)

	return err
}

func (userServ *UserService) GetUserByUsername(username string) (root.User, error) {
	model := userModel{}
	filter := bson.D{{"username", username}}

	err := userServ.collection.FindOne(context.TODO(), filter).Decode(&model)

	result := root.User{
		Identifier: model.Identifier,
		UserName:   model.UserName,
		Password:   "--"}

	return result, err
}

func (userServ *UserService) Login(cred root.Credentials) (root.User, error, bool) {
	model := userModel{}
	filter := bson.D{{"username", cred.UserName}}

	err := userServ.collection.FindOne(context.TODO(), filter).Decode(&model)
	if err != nil {
		return root.User{}, err, false
	}

	err = model.comparePassword(cred.Password)
	if err != nil {
		return root.User{}, err, false
	}

	result := root.User{
		Identifier: model.Identifier,
		UserName:   model.UserName,
		Password:   "--"}

	return result, nil, true
}
