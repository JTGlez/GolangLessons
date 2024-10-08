package main

type SalaryError struct {
	Message string
	Code    int
}

func (s *SalaryError) Error() string {
	return s.Message
}
