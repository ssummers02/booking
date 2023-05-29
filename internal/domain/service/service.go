package service

type Service struct {
	UsersService     UsersService
	ResortsService   ResortsService
	InventoryService InventoryService
	BookingService   BookingService
	CommentService   CommentService
}

func NewServices(r *Storages) *Service {
	user := NewUsersService(r.User)
	resort := NewResortsService(r.Resort)
	inventory := NewInventoryService(r.Inventory, resort)
	booking := NewBookingService(r.Booking, resort, inventory, user)
	comment := NewCommentService(r.Comment)

	return &Service{
		UsersService:     *user,
		ResortsService:   *resort,
		InventoryService: *inventory,
		BookingService:   *booking,
		CommentService:   *comment,
	}
}

type Storages struct {
	User      UserStorage
	Resort    ResortStorage
	Inventory InventoryStorage
	Booking   BookingStorage
	Comment   CommentStorage
}
