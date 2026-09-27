package file

import "github.com/google/uuid"

type File struct {
	UUID uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &File{
		UUID: fileId,
		Name: filename,
		Data: blob,
	}, nil
}

func (f *File) String() string {
	return "File{UUID: " + f.UUID.String() + ", Name: " + f.Name + ", Data: " + string(f.Data) + "}"
}
