package utils

import (
	"testing"
	"time"

	"github.com/RuhullahReza/Employee-App/app/domain"

	"github.com/stretchr/testify/assert"
)

func TestValidateAndSanitizeRequest(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "re  za",
			LastName:  "ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		err := ValidateAndSanitizeRequest(&req)
		assert.NoError(t, err)
		assert.Equal(t, "Re Za", req.FirstName)
	})

	t.Run("empty first name", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "",
			LastName:  "ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		err := ValidateAndSanitizeRequest(&req)
		assert.ErrorIs(t, err, ErrEmptyName)
	})

	t.Run("invalid name format", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "ozza 1",
			LastName:  "ozza",
			Email:     "test.test@gmail.com",
			HireDate:  "2024-03-03",
		}

		err := ValidateAndSanitizeRequest(&req)
		assert.ErrorIs(t, err, ErrInvalidName)
	})

	t.Run("invalid email", func(t *testing.T) {
		req := domain.EmployeeRequest{
			FirstName: "reza",
			LastName:  "ozza",
			Email:     "test.testgmail.com",
			HireDate:  "2024-03-03",
		}

		err := ValidateAndSanitizeRequest(&req)
		assert.ErrorIs(t, err, ErrInvalidEmail)
	})
}

func TestParseDateString(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		strDate := "2023-03-03"
		d, err := ParseDateString(strDate)
		assert.NoError(t, err)

		expectedDate := time.Date(2023, 3, 3, 0, 0, 0, 0, time.UTC)
		assert.Equal(t, expectedDate, d)

	})

	t.Run("invalid date", func(t *testing.T) {
		strDate := "2023-0303"
		_, err := ParseDateString(strDate)
		assert.Error(t, err)
	})
}
