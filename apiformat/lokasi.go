package apiformat

type MessageData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DataLokasi struct {
	MessageData
	Data []GetFormatLokasi `json:"provinsi"`
}

type GetFormatLokasi struct {
	Id   int    `json:"id"`
	Name string `json:"nama"`
}

type DataLokasiKota struct {
	MessageData
	Data []GetFormatLokasi `json:"kota_kabupaten"`
}