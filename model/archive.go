// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package model is the "model" layer which defines entity structures with ORM and controller.
package model

// Archive model.
type Archive struct {
	Model

	Year         string `gorm:"size:4" json:"year"`
	Month        string `gorm:"size:2" json:"month"`
	ArticleCount int    `json:"articleCount"`

	BlogID uint `sql:"index" json:"blogID"`
}
