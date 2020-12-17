package model

type Name struct {
	FirstName string
	MidName string
	ThirdName string
	LastName string
	LastLastName string
}

func NewName(firstName, midName, thirdName, lastName, lastLastName string) *Name {
	name := &Name{
		FirstName: firstName,
		MidName: midName,
		ThirdName: thirdName,
		LastName: lastName,
		LastLastName: lastLastName,
	}
	return name
}

func (n *Name) getShortName() string {
	return n.FirstName + " " + n.LastName
}

func (n *Name) getFullName() string {
	return n.FirstName + " " + n.MidName + " " + n.ThirdName + " " + n.LastName + " " + n.LastLastName
}
