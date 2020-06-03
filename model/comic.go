package model

import (
	"encoding/json"
	"fmt"
)

// XkcdResponse is the struct representation of XKCD comic http response
type XkcdResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// FormattedDate formats individual date elements into a single string
func (x XkcdResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", x.Day, x.Month, x.Year)
}

// Comic converts ComicResponse that we receive from the API to our
// application's output format
func (x XkcdResponse) Comic() Comic {
	return Comic{
		Title:       x.Title,
		Number:      x.Num,
		Date:        x.FormattedDate(),
		Description: x.Alt,
		Image:       x.Img,
	}
}

// Comic is the struct representation of the output of this app
type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

// PrettyString creates a pretty string of the Comic that we'll use as
// output
func (c Comic) PrettyString() string {
	p := fmt.Sprintf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		c.Title, c.Number, c.Date, c.Description, c.Image,
	)
	return p
}

// JSON converts the Comic struct to JSON, we'll use the JSON string as
// output
func (c Comic) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cJSON)
}
