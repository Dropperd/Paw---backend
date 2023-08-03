package repository

import (
	"time"
	"websiteapi/config"
	"websiteapi/entity"
)

func GetFeedbacksFromImageId(imageId uint64) []entity.Feedback {
	var feedbacks []entity.Feedback

	query := config.Db.Table("image_feedback").
		//Joins("JOIN image_clinicals ON image_feedback.id_clinical = image_clinicals.clinical_id").
		Where("image_feedback.image_id = ?", imageId).
		//Limit(1).
		Scan(&feedbacks)

	if query.Error != nil {
		return nil
	}
	//config.Db.Raw("SELECT p.id, p.user_id, p.feedback, p.id_clinical, p.image_id from image_feedback p , image_clinicals c WHERE p.id_clinical = c.clinical_id AND p.image_id = ?", imageId).Scan(&feedbacks)
	return feedbacks
}

func GetFeedbacksFromImageIdUserId(imageId uint64, userId uint64) []entity.Feedback {
	var feedbacks []entity.Feedback
	config.Db.Raw("SELECT p.id, p.user_id, p.feedback, p.id_clinical from image_feedback p , image_clinicals c WHERE p.id_clinical = c.clinical_id AND p.user_id = c.user_id AND p.image_id = ? AND p.clinical_id = ?", imageId, userId).Scan(&feedbacks)
	return feedbacks
}

func CreateUpdateFeedback(feedback entity.Feedback) entity.Feedback {
	var existingFeedback entity.Feedback
	config.Db.Raw("SELECT * FROM image_feedback WHERE image_id = ? AND id_clinical = ?", feedback.ImageId, feedback.IdClinical).Scan(&existingFeedback)
	if existingFeedback.ID != 0 {
		feedback.Updated_At = time.Now().Format("2006-01-02")
		config.Db.Table("image_feedback").Model(&existingFeedback).UpdateColumns(feedback)
	} else {
		feedback.Added_At = time.Now().Format("2006-01-02")
		feedback.Updated_At = time.Now().Format("2006-01-02")
		feedback.ID = existingFeedback.ID
		config.Db.Table("image_feedback").Save(&feedback)
	}

	return feedback
}

// DeleteFeedback delete feedback
func DeleteFeedback(feedbackId uint64) {
	config.Db.Delete(&feedbackId)
}

func GetFeedbacksFromImageIdCount(imageId uint64) uint64 {
	var count uint64
	config.Db.Raw("SELECT COUNT(*) from image_feedback where image_id = ?", imageId).Scan(&count)
	return count
}
