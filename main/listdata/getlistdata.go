package getlisdata

import (
	"context"
	"encoding/json"
	"fmt"
	"go-resfull/graph/model"
	"go-resfull/jwt"
	"go-resfull/mongo/maindbv2"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func StaffListData(ctx context.Context, input model.GetListData, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string, clinicDB string, clinicDB_profile_nouse string, clinicDB_profile_model string) model.GetStaffListReturn {
	var output model.GetStaffListReturn

	t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UIDS := SessionUIDPArS[0]

	SessionUIDC := jwt.ParseToken(input.SSessionUIDc, 7)
	SessionUIDPArC := strings.Split(SessionUIDC, `<>`)
	UIDC := SessionUIDPArC[0]

	fmt.Print(UIDS)

	dbs1S := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UIDS}, "_id", 1, 0, 0)
	if len(dbs1S) == 0 {
		output.SStatus = `someting wrong?? 1`
		return output
	}
	// fmt.Println(UIDS)
	// fmt.Println(UIDC)

	dbs1C := maindbv2.Finddb(ctx, clinicDB, clinicDB_profile_model, bson.M{"SUID": UIDC}, "_id", 1, 0, 0)
	if len(dbs1C) == 0 {
		output.SStatus = `someting wrong?? 2--> `
		return output
	}

	//MVP_only

	listStaff := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{}, "_id", 1, 0, 0)
	if len(listStaff) == 0 {
		output.SStatus = `someting wrong?? 3`
		return output
	}

	var rStaffListbuffer []*model.StaffList

	for i := 0; i < len(listStaff); i++ {
		var StaffBuffer model.StaffList
		StaffBuffer.SUID = jwt.GenerateToken(fmt.Sprintf("%v", listStaff[i][`SUID`])+`<>`+strconv.FormatInt(t1, 10), 7)
		StaffBuffer.SBytePhoto = fmt.Sprintf("%v", listStaff[i][`SBytePhoto`])
		StaffBuffer.NImgAvatarID = int(listStaff[0][`NImgAvatarID`].(int32))
		StaffBuffer.NTitle = int(listStaff[0][`NTitle`].(int32))
		StaffBuffer.SFirstName = fmt.Sprintf("%v", listStaff[i][`SFirstName`])
		StaffBuffer.SLastName = fmt.Sprintf("%v", listStaff[i][`SLastName`])
		StaffBuffer.SNickName = fmt.Sprintf("%v", listStaff[i][`SNickName`])
		StaffBuffer.NGender = int(listStaff[0][`NGender`].(int32))
		StaffBuffer.SMDCardNumbers = fmt.Sprintf("%v", listStaff[i][`SMDCardNumbers`])
		// StaffBuffer.SClinicalPractitioner = "456"
		// StaffBuffer.RsCaregiverUID = []string{}
		// StaffBufferlistStaff := intolistr(listStaff[i][`RsCaregiverUID`])

		// for j := 0; j < len(StaffBufferlistStaff); j++ {
		// 	SSessionUIDSC := jwt.GenerateToken(fmt.Sprintf("%v", StaffBufferlistStaff[j])+`<>`+strconv.FormatInt(t1, 10), 7)
		// 	StaffBuffer.RsCaregiverUID = append(StaffBuffer.RsCaregiverUID, SSessionUIDSC)
		// }

		StaffBuffer.RnSymptom = intoliin(listStaff[i][`RnSymptom`])
		// StaffBufferlistStaff := intolistr(listStaff[i][`RnSymptom`])

		rStaffListbuffer = append(rStaffListbuffer, &StaffBuffer)
	}

	output.SStatus = `ok`
	output.RStaffList = rStaffListbuffer

	return output
}

