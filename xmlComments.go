// Copyright 2016 - 2022 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to
// and read from XLSX / XLSM / XLTM files. Supports reading and writing
// spreadsheet documents generated by Microsoft Excel™ 2007 and later. Supports
// complex components by high compatibility, and provided streaming API for
// generating or reading data from a worksheet with huge amounts of data. This
// library needs Go version 1.15 or later.

package excelize

import "encoding/xml"

// xlsxComments directly maps the comments element from the namespace
// http://schemas.openxmlformats.org/spreadsheetml/2006/main. A comment is a
// rich text note that is attached to and associated with a cell, separate from
// other cell content. Comment content is stored separate from the cell, and is
// displayed in a drawing object (like a text box) that is separate from, but
// associated with, a cell. Comments are used as reminders, such as noting how a
// complex formula works, or to provide feedback to other users. Comments can
// also be used to explain assumptions made in a formula or to call out
// something special about the cell.
type xlsxComments struct {
	XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main comments"`
	Authors     xlsxAuthor      `xml:"authors"`
	CommentList xlsxCommentList `xml:"commentList"`
}

// xlsxAuthor directly maps the author element. This element holds a string
// representing the name of a single author of comments. Every comment shall
// have an author. The maximum length of the author string is an implementation
// detail, but a good guideline is 255 chars.
type xlsxAuthor struct {
	Author []string `xml:"author"`
}

// xlsxCommentList (List of Comments) directly maps the xlsxCommentList element.
// This element is a container that holds a list of comments for the sheet.
type xlsxCommentList struct {
	Comment []xlsxComment `xml:"comment"`
}

// xlsxComment directly maps the comment element. This element represents a
// single user entered comment. Each comment shall have an author and can
// optionally contain richly formatted text.
type xlsxComment struct {
	Ref      string   `xml:"ref,attr"`
	AuthorID int      `xml:"authorId,attr"`
	Text     xlsxText `xml:"text"`
}

// xlsxText directly maps the text element. This element contains rich text
// which represents the text of a comment. The maximum length for this text is a
// spreadsheet application implementation detail. A recommended guideline is
// 32767 chars.
type xlsxText struct {
	T          *string          `xml:"t"`
	R          []xlsxR          `xml:"r"`
	RPh        *xlsxPhoneticRun `xml:"rPh"`
	PhoneticPr *xlsxPhoneticPr  `xml:"phoneticPr"`
}

// xlsxPhoneticRun element represents a run of text which displays a phonetic
// hint for this String Item (si). Phonetic hints are used to give information
// about the pronunciation of an East Asian language. The hints are displayed
// as text within the spreadsheet cells across the top portion of the cell.
type xlsxPhoneticRun struct {
	Sb uint32 `xml:"sb,attr"`
	Eb uint32 `xml:"eb,attr"`
	T  string `xml:"t"`
}

// formatComment directly maps the format settings of the comment.
type formatComment struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

// Comment directly maps the comment information.
type Comment struct {
	Author   string `json:"author"`
	AuthorID int    `json:"author_id"`
	Ref      string `json:"ref"`
	Text     string `json:"text"`
}
