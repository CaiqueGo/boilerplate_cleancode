package user

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (ref User) GetID() int64 {
	return ref.ID
}

func (ref User) GetName() string {
	return ref.Name
}
