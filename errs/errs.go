package errs

import "errors"

var (
	ErrValidationFailed = errors.New("ErrValidationFailed")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrCategoryNotFound = errors.New("category not found")
	ErrOrderNotFound = errors.New("order not found")
	ErrOrderItemNotFound = errors.New("order item not found")
	ErrProductNotFound = errors.New("product not found")
	ErrUsernameUniquenessFailed = errors.New("username uniqueness failed")
	ErrUserNotFound = errors.New("user not found")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
	ErrRecordNotFound = errors.New("record not found")
	ErrPermissionDenied = errors.New("permisipn denied")
	ErrSomethingWentWrong = errors.New("something went wrong")
	ErrAdminNotFound = errors.New("admin not found")
	ErrCartNotFound     = errors.New("cart not found")
	ErrWishlistNotFound = errors.New("wishlist not found")
	ErrWishlistItemNotFound = errors.New("wishlist item not found")
	ErrIDIsNotCorrect = errors.New("ID not found")
	ErrIncorrectInput = errors.New("input not found") 
	ErrUnauthorized = errors.New("Unauthorized") 
)