func PatientListData(ctx context.Context, input model.GetListData, staffDB string, staffDB_profile_nouse string, staffDB_profile_model string, clinicDB string, clinicDB_profile string, clinicDB_profile_model string, authDB string, authDB_profile string, authDB_profile_model string) model.GetPatientListReturn {
	var output model.GetPatientListReturn

	t1 := time.Now().Unix()
	SessionUIDS := jwt.ParseToken(input.SSessionUIDs, 7)
	SessionUIDPArS := strings.Split(SessionUIDS, `<>`)
	UIDS := SessionUIDPArS[0]

	fmt.Println(UIDS)

	SessionUIDC := jwt.ParseToken(input.SSessionUIDc, 7)
	SessionUIDPArC := strings.Split(SessionUIDC, `<>`)
	UIDC := SessionUIDPArC[0]

	dbs1S := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{"SUID": UIDS}, "_id", 1, 0, 0)
	if len(dbs1S) == 0 {
		fmt.Print(dbs1S)
		output.SStatus = `someting wrong?? 1`
		return output
	}

	// fmt.Println(UIDC)

	dbs1C := maindbv2.Finddb(ctx, clinicDB, clinicDB_profile_model, bson.M{"SUID": UIDC}, "_id", 1, 0, 0)
	if len(dbs1C) == 0 {
		output.SStatus = `someting wrong?? 2`
		return output
	}

	listPatient := maindbv2.Finddb(ctx, authDB, authDB_profile_model, bson.M{}, "_id", 1, 0, 0)
	if len(listPatient) == 0 {
		output.SStatus = `someting wrong?? 3`
		return output
	}

	var rPatientListbuffer []*model.PatientList

	dbs1Sall := maindbv2.Finddb(ctx, staffDB, staffDB_profile_model, bson.M{}, "_id", 1, 0, 0)
	if len(dbs1Sall) == 0 {
		output.SStatus = `someting wrong?? 1`
		return output
	}

	// fmt.Println(len(dbs1Sall))

	for i := 0; i < len(listPatient); i++ {
		var PatientBuffer model.PatientList
		PatientBuffer.SUID = jwt.GenerateToken(fmt.Sprintf("%v", listPatient[i][`SUID`])+`<>`+strconv.FormatInt(t1, 10), 7)
		PatientBuffer.NImgAvatarID = int(listPatient[i][`NImgAvatarID`].(int32))
		PatientBuffer.SFirstName = fmt.Sprintf("%v", listPatient[i][`SFirstName`])
		PatientBuffer.SLastName = fmt.Sprintf("%v", listPatient[i][`SLastName`])
		PatientBuffer.SNickName = fmt.Sprintf("%v", listPatient[i][`SNickName`])
		PatientBuffer.NGender = int(listPatient[i][`NGender`].(int32))
		PrsDoctorUIDBuffer := intolistBsonM(listPatient[i][`PrsDoctorUID`])
		PatientBuffer.RsDoctorUID = []*model.MsDoctorReturn{}
		// fmt.Println(PrsDoctorUIDBuffer[0][`NSymptom`])
		// fmt.Printf("var1 = %T\n", PrsDoctorUIDBuffer[0][`NSymptom`])

		for j := 0; j < len(PrsDoctorUIDBuffer); j++ {
			var MsDoctorReturn model.MsDoctorReturn
			MsDoctorReturn.NSymptom = int(PrsDoctorUIDBuffer[j][`NSymptom`].(float64))
			// MsDoctorReturn.SSessionUIDs = jwt.GenerateToken(fmt.Sprintf("%v", PrsDoctorUIDBuffer[j][`SUIDS`])+`<>`+strconv.FormatInt(t1, 10), 7)
			for k := 0; k < len(dbs1Sall); k++ {
				if dbs1Sall[k][`SUID`] == fmt.Sprintf("%v", PrsDoctorUIDBuffer[j][`SUIDS`]) {
					MsDoctorReturn.SSessionUIDs = fmt.Sprintf("%v", dbs1Sall[k][`SFirstName`]) + ` ` + fmt.Sprintf("%v", dbs1Sall[k][`SLastName`])
					break
				} else {
					MsDoctorReturn.SSessionUIDs = `-`
				}
			}

			PatientBuffer.RsDoctorUID = append(PatientBuffer.RsDoctorUID, &MsDoctorReturn)
		}

		// var MsDoctorReturn model.MsDoctor
		// MsDoctorReturn.NSymptom = int(primitivetoBsonM(listPatient[i][`PrsDoctorUID`])[`NSymptom`].(int32))
		// PatientBuffer.RsDoctorUID

		PatientBuffer.RTeleSession = []string{}

		// fmt.Println(PatientBuffer)
		rPatientListbuffer = append(rPatientListbuffer, &PatientBuffer)
	}

	output.SStatus = `ok`
	output.RPatientList = rPatientListbuffer

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

func intolistBsonM(t interface{}) []bson.M {
	var output []bson.M
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			// fmt.Println(s.Index(i))
			//------------------------------
			bs, err := json.Marshal(s.Index(i).Interface())
			if err != nil {
				panic(err)
			}
			var o bson.M
			if err := json.Unmarshal(bs, &o); err != nil {
				panic(err)
			}
			//------------------------------
			output = append(output, o)
		}
	}
	return output
}

func primitivetoBsonM(input interface{}, my string) []bson.M {
	bs, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	var o bson.M
	if err := json.Unmarshal(bs, &o); err != nil {
		panic(err)
	}

	if o[my] == nil {
		var output []bson.M
		return output

	} else {
		output := intolistBsonM(o[my])
		return output
	}

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
