package service

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/attendance"
	"github.com/google/uuid"
)

func GetAttendances() ([]attendance.Attendance, error) {
	return attendance.GetAttendanceAll()
}

func CreateAttendance(v attendance.Attendance) (attendance.Attendance, error) {
	token := uuid.New().String()
	v = attendance.New(token, v.ExpireTime)
	if err := attendance.Insert(v); err != nil {
		return attendance.Attendance{}, err
	}
	return v, nil
}

func EditAttendance(id string, v attendance.Attendance) (result attendance.Attendance, err error) {
	result, err = attendance.GetAttendance(attendance.Attendance{Id: id})
	if err != nil {
		return
	}

	result.Update(v)
	return
}

func DelAttendance(id string) (result attendance.Attendance, err error) {
	result, err = attendance.GetAttendance(attendance.Attendance{Id: id})
	if err != nil {
		if err == attendance.ErrNotExistAttendance {
			err = attendance.ErrAlreadyDeleted
			return
		}

		return
	}

	err = attendance.Delete(result)
	return
}
