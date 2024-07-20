package postgres

import (
	"context"
	"log"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
)

func (ls *LearningStorage) GetResources(req *pb.GetResourcesRequest) (*pb.GetResourcesResponse, error) {
	coll := ls.db.Collection("resources")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get resources: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	var resources []*pb.Resource
	for cursor.Next(context.Background()) {
		var resource pb.Resource
		if err := cursor.Decode(&resource); err != nil {
			log.Printf("Failed to decode resource: %v", err)
			return nil, err
		}
		resources = append(resources, &resource)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return &pb.GetResourcesResponse{Resources: resources}, nil
}

func (ls *LearningStorage) CompleteResource(req *pb.CompleteResourceRequest) (*pb.CompleteResourceResponse, error) {
	_, err := ls.UpdateUserXp(&pb.Update{UserId: req.UserId, Xps: 10})
	if err != nil {
		log.Printf("Failed to update userxps: %v", err)
		return nil, err
	}
	_, err = ls.UpdateComplateds(&pb.CalculateCompleteds{ResourcesCompleted: 1, Userid: req.UserId})
	if err != nil {
		log.Printf("Failed to update complates: %v", err)
		return nil, err
	}
	return &pb.CompleteResourceResponse{
		Message:  "Resource completed successfully",
		XpEarned: 10,
	}, nil
}
