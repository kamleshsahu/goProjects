package main

type LibraryRepo struct {
	userMap     map[int]User
	UserIdIndex int
}

func GetLibraryRepo() LibraryRepoI {
	return &LibraryRepo{userMap: make(map[int]User), UserIdIndex: 1}
}

func (l *LibraryRepo) CreateUser(user User) (*User, CustomError) {
	user.Id = l.UserIdIndex
	l.userMap[user.Id] = user
	l.UserIdIndex = l.UserIdIndex + 1
	return &user, nil
}

func (l *LibraryRepo) UpdateUser(id int, user User) (updatedUser *User, customError CustomError) {
	_, oldUser := l.userMap[id]
	if !oldUser {
		return nil, "user not found"
	}
	l.userMap[id] = user

	return &user, nil
}

func (l *LibraryRepo) DeleteUser(id int) CustomError {
	_, user := l.userMap[id]
	if !user {
		return "user not found"
	}
	delete(l.userMap, id)
	return nil
}

func (l *LibraryRepo) GetUsers(i interface{}) ([]User, CustomError) {
	users := make([]User, 0)

	for _, user := range l.userMap {
		users = append(users, user)
	}

	return users, nil
}
