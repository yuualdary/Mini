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

type GetFormatLokasiKota struct {
	Id          int    `json:"id"`
	Id_provinsi string `json:"id_provinsi"`
	Name        string `json:"nama"`
}
type DataLokasiKota struct {
	MessageData
	Data []GetFormatLokasiKota `json:"kota_kabupaten"`
}

// type DataDetailLokasiKota struct {
// 	MessageData
// 	Data GetFormatLokasiKota `json:"data"`
// }
type DataListLokasiKota struct {
	MessageData
	Data []GetFormatLokasiKota `json:"data"`
}