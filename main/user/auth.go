package auth

import (
	"context"
	"fmt"
	model "go-resfull/graph/model"
	"go-resfull/jwt"
	internalmodel "go-resfull/main/model"
	"go-resfull/mongo/maindbv2"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/naamancurtis/mongo-go-struct-to-bson/mapper"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateStaffProfile(ctx context.Context, input model.StaffProfileRegister, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string, clinicDB string, clinicDB_profile_nouse string, clinicDB_profile_model string) string {
	t1 := time.Now().Unix()
	t2 := time.Now().UnixNano()

	// //------------------------------------- clinic ID

	SessionUIDC := jwt.ParseToken(input.SSessionUIDc, 7)
	SessionUIDPArC := strings.Split(SessionUIDC, `<>`)
	UIDC := SessionUIDPArC[0]

	dbs1C := maindbv2.Finddb(ctx, clinicDB, clinicDB_profile_model, bson.M{"SUID": UIDC}, "_id", 1, 0, 0)
	if len(dbs1C) == 0 {
		output := `no have clinic`
		return output
	}

	// //------------------------------------- clinic ID

	sUID := `s` + "-" + strconv.FormatInt(t1, 16) + strconv.FormatInt(t2, 16) + RandStringRunes(5)
	// fmt.Println(sUID)

	var StaffProMD internalmodel.StaffProfileModel

	StaffProMD.SUID = sUID
	StaffProMD.NStatus = input.NStatus
	StaffProMD.SBytePhoto = input.SBytePhoto
	StaffProMD.SEmail = input.SEmail
	StaffProMD.SPass = input.SPass
	StaffProMD.SFirstName = input.SFirstName
	StaffProMD.SLastName = input.SLastName
	StaffProMD.SNickName = input.SNickName
	StaffProMD.NGender = input.NGender
	StaffProMD.SDateBirth = input.SDateBirth
	StaffProMD.SDrOfficialID = input.SDrOfficialID
	StaffProMD.SAddress = input.SAddress
	StaffProMD.SKwang = input.SKwang
	StaffProMD.SKhet = input.SKhet
	StaffProMD.SProvince = input.SProvince
	StaffProMD.SPost = input.SPost
	StaffProMD.SCountry = input.SCountry
	StaffProMD.STel = input.STel
	StaffProMD.SLineID = input.SLineID
	StaffProMD.NImgAvatarID = input.NImgAvatarID
	StaffProMD.NTitle = input.NTitle
	StaffProMD.SMDCardNumbers = input.SMDCardNumbers
	StaffProMD.RnSymptom = input.RnSymptom

	StaffProMDB := mapper.ConvertStructToBSONMap(StaffProMD, nil)

	insPatiProMD := maindbv2.Insertdb(ctx, staffDB, staffDB_profile_model, StaffProMDB)
	if insPatiProMD == `nok` {
		return "nok"
	}

	var StaffPro internalmodel.StaffProfile
	StaffPro.SUID = sUID
	// StaffPro.RsClinicUid = append([]string{}, input.RsClinicUID...)
	StaffPro.RsClinicUID = append([]string{}, UIDC)
	// StaffPro.RsClinicUID = removeDuplicates(append(StaffPro.RsClinicUID, input.RsClinicUID...))
	// StaffPro.SClinicalPractitioner = append([]string{}, input.RsCaregiver...)
	StaffPro.RsCaregiverUID = []string{}
	StaffPro.RsPatientUID = []string{}

	StaffProB := mapper.ConvertStructToBSONMap(StaffPro, nil)

	insStaffPro := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": sUID}, StaffProB)
	if insStaffPro == `nok` {
		return "nok"
	}

	addStafftoClonic := maindbv2.UpdatePushArraycusStr(ctx, clinicDB, clinicDB_profile_model, bson.M{"SUID": UIDC}, sUID, "RsStaffUID")
	if addStafftoClonic == `nok` {
		return "nok"
	}
	var TeleSession internalmodel.TeleSession
	TeleSession.SUID = sUID
	TeleSessionB := mapper.ConvertStructToBSONMap(TeleSession, nil)

	insTeleSession := maindbv2.Insertdb(ctx, staffDB, `TeleSession`, TeleSessionB)
	if insTeleSession == `nok` {
		return "nok"
	}

	return "ok"
}

