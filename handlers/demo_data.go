package handlers

var DemoVideo = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		CommentCount:  0,
		FavoriteCount: 0,
		IsFavorite:    false,
		Title:         "bear",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestJack",
	FollowCount:   1,
	FollowerCount: 1,
	IsFollow:      true,
}

var DemoUserList = []User{
	DemoUser,
}

var DemoCommentList = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "This is a test comment.^_^",
		CreateDate: "2023-02-15",
	},
}

var DemoUserLoginInfo = map[string]User{
	"jack12345678": {
		Id:            1,
		Name:          "jack",
		FollowerCount: 20,
		FollowCount:   5,
		IsFollow:      true,
	},
	"mike12345678910": {
		Id:            2,
		Name:          "mike",
		FollowerCount: 30,
		FollowCount:   5,
		IsFollow:      true,
	},
}
