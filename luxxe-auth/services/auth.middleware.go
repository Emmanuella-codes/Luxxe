package services

var BaseAuthToken = authTokenService(&AuthTokenStruct{
	authPolicy:          headerBearerToken,
	allowExternalAccess: false,
})