func Login(ctx context.Context, input model.Login, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string) model.ReturnLogin {

	var output model.ReturnLogin
	t1 := time.Now().Unix()
	dbs1 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SEmail": input.SEmail}, "_id", 1, 0, 0)
	if len(dbs1) == 0 {
		output.SStatus = `db disconnect`
		return output
	}
	inputpass := jwt.ParseToken(input.SPass, 3)
	dbpass := jwt.ParseToken(fmt.Sprintf("%v", dbs1[0][`SPass`]), 3)

	if len(dbs1) == 0 {
		output.SStatus = `no have Email in system`
		return output
	} else {
		// if dbs1[0][`SPass`] == input.SPass {
		if dbpass == inputpass {
			token := jwt.GenerateToken(fmt.Sprintf("%v", dbs1[0][`SUID`])+`<>`+strconv.FormatInt(t1, 10), 7)
			if token == `` {
				output.SStatus = `system error`
				return output
			}
			output.SStatus = `ok`
			output.SSessionUIDs = token
			output.NRoleid = 0
			output.NAccountstatus = int(dbs1[0][`NStatus`].(int32))
			output.NSessionlifetime = 1000

			// dbs2 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": dbs1[0][`SUID`]}, "_id", 1, 0, 0)
			// if len(dbs2) == 0 {
			// 	output.SStatus = `db disconnect`
			// 	return output
			// }

			// fmt.Println(dbs2)
			// fmt.Println(dbs1[0][`RsClinicUID`])
			if len(intolistr(dbs1[0][`RsClinicUID`])) > 0 {

				for i := 0; i < len(intolistr(dbs1[0][`RsClinicUID`])); i++ {
					s := intolistr(dbs1[0][`RsClinicUID`])
					token2 := jwt.GenerateToken(fmt.Sprintf("%v", s[i])+`<>`+strconv.FormatInt(t1, 10), 7)
					if token2 == `` {
						output.SStatus = `system error`
						return output
					}
					output.SSessionUIDc = append(output.SSessionUIDc, token2)
				}

			}

		} else {
			output.SStatus = `incorrect password`
			return output
		}
	}
	return output

}

func GetProfile(ctx context.Context, input model.GetStaffData, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string) model.StaffProfile {
	var output model.StaffProfile

	// t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UIDS := SessionUIDPArS[0]

	dbs1S := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UIDS}, "_id", 1, 0, 0)
	if len(dbs1S) == 0 {
		output.SStatus = `someting wrong?? 1`
		return output
	}

	output.SStatus = `ok`
	output.NStatus = int(dbs1S[0][`NStatus`].(int32))
	output.SBytePhoto = fmt.Sprintf("%v", dbs1S[0][`SBytePhoto`])
	output.SEmail = fmt.Sprintf("%v", dbs1S[0][`SEmail`])
	output.SFirstName = fmt.Sprintf("%v", dbs1S[0][`SFirstName`])
	output.SLastName = fmt.Sprintf("%v", dbs1S[0][`SLastName`])
	output.SNickName = fmt.Sprintf("%v", dbs1S[0][`SNickName`])
	output.NGender = int(dbs1S[0][`NGender`].(int32))
	output.SDateBirth = fmt.Sprintf("%v", dbs1S[0][`SDateBirth`])
	output.SDrOfficialID = fmt.Sprintf("%v", dbs1S[0][`SDrOfficialID`])
	output.SAddress = fmt.Sprintf("%v", dbs1S[0][`SAddress`])
	output.SKwang = fmt.Sprintf("%v", dbs1S[0][`SKwang`])
	output.SKhet = fmt.Sprintf("%v", dbs1S[0][`SKhet`])
	output.SProvince = fmt.Sprintf("%v", dbs1S[0][`SProvince`])
	output.SPost = fmt.Sprintf("%v", dbs1S[0][`SPost`])
	output.SCountry = fmt.Sprintf("%v", dbs1S[0][`SCountry`])
	output.STel = fmt.Sprintf("%v", dbs1S[0][`STel`])
	output.SLineID = fmt.Sprintf("%v", dbs1S[0][`SLineID`])
	output.NTitle = int(dbs1S[0][`NTitle`].(int32))
	output.RnSymptom = (append(intoliin(dbs1S[0][`RnSymptom`]), output.RnSymptom...))

	return output
}

