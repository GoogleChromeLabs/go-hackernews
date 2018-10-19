/**
 * Copyright 2018 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package main

import (
	"strconv"

	"myitcv.io/react"
)

// CommentCardDef is the definition for the CommentCard component
type CommentCardDef struct {
	react.ComponentDef
}

// CommentCardProps is the prop types for the CommentCard component
type CommentCardProps struct {
	ID            int
	User          string
	Time          int
	TimeAgo       string
	Type          string
	Content       string
	Comments      []comment
	CommentsCount int
	Level         int
	URL           string
	Dead          bool
}

// CommentCardState is the state types for the CommentCard component
type CommentCardState struct {
	ToggleOpen bool
}

// CommentCard creates instances of the CommentCard component
func CommentCard(p CommentCardProps) *CommentCardElem {
	return buildCommentCardElem(CommentCardProps{ID: p.ID, User: p.User, Time: p.Time, TimeAgo: p.TimeAgo, Type: p.Type, Content: p.Content, Comments: p.Comments, CommentsCount: p.CommentsCount, Level: p.Level, URL: p.URL, Dead: p.Dead})
}

// Equals is used to define component re-rendering
func (c CommentCardState) Equals(v CommentCardState) bool {
	if c.ToggleOpen != v.ToggleOpen {
		return false
	}

	return true
}

// Equals is used to define component re-rendering
func (c CommentCardProps) Equals(v CommentCardProps) bool {
	if c.ID != v.ID {
		return false
	}

	return true
}

// GetInitialState defines the initial state of the component
func (f CommentCardDef) GetInitialState() CommentCardState {
	return CommentCardState{
		ToggleOpen: true,
	}
}

// Render renders the component
func (f CommentCardDef) Render() *react.LiElem {
	props := f.Props()

	var SubCommentsList *react.DivElem

	if props.CommentsCount > 0 {
		SubCommentsList = f.renderNestedComments(props.Comments)
	}
	return react.Li(&react.LiProps{ClassName: "comment"},
		react.Div(
			&react.DivProps{ClassName: "by"},
			react.A(
				&react.AProps{Href: "#/user/" + props.User},
				react.S(props.User),
			),
			react.S(" "+props.TimeAgo+" "),
		),
		react.Div(
			&react.DivProps{ClassName: "text"},
			react.Div(&react.DivProps{
				DangerouslySetInnerHTML: react.NewDangerousInnerHTML(props.Content),
			}),
		),
		SubCommentsList,
	)
}

// RendersLi is used to define rendered component as a list item element
func (f CommentCardDef) RendersLi(*react.LiElem) {}

type toggleReplies struct{ CommentCardDef }

// OnClick is used to define an onclick event for the component
func (f CommentCardDef) OnClick(e *react.SyntheticMouseEvent) {
	newState := f.State()
	newState.ToggleOpen = !newState.ToggleOpen
	f.SetState(newState)
}

func (f CommentCardDef) renderNestedComments(nestedComments []comment) *react.DivElem {
	state := f.State()
	props := f.Props()

	var comments []react.RendersLi

	for _, comment := range nestedComments {
		comments = append(comments, CommentCard(CommentCardProps{
			ID:            comment.ID,
			User:          comment.User,
			Time:          comment.Time,
			TimeAgo:       comment.TimeAgo,
			Type:          comment.Type,
			Content:       comment.Content,
			Comments:      comment.Comments,
			CommentsCount: comment.CommentsCount,
			Level:         comment.Level,
			URL:           comment.URL,
			Dead:          comment.Dead,
		}))
	}

	var content *react.DivElem

	if state.ToggleOpen {
		content = react.Div(nil,
			react.Div(
				&react.DivProps{ClassName: "toggle open"},
				react.Span(
					&react.SpanProps{OnClick: toggleReplies{f}},
					react.S("[-]"),
				),
			),
			react.Ul(
				&react.UlProps{ClassName: "comment-children"},
				comments...,
			),
		)
	} else {
		content = react.Div(nil,
			react.Div(
				&react.DivProps{ClassName: "toggle"},
				react.Span(
					&react.SpanProps{OnClick: toggleReplies{f}},
					react.S("[+] "+strconv.Itoa(props.CommentsCount)+" replies collapsed"),
				),
			),
		)
	}

	return content
}
