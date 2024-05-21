package main

type LibraryService struct {
	libraryRepo LibraryRepoI
}

func GetLibraryService(libraryRepo LibraryRepoI) LibraryServiceI {
	return &LibraryService{libraryRepo}
}

func (l *LibraryService) CreateUser(user User) (*User, CustomError) {
	return l.libraryRepo.CreateUser(user)
}

func (l *LibraryService) UpdateUser(id int, user User) (updatedUser *User, customError CustomError) {
	return l.libraryRepo.UpdateUser(id, user)
}

func (l *LibraryService) DeleteUser(id int) CustomError {
	return l.libraryRepo.DeleteUser(id)
}

func (l *LibraryService) GetUsers(i interface{}) ([]User, CustomError) {
	return l.libraryRepo.GetUsers(i)
}
