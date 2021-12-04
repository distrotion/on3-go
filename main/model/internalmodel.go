package internalmodel

type ClinicProfileModel struct {
	SUID       string `json:"sUID"`
	NStatus    int    `json:"nStatus"`
	SBytePhoto string `json:"sBytePhoto"`
	SName      string `json:"sName"`
	SAddress   string `json:"sAddress"`
	SKwang     string `json:"sKwang"`
	SKhet      string `json:"sKhet"`
	SProvince  string `json:"sProvince"`
	SPost      string `json:"sPost"`
	SCountry   string `json:"sCountry"`
	SEmail     string `json:"sEmail"`
	STel       string `json:"sTel"`
	SLineID    string `json:"sLineId"`
	SLat       string `json:"sLat"`
	SLong      string `json:"sLong"`
}

type ClinicProfile struct {
	SUID           string   `json:"sUID"`
	RsStaffUID     []string `json:"rsStaffUID"`
	RsCaregiverUID []string `json:"rsCaregiverUID"`
	RsPatientUID   []string `json:"rsPatientUID"`
}

type StaffProfileModel struct {
	SUID           string `json:"sUID"`
	NStatus        int    `json:"nStatus"`
	SBytePhoto     string `json:"sBytePhoto"`
	SEmail         string `json:"sEmail"`
	SPass          string `json:"sPass"`
	SFirstName     string `json:"sFirstName"`
	SLastName      string `json:"sLastName"`
	SNickName      string `json:"sNickName"`
	NGender        int    `json:"nGender"`
	SDateBirth     string `json:"sDateBirth"`
	SDrOfficialID  string `json:"sDrOfficialId"`
	SAddress       string `json:"sAddress"`
	SKwang         string `json:"sKwang"`
	SKhet          string `json:"sKhet"`
	SProvince      string `json:"sProvince"`
	SPost          string `json:"sPost"`
	SCountry       string `json:"sCountry"`
	STel           string `json:"sTel"`
	SLineID        string `json:"sLineId"`
	NImgAvatarID   int    `json:"nImgAvatarId"`
	NTitle         int    `json:"nTitle"`
	SMDCardNumbers string `json:"sMDCardNumbers"`
	RnSymptom      []int  `json:"rnSymptom"`
}

type StaffProfile struct {
	SUID           string   `json:"sUID"`
	RsClinicUID    []string `json:"rsClinicUID"`
	RsCaregiverUID []string `json:"rsCaregiverUID"`
	RsPatientUID   []string `json:"rsPatientUID"`
}

type TeleSession struct {
	SUID        string `json:"sUID"`
	SRoomId     string `json:"sRoomId"`
	SStartDate  string `json:"sStartDate"`
	SEndDate    string `json:"sEndDate"`
	SDoctorUID  string `json:"sDoctorUID"`
	SPatientUID string `json:"sPatientUID"`
	SGuest1UID  string `json:"sGuest1UID"`
	SGuest2UID  string `json:"sGuest2UID"`
}

type StaffProfileModelED struct {
	NStatus        int    `json:"nStatus"`
	SBytePhoto     string `json:"sBytePhoto"`
	SEmail         string `json:"sEmail"`
	SFirstName     string `json:"sFirstName"`
	SLastName      string `json:"sLastName"`
	SNickName      string `json:"sNickName"`
	NGender        int    `json:"nGender"`
	SDateBirth     string `json:"sDateBirth"`
	SDrOfficialID  string `json:"sDrOfficialId"`
	SAddress       string `json:"sAddress"`
	SKwang         string `json:"sKwang"`
	SKhet          string `json:"sKhet"`
	SProvince      string `json:"sProvince"`
	SPost          string `json:"sPost"`
	SCountry       string `json:"sCountry"`
	STel           string `json:"sTel"`
	SLineID        string `json:"sLineId"`
	NImgAvatarID   int    `json:"nImgAvatarId"`
	NTitle         int    `json:"nTitle"`
	SMDCardNumbers string `json:"sMDCardNumbers"`
	RnSymptom      []int  `json:"rnSymptom"`
}

type StaffProfileED struct {
	RsClinicUID    []string `json:"rsClinicUID"`
	RsCaregiverUID []string `json:"rsCaregiverUID"`
	RsPatientUID   []string `json:"rsPatientUID"`
}

