package location

import (
	"pasarwarga/apiformat"
	"strings"
)


type LocationFormatter struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
	 
}

type LocationCityFormatter struct{
	Id int `json:"id"`
	Id_provinsi string `json:"id_provinsi"`
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

func ListLocationKotaFormatter(lokasi []apiformat.GetFormatLokasiKota) []LocationCityFormatter {

	formatter := []LocationCityFormatter{}

	for _, listlokasi := range lokasi{
		Detailformatter := DetailLocationKotaFormatter(listlokasi)
		formatter = append(formatter, Detailformatter)
	}
	return formatter
}

func DetailLocationKotaFormatter(lokasi apiformat.GetFormatLokasiKota) LocationCityFormatter {

	detailformatter := LocationCityFormatter{}

	detailformatter.Id = lokasi.Id
	//fmt.Println(len(lokasi.Name))
	detailformatter.Id_provinsi = lokasi.Id_provinsi
	t := strings.Replace(lokasi.Name,"Kabupaten ", "", -1)
	detailformatter.City = t


	return detailformatter

}





