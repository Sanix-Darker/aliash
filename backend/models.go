package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Aliases isa raw alias defined by an uid and a raw_url where to fetch
type Aliases struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"title,omitempty"`
	Uid         string             `bson:"uid"`
	Content     string             `bson:"content,omitempty"`
	Description string             `bson:"description"`
	Hash512     string             `bson:"hash,omitempty"`
}

// Link define the link between analias an it user
type Link struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Alias  Aliases            `bson:"alias"`
	Author User               `bson:"author"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"pseudo"`
	CreatedAt time.Time          `bson:"created_at"`
}

type ProfileType struct {
	Github    string `default:"github"`
	Bitbucket string `default:"bitbucket"`
	Gitlab    string `default:"gitlab"`
}

type Profile struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	User User               `bson:"author"`
	Type ProfileType        `bson:"profile_type"`
}
