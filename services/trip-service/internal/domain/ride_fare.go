package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID          primitive.ObjectID
	UserID      string
	PackageSlug string // van van , luxury etc.
	//TotalPriceIncents float64
	//ExpiresAt         time.Time
}
