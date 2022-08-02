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
	Uid         string             `bson:"uid,omitempty"`
	Url         string             `bson:"url"`
	Description string             `bson:"description,omitempty"`
}

// Link define the link between analias an it user
type Link struct {
	ID     primitive.ObjectID `bson:"_id"`
	Alias  Aliases            `bson:"alias"`
	Author User               `bson:"author"`
}

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Pseudo        string             `bson:"pseudo"`
	GithubProfile string             `bson:"github_profile"`
	CreatedAt     time.Time          `bson:"created_at"`
}
