package misc

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringToObjectID(idStr string) (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("invalid objectID: %v", err)
	}
	return id, nil
}
