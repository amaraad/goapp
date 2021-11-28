package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	database "github.com/amaraad/goapp/db"
	"github.com/amaraad/goapp/graph/generated"
	"github.com/amaraad/goapp/graph/model"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	question := model.Question{}
	question.QuestionText = input.QuestionText
	question.PubDate = input.PubDate
	database.DB.Create(&question)
	return &question, nil
}

func (r *mutationResolver) UpdateQuestion(ctx context.Context, input model.QuestionUpdateInput) (*model.Question, error) {
	question := model.Question{}
	database.DB.First(&question, input.ID)
	question.PubDate = input.PubDate
	question.QuestionText = input.QuestionText
	database.DB.Save(&question)
	return &question, nil
}

func (r *mutationResolver) DeleteQuestion(ctx context.Context, input int) (string, error) {
	question := model.Question{}
	database.DB.First(&question, input)
	database.DB.Delete(&question)
	result := "success"
	return result, nil
}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
	choice := model.Choice{}
	question := model.Question{}
	choice.QuestionID = input.QuestionID
	choice.ChoiceText = input.ChoiceText
	database.DB.First(&question, choice.QuestionID)
	choice.Question = &question
	database.DB.Create(&choice)
	return &choice, nil
}

func (r *mutationResolver) UpdateChoice(ctx context.Context, input model.ChoiceUpdateInput) (*model.Choice, error) {
	choice := model.Choice{}
	database.DB.First(&choice, input.ID)
	choice.QuestionID = input.QuestionID
	choice.ChoiceText = input.ChoiceText
	database.DB.Save(&choice)
	return &choice, nil
}

func (r *mutationResolver) DeleteChoice(ctx context.Context, input int) (string, error) {
	choice := model.Choice{}
	database.DB.First(&choice, input)
	database.DB.Delete(&choice)
	result := "success"
	return result, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	database.DB.Find(&r.questions)
	for _, question := range r.questions {
		var choices []*model.Choice
		database.DB.Where(&model.Choice{QuestionID: question.ID}).Find(&choices)
		question.Choices = choices
	}
	return r.questions, nil
}

func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	database.DB.Find(&r.choices)
	return r.choices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
