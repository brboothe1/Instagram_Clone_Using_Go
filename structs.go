package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"time"
)

/*structs.go houses the structures used to build tables in the PostgreSQL database

github.com/jinzhu/gorm is used to provide unique creation values to declutter the struct fields
it is used to more easily integrate building tables with Go to PostgreSQL*/


type User struct {
	gorm.Model

	Username 			string		`gorm:"uniqueIndex"`
	Password 			string
	Email 				string		`gorm:"uniqueIndex"`
	Name				string
	Biography			string
	Website				string
	ProfilePic			string
	PhoneNumber			string
	Gender				string
	Followers			[]User
	FollowerCount		uint64
	Following			[]User
	FollowingCount		uint64
	StoryList			[]Story
	Photos				[]Photo
	Posts				[]Post
	Comments			[]Comment
	Likes				[]Like
	Newsfeeds			[]Newsfeed
	Stories				[]Story
	ChatRooms			[]ChatRoom
	DirectMessages		[]DirectMessage	`gorm:"uniqueIndex"`
	SenderUsername		string
	SenderUsernames		[]DirectMessage	`gorm:"foreignKey:SenderUsername;references:SenderUsername"`
	ReceiverUsername	string
	ReceiverUsernames	[]DirectMessage	`gorm:"foreignKey:ReceiverUsername;references:ReceiverUsername"`
}

type Photo struct {
	gorm.Model

	UserID 		uint64
	FileName 	string
	Location 	string
	Path 		string		`gorm:"uniqueIndex"`
	UploadTime  time.Time
	Likes 		uint64

}

type Post struct {
	gorm.Model

	Description string
	PhotoID		uint64
	UserID		uint64
	PhotoOrder	uint64
	CreateTime	time.Time
}

type Comment struct {
	gorm.Model

	UserID		uint64
	PhotoID		uint64
	Text 		string
	TimeWritten	time.Time
	Likes		uint64
}

type Like struct {
	gorm.Model

	UserID		uint64
	PhotoID		uint64
	CommentID	uint64
}

type Newsfeed struct {
	gorm.Model

	UserID 		uint64
	PostID		uint64
//	PostRank	int
}

type Story struct {
	gorm.Model

	UserID		uint64
	PhotoID		uint64
	PhotoOrder	uint8
	UploadTime	time.Time
	ExpireTime	time.Time
}

type ChatRoom struct {
	gorm.Model

	DirectMessages	[]DirectMessage
}

type DirectMessage struct {
	gorm.Model

	ChatRoomID			uint64
	SenderUserID 		uint64
	SenderUsername		string
	ReceiverUserID		uint64
	ReceiverUsername	string
	MessageText			string
	MessageTime			time.Time
}


// testing structs to make sure database does not error while adding to tables

var (
	user1 = &User{Username: "photo.boothe", Password: "323321", Email: "br.boothe1@gmail.com", Name: "Brandon Boothe", Biography: "Software Engineer, Boxer, Photographer", Website: "Mywebsite.com", ProfilePic: "Here", Followers: nil, Following: nil, StoryList: nil, Photos: nil, Posts: nil, Comments: nil, Likes: nil, Newsfeeds: nil, ChatRooms: nil, DirectMessages: nil, SenderUsernames: nil, ReceiverUsernames: nil}
	photo = []Photo{
		{UserID: 1, FileName: "Red Door.png", Path: "pathtophoto/photo1.png"},
		{UserID: 1, FileName: "Broward Park.png", Path: "pathtophoto/photo2.png"},
	}
)