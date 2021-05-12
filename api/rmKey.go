package api

func RmKey(token string) error {
	_, err := s.KeyRm(ctx, token)
	return err
}
