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

// StoryDef is the definition for the Story component
type StoryDef struct {
	react.ComponentDef
}

// StoryProps is the prop types for the Story component
type StoryProps struct {
	ID            int
	Title         string
	Points        int
	User          string
	Time          int
	TimeAgo       string
	Type          string
	Content       string
	Comments      []comment
	CommentsCount int
	URL           string
	Domain        string
}

// Story creates instances of the Story component
func Story(p StoryProps) *StoryElem {
	return buildStoryElem(StoryProps{ID: p.ID, Title: p.Title, Points: p.Points, User: p.User, Time: p.Time, TimeAgo: p.TimeAgo, Type: p.Type, Content: p.Content, Comments: p.Comments, CommentsCount: p.CommentsCount, URL: p.URL, Domain: p.Domain})
}

// Equals is used to define component re-rendering
func (c StoryProps) Equals(v StoryProps) bool {
	if c.ID != v.ID {
		return false
	}

	return true
}

// Render renders the Story component
func (f StoryDef) Render() react.Element {
	props := f.Props()

	var comments []react.RendersLi

	if len(props.Comments) > 0 {
		for _, comment := range props.Comments {
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
	}

	domainStr := ""
	if props.Domain != "" {
		domainStr = "        (" + props.Domain + ")"
	}

	return react.Div(nil,
		react.Div(&react.DivProps{ClassName: "wrapper"},
			react.Div(&react.DivProps{ClassName: "view"},
				react.Div(&react.DivProps{ClassName: "item-view-header"},
					react.A(
						&react.AProps{Target: "_blank", Href: props.URL, ClassName: "github"},
						react.H1(
							nil,
							react.S(props.Title),
						),
					),
					react.Span(
						&react.SpanProps{ClassName: "host"},
						react.S(domainStr),
					),
					react.P(
						&react.PProps{ClassName: "meta"},
						react.S("						"+strconv.Itoa(props.Points)+" points						| by  "),
						react.A(
							&react.AProps{Href: "#/user/" + props.User},
							react.S(props.User),
						),
						react.S(" 						"+props.TimeAgo+"					"),
					),
				),
				react.Div(&react.DivProps{ClassName: "item-view-comments"},
					react.P(&react.PProps{ClassName: "item-view-comments-header"},
						react.S(strconv.Itoa(props.CommentsCount)+" comments"),
					),
					react.Ul(
						&react.UlProps{ClassName: "comment-children"},
						comments...,
					),
				),
			),
		),
	)
}