func Editprofile(ctx context.Context, input model.StaffProfileEdit, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string) string {

	// t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	if SessionUIDS == `` {
		return `someting wrong?`
	}
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UID := SessionUIDPArS[0]

	// fmt.Println(UID)
	dbs1 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, "_id", 1, 0, 0)
	if len(dbs1) == 0 {
		return `someting wrong??`
	}

	// Uidtoken := jwt.GenerateToken(UID+`<>`+strconv.FormatInt(t1, 10), 7)
	// if Uidtoken == `` {
	// 	return `system error`
	// }

	var StaffProMDED internalmodel.StaffProfileModelED

	if input.NStatus != nil {
		// StaffProMDED.NStatus = int((dbs1[0][`NStatus`]).(int32))
		StaffProMDED.NStatus = *input.NStatus
	} else {
		StaffProMDED.NStatus = int(dbs1[0][`NStatus`].(int32))
	}

	if input.SBytePhoto != nil {
		// StaffProMDED.SBytePhoto = fmt.Sprintf("%v", dbs1[0][`SBytePhoto`])
		StaffProMDED.SBytePhoto = *input.SBytePhoto
	} else {
		StaffProMDED.SBytePhoto = fmt.Sprintf("%v", dbs1[0][`SBytePhoto`])
	}

	if input.SEmail != nil {
		// StaffProMDED.SEmail = fmt.Sprintf("%v", dbs1[0][`SEmail`])
		StaffProMDED.SEmail = *input.SEmail
	} else {
		StaffProMDED.SEmail = fmt.Sprintf("%v", dbs1[0][`SEmail`])
	}

	if input.SFirstName != nil {
		// StaffProMDED.SFirstName = fmt.Sprintf("%v", dbs1[0][`SFirstName`])
		StaffProMDED.SFirstName = *input.SFirstName
	} else {
		StaffProMDED.SFirstName = fmt.Sprintf("%v", dbs1[0][`SFirstName`])
	}

	if input.SLastName != nil {
		// StaffProMDED.SLastName = fmt.Sprintf("%v", dbs1[0][`SLastName`])
		StaffProMDED.SLastName = *input.SLastName
	} else {
		StaffProMDED.SLastName = fmt.Sprintf("%v", dbs1[0][`SLastName`])
	}

	if input.SNickName != nil {
		// StaffProMDED.SNickName = fmt.Sprintf("%v", dbs1[0][`SNickName`])
		StaffProMDED.SNickName = *input.SNickName
	} else {
		StaffProMDED.SNickName = fmt.Sprintf("%v", dbs1[0][`SNickName`])
	}

	if input.NGender != nil {
		// StaffProMDED.NGender = int((dbs1[0][`NGender`]).(int32))
		StaffProMDED.NGender = *input.NGender
	} else {
		StaffProMDED.NGender = int((dbs1[0][`NGender`]).(int32))
	}

	if input.SDateBirth != nil {
		// StaffProMDED.SDateBirth = fmt.Sprintf("%v", dbs1[0][`SDateBirth`])
		StaffProMDED.SDateBirth = *input.SDateBirth
	} else {
		StaffProMDED.SDateBirth = fmt.Sprintf("%v", dbs1[0][`SDateBirth`])
	}

	if input.SDrOfficialID != nil {
		// StaffProMDED.SDrOfficialID = fmt.Sprintf("%v", dbs1[0][`SDrOfficialID`])
		StaffProMDED.SDrOfficialID = *input.SDrOfficialID
	} else {
		StaffProMDED.SDrOfficialID = fmt.Sprintf("%v", dbs1[0][`SDrOfficialID`])
	}

	if input.SAddress != nil {
		// StaffProMDED.SAddress = fmt.Sprintf("%v", dbs1[0][`SAddress`])
		StaffProMDED.SAddress = *input.SAddress
	} else {
		StaffProMDED.SAddress = fmt.Sprintf("%v", dbs1[0][`SAddress`])
	}

	if input.SKwang != nil {
		// StaffProMDED.SKwang = fmt.Sprintf("%v", dbs1[0][`SKwang`])
		StaffProMDED.SKwang = *input.SKwang
	} else {
		StaffProMDED.SKwang = fmt.Sprintf("%v", dbs1[0][`SKwang`])
	}

	if input.SKhet != nil {
		// StaffProMDED.SKhet = fmt.Sprintf("%v", dbs1[0][`SKhet`])
		StaffProMDED.SKhet = *input.SKhet
	} else {
		StaffProMDED.SKhet = fmt.Sprintf("%v", dbs1[0][`SKhet`])
	}

	if input.SProvince != nil {
		// StaffProMDED.SProvince = fmt.Sprintf("%v", dbs1[0][`SProvince`])
		StaffProMDED.SProvince = *input.SProvince
	} else {
		StaffProMDED.SProvince = fmt.Sprintf("%v", dbs1[0][`SProvince`])
	}

	if input.SPost != nil {
		// StaffProMDED.SPost = fmt.Sprintf("%v", dbs1[0][`SPost`])
		StaffProMDED.SPost = *input.SPost
	} else {
		StaffProMDED.SPost = fmt.Sprintf("%v", dbs1[0][`SPost`])
	}

	if input.SCountry != nil {
		// StaffProMDED.SCountry = fmt.Sprintf("%v", dbs1[0][`SCountry`])
		StaffProMDED.SCountry = *input.SCountry
	} else {
		StaffProMDED.SCountry = fmt.Sprintf("%v", dbs1[0][`SCountry`])
	}

	if input.STel != nil {
		// StaffProMDED.STel = fmt.Sprintf("%v", dbs1[0][`STel`])
		StaffProMDED.STel = *input.STel
	} else {
		StaffProMDED.STel = fmt.Sprintf("%v", dbs1[0][`STel`])
	}

	if input.SLineID != nil {
		// StaffProMDED.SLineID = fmt.Sprintf("%v", dbs1[0][`SLineID`])
		StaffProMDED.SLineID = *input.SLineID
	} else {
		StaffProMDED.SLineID = fmt.Sprintf("%v", dbs1[0][`SLineID`])
	}

	if input.NImgAvatarID != nil {
		StaffProMDED.NImgAvatarID = *input.NImgAvatarID
	} else {
		StaffProMDED.NImgAvatarID = int(dbs1[0][`NImgAvatarID`].(int32))
	}

	if input.NTitle != nil {
		StaffProMDED.NTitle = *input.NTitle
	} else {
		StaffProMDED.NTitle = int(dbs1[0][`NTitle`].(int32))
	}

	if input.SMDCardNumbers != nil {
		StaffProMDED.SMDCardNumbers = *input.SMDCardNumbers
	} else {
		StaffProMDED.SMDCardNumbers = fmt.Sprintf("%v", dbs1[0][`SMDCardNumbers`])
	}

	if input.RnSymptom != nil {
		StaffProMDED.RnSymptom = input.RnSymptom
	} else {
		StaffProMDED.RnSymptom = (append(intoliin(dbs1[0][`RnSymptom`]), StaffProMDED.RnSymptom...))
	}

	StaffProMDEDB := mapper.ConvertStructToBSONMap(StaffProMDED, nil)

	insertprofilemodel := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, StaffProMDEDB)
	if insertprofilemodel == `nok` {
		return `database have some problem`
	}

	// // if len(input.RsCaregiver > 0 {
	// dbs2 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUid": UID}, "_id", 1, 0, 0)
	// if len(dbs2) == 0 {
	// 	return `someting wrong??`
	// }

	// var insertprofilestuc internalmodel.StaffProfileED
	// insertprofilestuc.RsCaregiver = removeDuplicates(append(intolistr(dbs2[0][`RsCaregiver`]), input.RsCaregiver...))
	// insertprofilestuc.RsCaregiverUID = input.RsCaregiver
	// insertprofilestuc.RsClinicUid = intolistr(dbs2[0][`RsClinicUid`])
	// insertprofilestuc.RsClinicUID = input.RsClinicUID
	// StaffProEDB := mapper.ConvertStructToBSONMap(insertprofilestuc, nil)

	// insertprofile := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUid": Uid}, StaffProEDB)
	// if insertprofile == `nok` {
	// 	return `database have some problem`
	// }
	// // }

	return `ok`

}

