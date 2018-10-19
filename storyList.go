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
	"myitcv.io/react"
)

// StoryListDef is the definition for the StoryList component
type StoryListDef struct {
	react.ComponentDef
}

// StoryListProps is the prop types for the StoryList component
type StoryListProps struct {
	StoryItems []story
}

// StoryList creates instances of the StoryList component
func StoryList(p StoryListProps) *StoryListElem {
	return buildStoryListElem(StoryListProps{StoryItems: p.StoryItems})
}

// Equals is used to define component re-rendering
func (c StoryListProps) Equals(v StoryListProps) bool {
	if len(v.StoryItems) != len(c.StoryItems) {
		return false
	}

	for i := range v.StoryItems {
		if v.StoryItems[i] != c.StoryItems[i] {
			return false
		}
	}

	return true
}

// Render renders the StoryList component
func (f StoryListDef) Render() react.Element {
	props := f.Props()

	var storyItems []react.RendersLi
	var storyList []react.RendersLi

	if len(props.StoryItems) > 0 {
		for _, story := range props.StoryItems {
			storyItems = append(storyItems, StoryCard(StoryCardProps{
				ID:            story.ID,
				title:         story.Title,
				points:        story.Points,
				commentsCount: story.CommentsCount,
				domain:        story.Domain,
				timeAgo:       story.TimeAgo,
				user:          story.User,
				URL:           story.URL,
				storyType:     story.Type,
			}))
		}

		storyList = storyItems
	} else {
		storyList = append(storyList, react.Li(&react.LiProps{ClassName: "story-card"},
			react.Span(
				&react.SpanProps{ClassName: "title"},
				react.S("No more story items!")),
		))
	}

	return react.Div(nil,
		react.Div(&react.DivProps{ClassName: "wrapper"},
			react.Div(&react.DivProps{ClassName: "story-list view"},
				react.Ul(
					nil,
					storyList...,
				),
			),
		),
	)
}
