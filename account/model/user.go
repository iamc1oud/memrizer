package model

type User struct {
	UID uuid.UUID `db:"uid" json:"uid"`
	Email string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Name string `db:"name" json:"name"`
	ImageUrl string `db:"image_url" json:"imageUrl"`
	Website string `db:"website" json:"website"`
}