package postgres

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "github.com/dilshodforever/4-oyimtixon-game-service/genprotos/learning"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Candidates    []Candidate   `json:"candidates"`
	UsageMetadata UsageMetadata `json:"usageMetadata"`
}

type Candidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason"`
	Index         int            `json:"index"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type Content struct {
	Parts []Part `json:"parts"`
	Role  string `json:"role"`
}

type Part struct {
	Text string `json:"text"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}

type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

type AiStorage struct {
	db *sql.DB
}

func NewAiStorage(db *sql.DB) *AiStorage {
	return &AiStorage{db: db}
}

func (ls *LearningStorage) SubmitChallengeSolution(req *pb.SubmitChallengeSolutionRequest) (*pb.SubmitChallengeSolutionResponse, error) {
	apiKey := "AIzaSyAqYyftXSGM9ImO8AJmLxhswmJ8BNmp6iE"
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text":"Sen fizika boyicha olim sifatida faqat s",
					},
					{
						"text": req.Solution,
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, err
	}

	res, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	res.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return nil, err
	}

	chat := &pb.SubmitChallengeSolutionResponse{Message: response.Candidates[0].Content.Parts[0].Text}
	coll := ls.db.Collection("aistorage")
	documents := []interface{}{
		bson.D{
			{Key: "Userid", Value: req.Userid},
			{Key: "RequestText", Value: req.Solution},
			{Key: "ResponsetText", Value: chat.Message},
		},
	}

	_, err = coll.InsertMany(context.Background(), documents)
	if err != nil {
		log.Printf("Failed to insert feedback: %v", err)
		return nil, err
	}
	return chat, nil
}
func (ls *LearningStorage) GetIAstorage(req *pb.AistorageRequest) (*pb.AistorageResponse, error) {
	coll := ls.db.Collection("aistorage")
	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Failed to get AI storage: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	airesponse := &pb.AistorageResponse{}
	for cursor.Next(context.Background()) {
		var topic pb.AistorageRespons
		if err := cursor.Decode(&topic); err != nil {
			log.Printf("Failed to decode AI storage item: %v", err)
			return nil, err
		}
		if topic.Userid == req.Userid {
			airesponse.Response = append(airesponse.Response, &topic)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return airesponse, nil
}
