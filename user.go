package oauth2providers

type UserInfo interface {
	setPictureURL(pictureURL string)
	setEmail(email string)
	setName(name string)
	setID(id string)

	GetPictureURL() string
	GetEmail() string
	GetName() string
	GetID() string
}

type userInfo struct {
	ID         string
	Email      string
	Name       string
	PictureURL string
}

func NewUserInfo() UserInfo {
	return &userInfo{}
}

func (u *userInfo) setPictureURL(pictureURL string) {
	u.PictureURL = pictureURL
}

func (u *userInfo) GetPictureURL() string {
	return u.PictureURL
}

func (u *userInfo) setEmail(email string) {
	u.Email = email
}

func (u *userInfo) GetEmail() string {
	return u.Email
}

func (u *userInfo) setName(name string) {
	u.Name = name
}

func (u *userInfo) GetName() string {
	return u.Name
}

func (u *userInfo) setID(id string) {
	u.ID = id
}
func (u *userInfo) GetID() string {
	return u.ID
}
