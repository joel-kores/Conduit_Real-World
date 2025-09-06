package converter

import (
	"github.com/joel-kores/Conduit_Real-World/internal/domain/user"
	"github.com/joel-kores/Conduit_Real-World/internal/generated"
)

func ToDBUser(apiUser *generated.User) *user.User {
    return &user.User{
        Email:    apiUser.Email,
        Username: apiUser.Username,
        Bio:      apiUser.Bio,
        Image:    apiUser.Image,
    }
}

func ToAPIUser(dbUser *user.User) *generated.User {
    return &generated.User{
        Email:     dbUser.Email,
        Username:  dbUser.Username,
        Bio:       dbUser.Bio,
        Image:     dbUser.Image,
    }
}
