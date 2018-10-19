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

// PageNavDef is the definition for the PageNav component
type PageNavDef struct {
	react.ComponentDef
}

// PageNavProps is the prop types for the PageNav component
type PageNavProps struct {
	CurrPage   int
	StoryType  string
	NumStories int
}

// PageNav creates instances of the PageNav component
func PageNav(p PageNavProps) *PageNavElem {
	return buildPageNavElem(PageNavProps{CurrPage: p.CurrPage, StoryType: p.StoryType, NumStories: p.NumStories})
}

// Render renders the PageNav component
func (f PageNavDef) Render() react.Element {
	props := f.Props()

	var prevLink react.Element
	var nextLink react.Element

	if props.CurrPage == 1 {
		prevLink = react.Span(
			&react.SpanProps{ClassName: "disabled"},
			react.S("< prev"),
		)
	} else {
		prevLink = react.A(
			&react.AProps{Href: "/#/" + props.StoryType + "/" + strconv.Itoa(props.CurrPage-1)},
			react.S("< prev"),
		)
	}

	if props.NumStories == 30 {
		nextLink = react.A(
			&react.AProps{Href: "/#/" + props.StoryType + "/" + strconv.Itoa(props.CurrPage+1)},
			react.S("more >"),
		)
	} else {
		nextLink = react.Span(
			&react.SpanProps{ClassName: "disabled"},
			react.S("more >"),
		)
	}

	return react.Div(&react.DivProps{ClassName: "page-nav-container"},
		react.Div(&react.DivProps{ClassName: "page-nav"},
			prevLink,
			react.Span(
				nil,
				react.S(react.S(strconv.Itoa(props.CurrPage))),
			),
			nextLink,
		),
	)
}
