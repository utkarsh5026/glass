package services

import (
	"errors"
	"fmt"
	"server/app/models"

	"gorm.io/gorm"
)

type EnrollmentService struct {
	db *gorm.DB
}

func NewEnrollmentService(db *gorm.DB) *EnrollmentService {
	return &EnrollmentService{db: db}
}

// JoinCourseByCode enrolls a user in a course using an invitation code.
// Initially it sets the enrollment status to pending.
//
// Parameters:
//   - userID: The ID of the user to enroll.
//   - invitationCode: The invitation code for the course.
//   - role: The role the user will have in the course.
//
// Returns:
//   - error: An error if the enrollment process fails, nil otherwise.
func (s *EnrollmentService) JoinCourseByCode(userID uint, invitationCode string, role models.Role) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var course models.Course
		if err := s.db.Where("invitation_code = ?", invitationCode).First(&course).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("invalid invitation code")
			}
			return err
		}

		enrollment := models.Enrollment{
			UserID:   userID,
			CourseID: course.ID,
			Role:     role,
			Status:   models.EnrollmentStatusPending,
		}

		if err := s.db.Create(&enrollment).Error; err != nil {
			return CreateEntityFailure(err)
		}

		return nil
	})
}

// ApproveEnrolment approves a pending enrollment request for a course.
//
// Parameters:
//   - adminId: The ID of the admin (course creator) approving the enrollment.
//   - enrollmentId: The ID of the enrollment to be approved.
//
// Returns:
//   - error: An error if the approval process fails, nil otherwise.
func (s *EnrollmentService) ApproveEnrolment(adminId, enrollmentId uint) error {
	if adminId == 0 || enrollmentId == 0 {
		return errors.New("invalid input: adminId and enrollmentId must be non-zero")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		var enrollment models.Enrollment
		if err := tx.Preload(models.CoursesTable).Where("id = ?", enrollmentId).First(&enrollment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return EntityNotFound(fmt.Errorf("enrollment with id %d not found", enrollmentId))
			}
			return fmt.Errorf("error fetching enrollment: %w", err)
		}

		if enrollment.Course.ID == 0 {
			return errors.New("course information not loaded")
		}

		if adminId != enrollment.Course.CreatorID {
			return errors.New("only course creator can approve enrollments")
		}

		if enrollment.Status != models.EnrollmentStatusPending {
			return fmt.Errorf("cannot approve enrollment with status: %s", enrollment.Status)
		}

		enrollment.Status = models.EnrollmentStatusApproved
		if err := tx.Save(&enrollment).Error; err != nil {
			return UpdateEntityFailure(fmt.Errorf("failed to update enrollment status: %w", err))
		}

		return nil
	})
}

// RejectEnrolment rejects a pending enrollment request for a course.
//
// Parameters:
//   - adminId: The ID of the admin (course creator) rejecting the enrollment.
//   - enrollmentId: The ID of the enrollment to be rejected.
//
// Returns:
//   - error: An error if the rejection process fails, nil otherwise.
func (s *EnrollmentService) RejectEnrolment(adminId, enrollmentId uint) error {
	if adminId == 0 || enrollmentId == 0 {
		return errors.New("invalid input: adminId and enrollmentId must be non-zero")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		var enrollment models.Enrollment
		if err := tx.Preload(models.CoursesTable).Where("id = ?", enrollmentId).First(&enrollment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return EntityNotFound(fmt.Errorf("enrollment with id %d not found", enrollmentId))
			}
			return fmt.Errorf("error fetching enrollment: %w", err)
		}

		if enrollment.Course.ID == 0 {
			return errors.New("course information not loaded")
		}

		if adminId != enrollment.Course.CreatorID {
			return errors.New("only course creator can reject enrollments")
		}

		if enrollment.Status != models.EnrollmentStatusPending {
			return fmt.Errorf("cannot reject enrollment with status: %s", enrollment.Status)
		}

		enrollment.Status = models.EnrollmentStatusRejected
		if err := tx.Save(&enrollment).Error; err != nil {
			return UpdateEntityFailure(fmt.Errorf("failed to update enrollment status: %w", err))
		}

		return nil
	})
}
