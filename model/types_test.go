package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {

	data := `{
  "fromUser": {
   "avatarUrl": "https://avatars-02.gitter.im/gh/uv/3/foobar",
   "avatarUrlMedium": "https://avatars1.githubusercontent.com/u/626234?v=3&s=128",
   "avatarUrlSmall": "https://avatars1.githubusercontent.com/u/6264234?v=3&s=60",
   "displayName": "Foo bar",
   "gv": "3",
   "id": "234023420349",
   "url": "/foo",
   "username": "foo",
   "v": 11
  },
  "html": "New private room. This one is actually private and not just obscured ",
  "id": "551be8de4ceaec8225f8f17b",
  "issues": [],
  "mentions": [],
  "meta": [],
  "readBy": 10,
  "sent": "2015-04-01T12:47:26.093Z",
  "text": "New private room. This one is actually private and not just obscured ",
  "unread": false,
  "urls": [],
  "v": 1
 }`
	var message Message
	err := json.Unmarshal([]byte(data), &message)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("text : %v", message.Text)
}
