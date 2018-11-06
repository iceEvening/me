package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/iceEvening/me/service/model/tables"
)

//SignUp model func
func (m *Model) SignUp(email string, nickname string, password string) error {
	passwd := md5.Sum([]byte(password))

	user := tables.User{
		Email:    email,
		Nickname: nickname,
		Passwd:   hex.EncodeToString(passwd[:]),
		UserRole: 1,
		Birthday: time.Now(),
	}
	err := m.DB.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

//SignIn model func
func (m *Model) SignIn(email string, password string) (*tables.User, error) {
	passwd := md5.Sum([]byte(password))

	user := tables.User{
		Email:  email,
		Passwd: hex.EncodeToString(passwd[:]),
	}
	err := m.DB.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//EditCareers model func
func (m *Model) EditCareers(
	UID uint,
	ID uint,
	company string,
	department string,
	post string,
	rank string,
	description string,
	start time.Time,
	end time.Time,
	isCurrent bool,
) error {
	job := tables.Job{
		UserID:     UID,
		Company:    company,
		Department: department,
		Post:       post,
		JobLevel:   rank,
		JobDesc:    description,
		StartTime:  start,
		EndTime:    end,
		IsCurrent:  isCurrent,
	}

	if ID != 0 {
		job.ID = ID
		err := m.DB.Save(&job).Error
		if err != nil {
			return err
		}
	} else {
		err := m.DB.Create(&job).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//DelCareer model func
func (m *Model) DelCareer(userID uint, eduID uint) error {
	job := tables.Job{
		Model:  gorm.Model{ID: eduID},
		UserID: userID,
	}
	err := m.DB.Delete(&job).Error
	if err != nil {
		return err
	}
	return nil
}

//GetCareers model func
func (m *Model) GetCareers(ID int64) (*[]tables.Job, error) {
	var jobs []tables.Job
	err := m.DB.Where("user_id = ?", ID).Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	return &jobs, nil
}

//EditEducations model func
func (m *Model) EditEducations(
	UID uint,
	ID uint,
	school string,
	major string,
	degree string,
	start time.Time,
	end time.Time,
	isCurrent bool,
) error {
	education := tables.Education{
		UserID:    UID,
		School:    school,
		Major:     major,
		Degree:    degree,
		StartTime: start,
		EndTime:   end,
		IsCurrent: isCurrent,
	}

	if ID != 0 {
		education.ID = ID
		err := m.DB.Save(&education).Error
		if err != nil {
			return err
		}
	} else {
		err := m.DB.Create(&education).Error
		if err != nil {
			return err
		}
	}
	return nil
}

//DelEducation model func
func (m *Model) DelEducation(userID uint, eduID uint) error {
	education := tables.Education{
		Model:  gorm.Model{ID: eduID},
		UserID: userID,
	}
	err := m.DB.Delete(&education).Error
	if err != nil {
		return err
	}
	return nil
}

//GetEducations model func
func (m *Model) GetEducations(ID int64) (*[]tables.Education, error) {
	var educations []tables.Education
	err := m.DB.Where("user_id = ?", ID).Find(&educations).Error
	if err != nil {
		return nil, err
	}
	return &educations, nil
}

//EditUser model func
func (m *Model) EditUser(
	ID uint,
	img []byte,
	nickname string,
	me string,
	name string,
	gender uint,
	email string,
	birthday time.Time,
	hometown string,
	city string,
) error {
	user := tables.User{
		Model:    gorm.Model{ID: ID},
		Img:      img,
		Nickname: nickname,
		Me:       me,
		Name:     name,
		Gender:   gender,
		Email:    email,
		Birthday: birthday,
		Hometown: hometown,
		City:     city,
	}

	err := m.DB.Model(&user).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

//GetUser model func
func (m *Model) GetUser(ID uint) (*tables.User, error) {
	user := tables.User{
		Model: gorm.Model{ID: ID},
	}
	err := m.DB.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
