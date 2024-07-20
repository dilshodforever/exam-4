package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	
	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
)

func setup(t *testing.T) (*LearningStorage, func()) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	db := mt.Client.Database("topics")
	ls := NewLearningStorage(db)
	return ls, func() {
		//mt.Coll.Clone()
	}
}

func TestGetTopics(t *testing.T) {
	ls, teardown := setup(t)
	defer teardown()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	

	mt.Run("success", func(mt *mtest.T) {
		expectedTopics := []*pb.Topic{
			{Id: "topic1", Name: "Topic 1"},
			{Id: "topic2", Name: "Topic 2"},
		}

		first := mtest.CreateCursorResponse(1, "test_db.topics", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "id", Value: "topic1"},
			{Key: "name", Value: "Topic 1"},
		})
		second := mtest.CreateCursorResponse(1, "test_db.topics", mtest.NextBatch, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "id", Value: "topic2"},
			{Key: "name", Value: "Topic 2"},
		})
		killCursors := mtest.CreateCursorResponse(0, "test_db.topics", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		req := &pb.GetTopicsRequest{}
		resp, err := ls.GetTopics(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.ElementsMatch(t, expectedTopics, resp.Topics)
	})

	mt.Run("fail", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "Mock error",
		}))

		req := &pb.GetTopicsRequest{}
		resp, err := ls.GetTopics(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestSubmitFeedback(t *testing.T) {
	ls, teardown := setup(t)
	defer teardown()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		req := &pb.SubmitFeedbackRequest{
			Userid:  "user1",
			TopicId: "topic1",
			Rating:  5,
			Comment: "Great topic!",
		}
		resp, err := ls.SubmitFeedback(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Feedback submitted successfully", resp.Message)
		assert.Equal(t, int32(10), resp.XpEarned)
	})

	mt.Run("fail", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "Mock error",
		}))

		req := &pb.SubmitFeedbackRequest{
			Userid:  "user1",
			TopicId: "topic1",
			Rating:  5,
			Comment: "Great topic!",
		}
		resp, err := ls.SubmitFeedback(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestStartGame(t *testing.T) {
	ls, teardown := setup(t)
	defer teardown()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		req := &pb.StartRequest{
			Userid: "user1",
		}
		resp, err := ls.StartGame(req)
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Success", resp.Message)
	})

	mt.Run("fail", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "Mock error",
		}))

		req := &pb.StartRequest{
			Userid: "user1",
		}
		resp, err := ls.StartGame(req)
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