func ChangeEmailtg(ctx context.Context, input model.ChangeEmailIn, staffDB string, staffDB_profile_model string) string {

	// t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	if SessionUIDS == `` {
		return `someting wrong?`
	}
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UID := SessionUIDPArS[0]

	// fmt.Println(UID)
	dbs1 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, "_id", 1, 0, 0)
	if len(dbs1) == 0 {
		return `someting wrong??`
	}

	inputpass := jwt.ParseToken(input.SPassword, 3)
	dbpass := jwt.ParseToken(fmt.Sprintf("%v", dbs1[0][`SPass`]), 3)

	if dbpass == inputpass {

		insertprofilemodel := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, bson.M{"SEmail": input.SNewEmail})
		if insertprofilemodel == `nok` {
			return `database have some problem`
		}

	} else {
		return `Password Incorrect`
	}

	return `ok`
}

func ChangePWDtg(ctx context.Context, input model.ChangePwdin, staffDB string, staffDB_profile_model string) string {

	// t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	if SessionUIDS == `` {
		return `someting wrong?`
	}
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UID := SessionUIDPArS[0]

	// fmt.Println(UID)
	dbs1 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, "_id", 1, 0, 0)
	if len(dbs1) == 0 {
		return `someting wrong??`
	}

	inputpass := jwt.ParseToken(input.SPassword, 3)
	dbpass := jwt.ParseToken(fmt.Sprintf("%v", dbs1[0][`SPass`]), 3)

	if dbpass == inputpass {

		insertprofilemodel := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, bson.M{"SPass": input.SNewPassword})
		if insertprofilemodel == `nok` {
			return `database have some problem`
		}

	} else {
		return `Password Incorrect`
	}
	return `ok`
}

