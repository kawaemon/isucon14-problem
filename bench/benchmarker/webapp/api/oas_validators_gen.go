// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *AppChair) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Stats.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "stats",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AppChairStats) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.RecentRides == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "recent_rides",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Float{}).Validate(float64(s.TotalEvaluationAvg)); err != nil {
			return errors.Wrap(err, "float")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "total_evaluation_avg",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AppGetNearbyChairsOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Chairs == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.Chairs {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chairs",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AppGetNotificationOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Chair.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chair",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AppPostRequestEvaluateReq) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:        true,
			Min:           1,
			MaxSet:        true,
			Max:           5,
			MinExclusive:  false,
			MaxExclusive:  false,
			MultipleOfSet: false,
			MultipleOf:    0,
		}).Validate(int64(s.Evaluation)); err != nil {
			return errors.Wrap(err, "int")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "evaluation",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *AppRequest) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if err := func() error {
		if value, ok := s.Chair.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chair",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ChairGetNotificationOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ChairRequest) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if value, ok := s.Status.Get(); ok {
			if err := func() error {
				if err := value.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OwnerGetChairsOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Chairs == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chairs",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *OwnerGetSalesOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Chairs == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chairs",
			Error: err,
		})
	}
	if err := func() error {
		if s.Models == nil {
			return errors.New("nil is invalid value")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "models",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s RequestStatus) Validate() error {
	switch s {
	case "MATCHING":
		return nil
	case "DISPATCHING":
		return nil
	case "DISPATCHED":
		return nil
	case "CARRYING":
		return nil
	case "ARRIVED":
		return nil
	case "COMPLETED":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
