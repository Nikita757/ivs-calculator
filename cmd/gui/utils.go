/*
 * Copyright (c) 2021 Tomas Martykan
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"github.com/gotk3/gotk3/gtk"
)

/**
 * Utility function to get the text content of a Gtk TextView
 * @param textView A Gtk TextView widget
 * @return TextView content
 */
func TextView_GetText(textView *gtk.TextView) string {
	buffer, _ := textView.GetBuffer()
	start, end := buffer.GetBounds()
	text, _ := buffer.GetText(start, end, true)
	return text
}

/**
 * Utility function to set the text content of a Gtk TextView
 * @param textView A Gtk TextView widget
 * @param text Text content
 */
func TextView_SetText(textView *gtk.TextView, text string) {
	buffer, _ := textView.GetBuffer()
	buffer.SetText(text)
	textView.SetBuffer(buffer)
}
