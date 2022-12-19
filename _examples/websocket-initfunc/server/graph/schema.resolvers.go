package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"
)

// PostMessageTo is the resolver for the postMessageTo field.
func (r *mutationResolver) PostMessageTo(ctx context.Context, subscriber string, content string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// PostAnnouncement is the resolver for the postAnnouncement field.
func (r *mutationResolver) PostAnnouncement(ctx context.Context, content string) (string, error) {
	if ch, ok := chanMap["ann"]; ok {
		select {
		case ch <- content:
		default:
			delete(chanMap, "ann")
		}
	} else {
		fmt.Println("No subscription found!")
	}

	return content, nil
}

// Foo is the resolver for the foo field.
func (r *queryResolver) Foo(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented: Foo - foo"))
}

// Subscribe is the resolver for the subscribe field.
func (r *subscriptionResolver) Subscribe(ctx context.Context, subscriber string) (<-chan string, error) {
	fmt.Println(ctx.Value("username"))
	ch := make(<-chan string)

	return ch, nil
}

// Announcement is the resolver for the announcement field.
func (r *subscriptionResolver) Announcement(ctx context.Context, subscriber string) (<-chan string, error) {
	var ch chan string

	if _, ok := chanMap["ann"]; !ok {
		ch = make(chan string)
		chanMap["ann"] = ch
	} else {
		ch = chanMap["ann"]
	}

	return ch, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
