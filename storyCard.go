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
	"strings"

	"myitcv.io/react"
)

// StoryCardDef is the definition for the StoryCard component
type StoryCardDef struct {
	react.ComponentDef
}

// StoryCardProps is the prop types for the StoryCard component
type StoryCardProps struct {
	ID            int
	title         string
	points        int
	commentsCount int
	domain        string
	timeAgo       string
	user          string
	storyType     string
	URL           string
}

// StoryCard creates instances of the StoryCard component
func StoryCard(s StoryCardProps) *StoryCardElem {
	return buildStoryCardElem(StoryCardProps{ID: s.ID, title: s.title, points: s.points, commentsCount: s.commentsCount, domain: s.domain, timeAgo: s.timeAgo, user: s.user, URL: s.URL, storyType: s.storyType})
}

// Render renders the StoryCard component
func (f StoryCardDef) Render() *react.LiElem {
	props := f.Props()

	var userSpan react.Element
	if props.user == "" {
		userSpan = nil
	} else {
		userSpan = react.Span(
			&react.SpanProps{ClassName: "by"},
			react.S(" by "),
			react.A(
				&react.AProps{Href: "#/user/" + props.user},
				react.S(" "+props.user),
			),
		)
	}

	var commentsSpan react.Element
	if props.storyType == "job" {
		commentsSpan = nil
	} else {
		commentsSpan = react.Span(
			&react.SpanProps{ClassName: "comments-link"},
			react.S(" | "),
			react.A(
				&react.AProps{Href: "#/item/" + strconv.Itoa(props.ID)},
				react.S(strconv.Itoa(props.commentsCount)+" comments"),
			),
		)
	}

	link := props.URL
	if strings.HasPrefix(props.URL, "item") {
		link = "#/item/" + strconv.Itoa(props.ID)
	}

	domainStr := ""
	if props.domain != "" {
		domainStr = " (" + props.domain + ")"
	}

	return react.Li(&react.LiProps{ClassName: "story-card"},
		react.Span(
			&react.SpanProps{ClassName: "score"},
			react.S(strconv.Itoa(props.points)),
		),
		react.Span(
			&react.SpanProps{ClassName: "title"},
			react.A(
				&react.AProps{Href: link},
				react.S(props.title),
			),
			react.Span(
				&react.SpanProps{ClassName: "host"},
				react.S(domainStr),
			),
		),
		react.Br(nil, nil),
		react.Span(
			&react.SpanProps{ClassName: "meta"},
			userSpan,
			react.Span(
				&react.SpanProps{ClassName: "time"},
				react.S(" "+props.timeAgo+" "),
			),
			commentsSpan,
		),
	)
}

// RendersLi is used to return the rendered StoryCard component as a list item element
func (f StoryCardDef) RendersLi(*react.LiElem) {}
