package controllers

import (
	"fmt"

	"github.com/brunoeduardodev/simple-ecommerce-go/internal/repositories"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/utils"
	"github.com/brunoeduardodev/simple-ecommerce-go/internal/views"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userRepository repositories.UserRepository
}

// @BasePath /
// Sign Up
// @Summary Does a sign up
// @Param request body repositories.CreateUserInput true  "Sign up input"
// @Description Does a sign up
// @Produce json
// @Router /sign-up [post]
func (c AuthController) SignUp(ctx *gin.Context) {
	var input repositories.CreateUserInput

	err := ctx.ShouldBind(&input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	existingUser, err := c.userRepository.FindByEmail(input.Email)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		ctx.JSON(503, gin.H{"error": "Couldn't sign up #101"})
		return
	}

	if existingUser != nil {
		ctx.JSON(400, gin.H{"error": "User with same email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		ctx.JSON(503, gin.H{"error": "Couldn't sign up #102"})
	}

	input.Password = string(hashedPassword)

	err = c.userRepository.Create(input)

	if err != nil {
		fmt.Printf("Error: %v", err)
		ctx.JSON(503, gin.H{"error": "Unable to sign up #103"})
		return
	}

	user, err := c.userRepository.FindByEmail(input.Email)

	if err != nil {
		fmt.Printf("Error: %v", err)
		ctx.JSON(503, gin.H{"error": "Something went wrong #104"})
		return
	}

	token, err := utils.GenerateToken(*user)

	if err != nil {
		fmt.Printf("Error: %v", err)
		ctx.JSON(503, gin.H{"error": "Something went wrong #105"})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "user": user})
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @BasePath /
// Sign In
// @Summary Does a sign in
// @Description Does a sign in
// @Param request body SignInInput true "Sign In input"
// @Produce json
// @Router /sign-in [post]
func (c AuthController) SignIn(ctx *gin.Context) {
	var input SignInInput

	err := ctx.ShouldBind(&input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid body"})
		return
	}

	existingUser, err := c.userRepository.FindByEmail(input.Email)

	if err != nil {
		ctx.JSON(503, gin.H{"error": "Couldn't sign in"})
		return
	}

	if existingUser == nil {
		ctx.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(input.Password))

	fmt.Printf("Comparing %v and %v\n", existingUser.Password, input.Password)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(*existingUser)

	if err != nil {
		ctx.JSON(400, gin.H{"error": "Couldn't sign in"})
		return

	}

	ctx.JSON(200, gin.H{"token": token, "user": views.DbUserToView(existingUser)})
}

func NewAuthController(repository repositories.UserRepository) AuthController {
	return AuthController{userRepository: repository}
}
