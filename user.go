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

// UserDef is the definition for the User component
type UserDef struct {
	react.ComponentDef
}

// UserProps is the prop types for the User component
type UserProps struct {
	About       string
	CreatedTime int
	Created     string
	ID          string
	Karma       int
}

// User creates instances of the User component
func User(p UserProps) *UserElem {
	return buildUserElem(UserProps{ID: p.ID, CreatedTime: p.CreatedTime, Created: p.Created, Karma: p.Karma, About: p.About})
}

// Render renders the User component
func (f UserDef) Render() react.Element {
	props := f.Props()

	return react.Div(nil,
		react.Div(&react.DivProps{ClassName: "wrapper"},
			react.Div(&react.DivProps{ClassName: "user-view view"},
				react.H1(
					nil,
					react.S("User : "+props.ID),
				),
				react.Ul(
					&react.UlProps{ClassName: "meta"},
					react.Li(
						nil,
						react.Span(
							&react.SpanProps{ClassName: "label"},
							react.S("Created:"),
						),
						react.S(" "+props.Created),
					),
					react.Li(
						nil,
						react.Span(
							&react.SpanProps{ClassName: "label"},
							react.S("Karma:"),
						),
						react.S(" "+strconv.Itoa(props.Karma)),
					),
					react.Li(
						&react.LiProps{ClassName: "about", DangerouslySetInnerHTML: react.NewDangerousInnerHTML(props.About)},
					),
				),
				react.P(
					&react.PProps{ClassName: "links"},
					react.A(
						&react.AProps{Target: "_blank", Href: "https://news.ycombinator.com/submitted?id=" + props.ID},
						react.S("submissions"),
					),
					react.S("  |  "),
					react.A(
						&react.AProps{Target: "_blank", Href: "https://news.ycombinator.com/threads?id=" + props.ID},
						react.S("comments"),
					),
				),
			),
		),
	)
}
