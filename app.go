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
	"bytes"
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-humble/router"
	"honnef.co/go/js/xhr"
	"myitcv.io/react"
)

type story struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	User          string `json:"user"`
	Time          int    `json:"time"`
	TimeAgo       string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Domain        string `json:"domain"`
}

type comment struct {
	ID            int       `json:"id"`
	User          string    `json:"user"`
	Time          int       `json:"time"`
	TimeAgo       string    `json:"time_ago"`
	Type          string    `json:"type"`
	Content       string    `json:"content"`
	Comments      []comment `json:"comments"`
	CommentsCount int       `json:"comments_count"`
	Level         int       `json:"level"`
	URL           string    `json:"url"`
	Dead          bool      `json:"dead,omitempty"`
}

type storyItem struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Points        int       `json:"points"`
	User          string    `json:"user"`
	Time          int       `json:"time"`
	TimeAgo       string    `json:"time_ago"`
	Type          string    `json:"type"`
	Content       string    `json:"content"`
	Comments      []comment `json:"comments"`
	CommentsCount int       `json:"comments_count"`
	URL           string    `json:"url"`
	Domain        string    `json:"domain"`
}

type user struct {
	About       string `json:"about"`
	CreatedTime int    `json:"created_time"`
	Created     string `json:"created"`
	ID          string `json:"id"`
	Karma       int    `json:"karma"`
}

// AppDef is the definition for the App component
type AppDef struct {
	react.ComponentDef
}

// AppState is the state types for the App component
type AppState struct {
	Loading       bool
	CurrRoute     string
	CurrPage      int
	Stories       []story
	SelectedStory storyItem
	SelectedUser  user
}

// App creates instances of the App component
func App() *AppElem {
	return buildAppElem()
}

// GetInitialState defines the initial state of the component
func (a AppDef) GetInitialState() AppState {
	return AppState{
		Loading:   false,
		CurrRoute: "news",
		CurrPage:  1,
	}
}

// Equals is used to define component re-rendering
func (c AppState) Equals(v AppState) bool {
	if len(v.Stories) != len(c.Stories) {
		return false
	}

	for i := range v.Stories {
		if v.Stories[i] != c.Stories[i] {
			return false
		}
	}

	if v.Loading != c.Loading || v.CurrPage != c.CurrPage || v.CurrRoute != c.CurrRoute || v.SelectedStory.ID != c.SelectedStory.ID || v.SelectedUser.ID != c.SelectedUser.ID {
		return false
	}

	return true
}

// ComponentDidMount is the hook that fires as soon as the component finishes mounting
func (a AppDef) ComponentDidMount() {
	Router := router.New()
	Router.ForceHashURL = true

	Router.HandleFunc("/", func(context *router.Context) {
		Router.Navigate("/news/1")
	})
	Router.HandleFunc("/news/{page}", a.showStories)
	Router.HandleFunc("/newest/{page}", a.showStories)
	Router.HandleFunc("/show/{page}", a.showStories)
	Router.HandleFunc("/ask/{page}", a.showStories)
	Router.HandleFunc("/jobs/{page}", a.showStories)

	Router.HandleFunc("/item/{id}", a.showStory)

	Router.HandleFunc("/user/{id}", a.showUser)

	Router.Start()
}

func (a AppDef) showUser(context *router.Context) {
	re, _ := regexp.Compile("^(.*?/.*?)/")
	newState := a.State()
	path := re.FindString(context.Path)

	newState.Loading = true
	newState.CurrRoute = strings.Replace(path, "/", "", 2)
	a.SetState(newState)

	go func() {
		a.getUser(context.Params["id"])
	}()
}

func (a AppDef) getUser(userID string) {
	newState := a.State()

	userChannel := make(chan user)
	go a.getUserRequest(userChannel, userID)
	newState.SelectedUser = <-userChannel
	newState.Loading = false

	a.SetState(newState)
}

func (a AppDef) getUserRequest(userChannel chan user, userID string) {
	url := "https://api.hnpwa.com/v0/user/" + userID + ".json"

	data, err := xhr.Send("GET", url, nil)
	if err != nil {
		println("Encountered error: ", err)
	}
	var user user
	json.NewDecoder(bytes.NewReader(data)).Decode(&user)

	userChannel <- user
}

func (a AppDef) showStory(context *router.Context) {
	re, _ := regexp.Compile("^(.*?/.*?)/")
	newState := a.State()
	path := re.FindString(context.Path)

	newState.Loading = true
	newState.CurrRoute = strings.Replace(path, "/", "", 2)
	a.SetState(newState)

	go func() {
		a.getStory(context.Params["id"])
	}()
}