func ChangePHONEtg(ctx context.Context, input model.ChangePhoneIn, staffDB string, staffDB_profile_model string) string {

	// t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	if SessionUIDS == `` {
		return `someting wrong?`
	}
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UID := SessionUIDPArS[0]

	// fmt.Println(UID)
	dbs1 := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, "_id", 1, 0, 0)
	if len(dbs1) == 0 {
		return `someting wrong??`
	}

	inputpass := jwt.ParseToken(input.SPassword, 3)
	dbpass := jwt.ParseToken(fmt.Sprintf("%v", dbs1[0][`SPass`]), 3)

	if dbpass == inputpass {

		insertprofilemodel := maindbv2.UpdateArchive(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UID}, bson.M{"STel": input.SNewPhone})
		if insertprofilemodel == `nok` {
			return `database have some problem`
		}

	} else {
		return `Password Incorrect`
	}
	return `ok`
}

func CreatePatientProfile(ctx context.Context, input model.PatientProfileCreate, authDB string, authDB_profile_nouse string, authDB_profile_model string) string {

	// t1 := time.Now().Unix()
	// t2 := time.Now().UnixNano()
	// sUid := `p` + "-" + strconv.FormatInt(t1, 16) + strconv.FormatInt(t2, 16) + RandStringRunes(5)
	// // fmt.Println(sUid)

	// var PatiProMD internalmodel.PatientProfileModel

	// PatiProMD.SUid = sUid
	// PatiProMD.NStatus = input.NStatus
	// PatiProMD.SBytePhoto = input.SBytePhoto
	// PatiProMD.NImgAvatarID = input.NImgAvatarID
	// PatiProMD.SEmail = input.SEmail
	// PatiProMD.SFirstName = input.SFirstName
	// PatiProMD.SLastName = input.SLastName
	// PatiProMD.SNickName = input.SNickName
	// PatiProMD.NGender = input.NGender
	// PatiProMD.SDateBirth = input.SDateBirth
	// PatiProMD.SNationID = input.SNationID
	// PatiProMD.SAddress = input.SAddress
	// PatiProMD.SKwang = input.SKwang
	// PatiProMD.SKhet = input.SKhet
	// PatiProMD.SProvince = input.SProvince
	// PatiProMD.SPost = input.SPost
	// PatiProMD.SCountry = input.SCountry
	// PatiProMD.STel = input.STel
	// PatiProMD.SLineID = input.SLineID
	// PatiProMD.NWhoPay = input.NWhoPay
	// PatiProMD.NtimestampFront = input.NtimestampFront
	// PatiProMD.Ntimestamp = int(t1)

	// PatiProMDB := mapper.ConvertStructToBSONMap(PatiProMD, nil)

	// insPatiProMD := maindbv2.Insertdb(ctx, authDB, authDB_profile_model, PatiProMDB)
	// if insPatiProMD == `nok` {
	// 	return "nok"
	// }

	// //------------------------------------------------------

	// var PatiPro internalmodel.PatientProfile

	// PatiPro.SUid = sUid
	// // PatiPro.RnSymptom = input.RnSymptom
	// PatiPro.RsCaregiverUID = input.RsCaregiverUID
	// PatiPro.RsStaffUID = input.RsStaffUID
	// PatiPro.RsClinicUID = input.RsClinicUID
	// // PatiPro.RsDoctorUID = input.RsDoctorUID

	// PatiProB := mapper.ConvertStructToBSONMap(PatiPro, nil)
	// insPatiPro := maindbv2.UpdateArchive(ctx, authDB, authDB_profile_model, bson.M{"SUid": sUid}, PatiProB)
	// if insPatiPro == `nok` {
	// 	return "nok"
	// }

	// return "ok"
	return `unavailable`
}

