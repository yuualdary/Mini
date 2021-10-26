package location

import "pasarwarga/apiformat"

type LocationFormatter struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
	 
}

type LocationCityFormatter struct{
	Id int `json:"id"`
	Id_provinsi int `json:"id_provinsi"`
	City string `json:"nama"`
}

func ListLocationFormatter(lokasi []apiformat.GetFormatLokasi) []LocationFormatter {

	formatter := []LocationFormatter{}

	for _, listlokasi := range lokasi{
		Detailformatter := DetailLocationFormatter(listlokasi)
		formatter = append(formatter, Detailformatter)
	}
	return formatter
}

func DetailLocationFormatter(lokasi apiformat.GetFormatLokasi) LocationFormatter {

	detailformatter := LocationFormatter{}

	detailformatter.Id = lokasi.Id
	detailformatter.Nama = lokasi.Name

	return detailformatter

}




