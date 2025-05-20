package model

type TokenPair struct {
	AccessToken  string
	RefreshToken string
	AccessExp    int64
	RefreshExp   int64
}