func EditPatientProfile(ctx context.Context, input model.PatientProfileEdit, authDB string, authDB_profile_nouse string, authDB_profile_model string) string {

	// // fmt.Print(input.SSessionUIDP)
	// t1 := time.Now().Unix()
	// SessionUIDP := jwt.ParseToken(input.SSessionUIDP, 7)
	// if SessionUIDP == `` {
	// 	return `someting wrong?`
	// }
	// SessionUIDPAr := strings.Split(SessionUIDP, `<>`)
	// Uid := SessionUIDPAr[0]

	// fmt.Println(Uid)
	// dbs1 := maindbv2.Finddb(ctx, authDB, authDB_profile_model, bson.M{"SUid": Uid}, "_id", 1, 0, 0)
	// if len(dbs1) == 0 {

	// 	return `someting wrong??====`
	// }

	// Uidtoken := jwt.GenerateToken(Uid+`<>`+strconv.FormatInt(t1, 10), 7)
	// if Uidtoken == `` {
	// 	return `system error`
	// }

	// var PatiProMDED internalmodel.PatientProfileModelEdir

	// if input.NtimestampFront == 0 {
	// 	// PatiProMDED.NtimestampFront = int((dbs1[0][`NtimestampFront`]).(int32))
	// 	PatiProMDED.NtimestampFront = input.NtimestampFront
	// } else {
	// 	PatiProMDED.NtimestampFront = input.NtimestampFront
	// }

	// if input.NStatus == 0 {
	// 	// PatiProMDED.NStatus = int((dbs1[0][`NStatus`]).(int32))
	// 	PatiProMDED.NStatus = input.NStatus
	// } else {
	// 	PatiProMDED.NStatus = input.NStatus
	// }

	// if input.SBytePhoto == `` {
	// 	// PatiProMDED.SBytePhoto = fmt.Sprintf("%v", dbs1[0][`SBytePhoto`])
	// 	PatiProMDED.SBytePhoto = input.SBytePhoto
	// } else {
	// 	PatiProMDED.SBytePhoto = input.SBytePhoto
	// }

	// if input.NImgAvatarID == 0 {
	// 	// PatiProMDED.NImgAvatarID = int((dbs1[0][`NImgAvatarID`]).(int32))
	// 	PatiProMDED.NImgAvatarID = input.NImgAvatarID
	// } else {
	// 	PatiProMDED.NImgAvatarID = input.NImgAvatarID
	// }

	// if input.SEmail == `` {
	// 	// PatiProMDED.SEmail = fmt.Sprintf("%v", dbs1[0][`SEmail`])
	// 	PatiProMDED.SEmail = input.SEmail
	// } else {
	// 	PatiProMDED.SEmail = input.SEmail
	// }

	// if input.SFirstName == `` {
	// 	// PatiProMDED.SFirstName = fmt.Sprintf("%v", dbs1[0][`SFirstName`])
	// 	PatiProMDED.SFirstName = input.SFirstName
	// } else {
	// 	PatiProMDED.SFirstName = input.SFirstName
	// }

	// if input.SLastName == `` {
	// 	// PatiProMDED.SLastName = fmt.Sprintf("%v", dbs1[0][`SLastName`])
	// 	PatiProMDED.SLastName = input.SLastName
	// } else {
	// 	PatiProMDED.SLastName = input.SLastName
	// }

	// if input.SNickName == `` {
	// 	// PatiProMDED.SNickName = fmt.Sprintf("%v", dbs1[0][`SNickName`])
	// 	PatiProMDED.SNickName = input.SNickName
	// } else {
	// 	PatiProMDED.SNickName = input.SNickName
	// }

	// if input.NGender == 0 {
	// 	// PatiProMDED.NGender = int((dbs1[0][`NGender`]).(int32))
	// 	PatiProMDED.NGender = input.NGender
	// } else {
	// 	PatiProMDED.NGender = input.NGender
	// }

	// if input.SDateBirth == `` {
	// 	// PatiProMDED.SDateBirth = fmt.Sprintf("%v", dbs1[0][`SDateBirth`])
	// 	PatiProMDED.SDateBirth = input.SDateBirth
	// } else {
	// 	PatiProMDED.SDateBirth = input.SDateBirth
	// }

	// if input.SNationID == `` {
	// 	// PatiProMDED.SNationID = fmt.Sprintf("%v", dbs1[0][`SNationID`])
	// 	PatiProMDED.SNationID = input.SNationID
	// } else {
	// 	PatiProMDED.SNationID = input.SNationID
	// }

	// if input.SAddress == `` {
	// 	// PatiProMDED.SAddress = fmt.Sprintf("%v", dbs1[0][`SAddress`])
	// 	PatiProMDED.SAddress = input.SAddress
	// } else {
	// 	PatiProMDED.SAddress = input.SAddress
	// }

	// if input.SKwang == `` {
	// 	// PatiProMDED.SKwang = fmt.Sprintf("%v", dbs1[0][`SKwang`])
	// 	PatiProMDED.SKwang = input.SKwang
	// } else {
	// 	PatiProMDED.SKwang = input.SKwang
	// }

	// if input.SKhet == `` {
	// 	// PatiProMDED.SKhet = fmt.Sprintf("%v", dbs1[0][`SKhet`])
	// 	PatiProMDED.SKhet = input.SKhet
	// } else {
	// 	PatiProMDED.SKhet = input.SKhet
	// }

	// if input.SProvince == `` {
	// 	// PatiProMDED.SProvince = fmt.Sprintf("%v", dbs1[0][`SProvince`])
	// 	PatiProMDED.SProvince = input.SProvince
	// } else {
	// 	PatiProMDED.SProvince = input.SProvince
	// }

	// if input.SPost == `` {
	// 	// PatiProMDED.SPost = fmt.Sprintf("%v", dbs1[0][`SPost`])
	// 	PatiProMDED.SPost = input.SPost
	// } else {
	// 	PatiProMDED.SPost = input.SPost
	// }

	// if input.SCountry == `` {
	// 	// PatiProMDED.SCountry = fmt.Sprintf("%v", dbs1[0][`SCountry`])
	// 	PatiProMDED.SCountry = input.SCountry
	// } else {
	// 	PatiProMDED.SCountry = input.SCountry
	// }

	// if input.STel == `` {
	// 	// PatiProMDED.STel = fmt.Sprintf("%v", dbs1[0][`STel`])
	// 	PatiProMDED.STel = input.STel
	// } else {
	// 	PatiProMDED.STel = input.STel
	// }

	// if input.SLineID == `` {
	// 	// PatiProMDED.SLineID = fmt.Sprintf("%v", dbs1[0][`SLineID`])
	// 	PatiProMDED.SLineID = input.SLineID
	// } else {
	// 	PatiProMDED.SLineID = input.SLineID
	// }

	// if input.NWhoPay == 0 {
	// 	// PatiProMDED.NWhoPay = int((dbs1[0][`NWhoPay`]).(int32))
	// 	PatiProMDED.NWhoPay = input.NWhoPay
	// } else {
	// 	PatiProMDED.NWhoPay = input.NWhoPay
	// }

	// PatiProMDED.Ntimestamp = int(t1)

	// PatiProMDEDB := mapper.ConvertStructToBSONMap(PatiProMDED, nil)

	// // fmt.Println(PatiProMDEDB)

	// insertprofilemodel := maindbv2.UpdateArchive(ctx, authDB, authDB_profile_model, bson.M{"SUid": Uid}, PatiProMDEDB)
	// if insertprofilemodel == `nok` {
	// 	return `database have some problem`
	// }

	// // if input.IsUpdateonlyProfileModel == true {
	// dbs2 := maindbv2.Finddb(ctx, authDB, authDB_profile_model, bson.M{"SUid": Uid}, "_id", 1, 0, 0)
	// if len(dbs2) == 0 {

	// 	return `someting wrong??`
	// }

	// var insertprofilestuc internalmodel.PatientProfileEdit

	// // insertprofilestuc.RnSymptom = removeDuplicatesint(append(intoliin(dbs2[0][`RnSymptom`]), input.RnSymptom...))
	// insertprofilestuc.RnSymptom = input.RnSymptom
	// // insertprofilestuc.RsCaregiverUID = removeDuplicates(append(intolistr(dbs2[0][`RsCaregiverUID`]), input.RsCaregiverUID...))
	// insertprofilestuc.RsCaregiverUID = input.RsCaregiverUID
	// // insertprofilestuc.RsStaffUID = removeDuplicates(append(intolistr(dbs2[0][`RsStaffUID`]), input.RsStaffUID...))
	// insertprofilestuc.RsStaffUID = input.RsStaffUID
	// // insertprofilestuc.RsClinicUID = removeDuplicates(append(intolistr(dbs2[0][`RsClinicUID`]), input.RsClinicUID...))
	// insertprofilestuc.RsClinicUID = input.RsClinicUID
	// // insertprofilestuc.RsDoctorUID = removeDuplicates(append(intolistr(dbs2[0][`RsDoctorUID`]), input.RsDoctorUID...))
	// insertprofilestuc.RsDoctorUID = input.RsDoctorUID
	// insertprofilestuc.Ntimestamp = int(t1)
	// PatiProEDB := mapper.ConvertStructToBSONMap(insertprofilestuc, nil)

	// insertprofile := maindbv2.UpdateArchive(ctx, authDB, authDB_profile_model, bson.M{"SUid": Uid}, PatiProEDB)
	// if insertprofile == `nok` {
	// 	return `database have some problem`
	// }
	// // }

	// return `ok`

	return `unavailable`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		b[i] = letterRunes[r1.Intn(len(letterRunes))]
	}
	return string(b)
}

func intoliin(t interface{}) []int {
	var output []int
	if t != nil {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				// fmt.Println(s.Index(i))
				output = append(output, int(s.Index(i).Interface().(int32)))
			}
		}
	} else {
		return output
	}

	return output
}

func intolistr(t interface{}) []string {
	var output []string
	if t != nil {
		switch reflect.TypeOf(t).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				// fmt.Println(s.Index(i))
				output = append(output, fmt.Sprintf("%v", s.Index(i).Interface()))
			}
		}
	} else {
		return output
	}

	return output
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	if strList != nil {
		for _, item := range strList {
			fmt.Println(item)
			if contains(list, item) == false {
				list = append(list, item)
			}
		}
	} else {
		return list
	}

	return list
}
