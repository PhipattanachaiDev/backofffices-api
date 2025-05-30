package models

type Province struct {
	ProvinceID   uint   
	ProvinceName string 

}

type District struct {
	DistrictID   uint  
	DistrictName string 
}

type SubDistrict struct {
	SubDistrictID   uint  
	SubDistrictName string 
	ZipCode         string
}