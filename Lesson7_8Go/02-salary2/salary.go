package main

type SalaryError struct {
	Message string
	Code    int
}

type ErrorWrapper struct {
	TargetError error
	OtherError  error
}

func (s *SalaryError) Error() string {
	return s.Message
}