type PatientProfileModel struct {
	SUID            string `json:"sUID"`
	NtimestampFront int    `json:"ntimestampFront"`
	NStatus         int    `json:"nStatus"`
	SBytePhoto      string `json:"sBytePhoto"`
	NImgAvatarID    int    `json:"nImgAvatarId"`
	SEmail          string `json:"sEmail"`
	SPass           string `json:"sPass"`
	SFirstName      string `json:"sFirstName"`
	SLastName       string `json:"sLastName"`
	SNickName       string `json:"sNickName"`
	NGender         int    `json:"nGender"`
	SDateBirth      string `json:"sDateBirth"`
	SNationID       string `json:"sNationId"`
	SAddress        string `json:"sAddress"`
	SKwang          string `json:"sKwang"`
	SKhet           string `json:"sKhet"`
	SProvince       string `json:"sProvince"`
	SPost           string `json:"sPost"`
	SCountry        string `json:"sCountry"`
	STel            string `json:"sTel"`
	SLineID         string `json:"sLineId"`
	NWhoPay         int    `json:"nWhoPay"`
	Ntimestamp      int    `json:"ntimestamp"` //23
	IsGotCaregiver  bool   `json:"isGotCaregiver"`
}

type PatientProfile struct {
	SUID string `json:"sUID"`
	// RnSymptom      []int       `json:"rnSymptom"`
	RsCaregiverUID []string `json:"rsCaregiverUID"`
	RsStaffUID     []string `json:"rsStaffUID"`
	RsClinicUID    []string `json:"rsClinicUID"`
	// RsDoctorUID    []string    `json:"rsDoctorUID"`
	Ntimestamp   int         `json:"ntimestamp"`
	STemppass    string      `json:"sTemppass"`
	PrsDoctorUID []*MsDoctor `json:"prsDoctorUID"`
}

type MsDoctor struct {
	NSymptom int    `json:"nSymptom"`
	SUIDS    string `json:"sUIDS"`
}

type PatientProfileModelEdir struct {
	NtimestampFront int    `json:"ntimestampFront"`
	NStatus         int    `json:"nStatus"`
	SBytePhoto      string `json:"sBytePhoto"`
	NImgAvatarID    int    `json:"nImgAvatarId"`
	SEmail          string `json:"sEmail"`
	SFirstName      string `json:"sFirstName"`
	SLastName       string `json:"sLastName"`
	SNickName       string `json:"sNickName"`
	NGender         int    `json:"nGender"`
	SDateBirth      string `json:"sDateBirth"`
	SNationID       string `json:"sNationId"`
	SAddress        string `json:"sAddress"`
	SKwang          string `json:"sKwang"`
	SKhet           string `json:"sKhet"`
	SProvince       string `json:"sProvince"`
	SPost           string `json:"sPost"`
	SCountry        string `json:"sCountry"`
	STel            string `json:"sTel"`
	SLineID         string `json:"sLineId"`
	NWhoPay         int    `json:"nWhoPay"`
	Ntimestamp      int    `json:"ntimestamp"`
	IsGotCaregiver  bool   `json:"isGotCaregiver"`
}

type PatientProfileEdit struct {
	RnSymptom      []int    `json:"rnSymptom"`
	RsCaregiverUID []string `json:"rsCaregiverUID"`
	RsStaffUID     []string `json:"rsStaffUID"`
	RsClinicUID    []string `json:"rsClinicUID"`
	RsDoctorUID    []string `json:"rsDoctorUID"`
	Ntimestamp     int      `json:"ntimestamp"`
}

