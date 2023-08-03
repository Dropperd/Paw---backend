package service

import (
	"websiteapi/entity"
	"websiteapi/repository"
)

func GetFeedbacks(imageId uint64, userId uint64) []entity.Feedback {
	if userId == 0 {
		return repository.GetFeedbacksFromImageId(imageId)
	} else {

		if repository.IsClinical(userId) {
			return repository.GetFeedbacksFromImageIdUserId(imageId, userId)
		}
	}
	return nil
}

func UpdateFeedback(feedback entity.Feedback) entity.Feedback {
	return repository.CreateUpdateFeedback(feedback)
}

func CreateFeedback(feedback entity.Feedback) entity.Feedback {
	return repository.CreateUpdateFeedback(feedback)
}

func DeleteFeedback(feedback_id uint64) {
	repository.DeleteFeedback(feedback_id)
}

func GetFeedbacksCount(imageId uint64) uint64 {
	return repository.GetFeedbacksFromImageIdCount(imageId)
}
