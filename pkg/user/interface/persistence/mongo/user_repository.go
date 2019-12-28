package mongo

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/stereoit/eventival/pkg/user/domain/model"
	"github.com/stereoit/eventival/pkg/user/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var errNotImplemented = errors.New("Not implemented")

// UserRepositoryOpts are configuration options for the DB
type UserRepositoryOpts struct {
	ConnectionURI string
	Database      string
	Collection    string
}

// User struct maps the domain entity model.User into storage model
type User struct {
	ID        string `bson:"_id"`
	Email     string `bson:"email"`
	FirstName string `bson:"firstname"`
	LastName  string `bson:"lastname"`
}

type userRepository struct {
	*mongo.Client
	connectionURI string
	database      string
	collection    string
}

// NewUserRepository returns instance of User MongoDB repository
func NewUserRepository(opts *UserRepositoryOpts) (repository.UserRepository, error) {
	client, err := getClient(opts.ConnectionURI)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
		return nil, err
	}

	return &userRepository{
		client,
		opts.ConnectionURI,
		opts.Database,
		opts.Collection,
	}, nil
}

func getClient(connection string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connection)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Doublecheck this behaviour

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Here we implement full repository interface methods from domain.UserRepository

func (r *userRepository) FindAll() ([]*model.User, error) {
	filter := bson.D{}
	collection := r.Client.Database(r.database).Collection(r.collection)

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var users []*User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return toDomainUserList(users), nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	findFilter := bson.D{primitive.E{
		Key:   "email",
		Value: email,
	}}
	collection := r.Client.Database(r.database).Collection(r.collection)

	var result User
	err := collection.FindOne(context.TODO(), findFilter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	user := model.NewUser(result.ID, result.Email)
	return user, nil
}

func (r *userRepository) FindByID(id string) (*model.User, error) {
	findFilter := bson.D{primitive.E{Key: "_id", Value: id}}
	collection := r.Client.Database(r.database).Collection(r.collection)
	var result User

	err := collection.FindOne(context.TODO(), findFilter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	user := model.NewUser(result.ID, result.Email)
	return user, nil
}

func (r *userRepository) Save(user *model.User) error {

	collection := r.Client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(context.Background(), toMongoUser(user))
	if err != nil {
		return err
	}

	return nil
}

func toMongoUser(user *model.User) *User {
	return &User{
		ID:        user.GetID(),
		Email:     user.GetEmail(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
	}
}

func toDomainUserList(users []*User) []*model.User {
	var domainUsers []*model.User
	for _, user := range users {
		domainUsers = append(domainUsers, model.NewUser(user.ID, user.Email))
	}
	return domainUsers
}