type CaregiverProfileModel struct {
	SUID          string `json:"sUID"`
	NStatus       int    `json:"nStatus"`
	SBytePhoto    string `json:"sBytePhoto"`
	NImgAvatarID  int    `json:"nImgAvatarId"`
	SFirstName    string `json:"sFirstName"`
	SLastName     string `json:"sLastName"`
	SNickName     string `json:"sNickName"`
	NRelation     int    `json:"nRelation"`
	NFamilySize   int    `json:"nFamilySize"`
	NMarriage     int    `json:"nMarriage"`
	SDateBirth    string `json:"sDateBirth"`
	SNationID     string `json:"sNationId"`
	IsSameAddress bool   `json:"IsSameAddress"`
	SAddress      string `json:"sAddress"`
	SKwang        string `json:"sKwang"`
	SKhet         string `json:"sKhet"`
	SProvince     string `json:"sProvince"`
	SPost         string `json:"sPost"`
	SCountry      string `json:"sCountry"`
	SEmail        string `json:"sEmail"`
	SLineID       string `json:"sLineId"`
	STel          string `json:"sTel"`
	// RsPatientUID  []string `json:"rsPatientUid"`
}

type CaregiverProfile struct {
	SUID         string   `json:"sUID"`
	RsPatientUID []string `json:"rsPatientUID"`
	RsClinicUID  []string `json:"rsClinicUID"`
	RsStaffUID   []string `json:"rsStaffUID"`
}

type PatientEdit struct {
	NtimestampFront int         `json:"ntimestampFront"`
	NStatus         int         `json:"nStatus"`
	SBytePhoto      string      `json:"sBytePhoto"`
	NImgAvatarID    int         `json:"nImgAvatarId"`
	SEmail          string      `json:"sEmail"`
	SFirstName      string      `json:"sFirstName"`
	SLastName       string      `json:"sLastName"`
	SNickName       string      `json:"sNickName"`
	NGender         int         `json:"nGender"`
	SDateBirth      string      `json:"sDateBirth"`
	SNationID       string      `json:"sNationId"`
	SAddress        string      `json:"sAddress"`
	SKwang          string      `json:"sKwang"`
	SKhet           string      `json:"sKhet"`
	SProvince       string      `json:"sProvince"`
	SPost           string      `json:"sPost"`
	SCountry        string      `json:"sCountry"`
	STel            string      `json:"sTel"`
	SLineID         string      `json:"sLineId"`
	NWhoPay         int         `json:"nWhoPay"`
	RsDoctorUID     []*MsDoctor `json:"rsDoctorUID"`
	IsGotCaregiver  bool        `json:"isGotCaregiver"`
}

type CaregiverEdit struct {
	NStatus       int    `json:"cnStatus"`
	SBytePhoto    string `json:"csBytePhoto"`
	NImgAvatarID  int    `json:"cnImgAvatarId"`
	SFirstName    string `json:"csFirstName"`
	SLastName     string `json:"csLastName"`
	SNickName     string `json:"csNickName"`
	NRelation     int    `json:"cnRelation"`
	NFamilySize   int    `json:"cnFamilySize"`
	NMarriage     int    `json:"nMarriage"`
	SDateBirth    string `json:"csDateBirth"`
	SNationID     string `json:"csNationId"`
	IsSameAddress bool   `json:"cIsSameAddress"`
	SAddress      string `json:"csAddress"`
	SKwang        string `json:"csKwang"`
	SKhet         string `json:"csKhet"`
	SProvince     string `json:"csProvince"`
	SPost         string `json:"csPost"`
	SCountry      string `json:"csCountry"`
	SEmail        string `json:"csEmail"`
	SLineID       string `json:"csLineId"`
	STel          string `json:"csTel"`
}

type StaffProfileEditIN struct {
	NStatus        int    `json:"nStatus"`
	SBytePhoto     string `json:"sBytePhoto"`
	SEmail         string `json:"sEmail"`
	SFirstName     string `json:"sFirstName"`
	SLastName      string `json:"sLastName"`
	SNickName      string `json:"sNickName"`
	NGender        int    `json:"nGender"`
	SDateBirth     string `json:"sDateBirth"`
	SDrOfficialID  string `json:"sDrOfficialId"`
	SAddress       string `json:"sAddress"`
	SKwang         string `json:"sKwang"`
	SKhet          string `json:"sKhet"`
	SProvince      string `json:"sProvince"`
	SPost          string `json:"sPost"`
	SCountry       string `json:"sCountry"`
	STel           string `json:"sTel"`
	SLineID        string `json:"sLineId"`
	NImgAvatarID   int    `json:"nImgAvatarId"`
	NTitle         int    `json:"nTitle"`
	SMDCardNumbers string `json:"sMDCardNumbers"`
}