func (a AppDef) getStory(storyD string) {
	newState := a.State()

	storyChannel := make(chan storyItem)
	go a.getStoryRequest(storyChannel, storyD)
	newState.SelectedStory = <-storyChannel
	newState.Loading = false

	a.SetState(newState)
}

func (a AppDef) getStoryRequest(storyChannel chan storyItem, storyD string) {
	url := "https://api.hnpwa.com/v0/item/" + storyD + ".json"

	data, err := xhr.Send("GET", url, nil)
	if err != nil {
		println("Encountered error: ", err)
	}
	var story storyItem
	json.NewDecoder(strings.NewReader(string(data))).Decode(&story)
	storyChannel <- story
}

func (a AppDef) showStories(context *router.Context) {
	re, _ := regexp.Compile("^(.*?/.*?)/")
	newState := a.State()
	path := re.FindString(context.Path)

	currPage, err := strconv.Atoi(context.Params["page"])

	if err != nil {
		print("Something went wrong!")
	}

	newState.Loading = true
	newState.CurrPage = currPage
	newState.CurrRoute = strings.Replace(path, "/", "", 2)

	a.SetState(newState)

	go func() {
		a.getStories(path, context.Params["page"])
	}()
}

func (a AppDef) getStories(storyType string, pageNum string) {
	newState := a.State()

	storiesChannel := make(chan []story)
	go a.getStoriesRequest(storiesChannel, storyType, pageNum)
	newState.Stories = <-storiesChannel
	newState.Loading = false

	a.SetState(newState)
}

func (a AppDef) getStoriesRequest(storiesChannel chan []story, storyType string, pageNum string) {
	url := "https://api.hnpwa.com/v0" + storyType + pageNum + ".json"

	data, err := xhr.Send("GET", url, nil)
	if err != nil {
		println("Encountered error: ", err)
	}
	var stories []story
	json.NewDecoder(strings.NewReader(string(data))).Decode(&stories)
	storiesChannel <- stories
}

func (a AppDef) renderLoader() react.Element {
	return react.Div(&react.DivProps{ClassName: "loader-container"},
		react.Div(&react.DivProps{ClassName: "loader"},
			react.Div(nil),
			react.Div(nil),
			react.Div(nil),
			react.Div(nil),
		),
	)
}

// Render renders the component
func (a AppDef) Render() react.Element {
	state := a.State()

	var content react.Element

	if state.Loading {
		if state.CurrRoute != "item" && state.CurrRoute != "user" {
			content = react.Fragment(nil,
				PageNav(PageNavProps{CurrPage: state.CurrPage, StoryType: state.CurrRoute, NumStories: len(state.Stories)}),
				react.Div(nil,
					react.Div(&react.DivProps{ClassName: "wrapper"},
						react.Div(&react.DivProps{ClassName: "story-list view"},
							react.Ul(
								&react.UlProps{ClassName: "skeleton"},
							),
						),
					),
				),
			)
		} else {
			content = a.renderLoader()
		}
	} else if state.CurrRoute == "item" {
		content = Story(StoryProps{
			ID:            state.SelectedStory.ID,
			Title:         state.SelectedStory.Title,
			Points:        state.SelectedStory.Points,
			User:          state.SelectedStory.User,
			Time:          state.SelectedStory.Time,
			TimeAgo:       state.SelectedStory.TimeAgo,
			Type:          state.SelectedStory.Type,
			Content:       state.SelectedStory.Content,
			Comments:      state.SelectedStory.Comments,
			CommentsCount: state.SelectedStory.CommentsCount,
			URL:           state.SelectedStory.URL,
			Domain:        state.SelectedStory.Domain,
		})
	} else if state.CurrRoute == "user" {
		content = User(UserProps{
			About:       state.SelectedUser.About,
			CreatedTime: state.SelectedUser.CreatedTime,
			Created:     state.SelectedUser.Created,
			ID:          state.SelectedUser.ID,
			Karma:       state.SelectedUser.Karma,
		})
	} else {
		content = react.Div(nil,
			PageNav(PageNavProps{CurrPage: state.CurrPage, StoryType: state.CurrRoute, NumStories: len(state.Stories)}),
			StoryList(StoryListProps{StoryItems: state.Stories}),
		)
	}

	return react.Div(nil,
		Header(HeaderProps{path: state.CurrRoute}),
		content,
	)
}
