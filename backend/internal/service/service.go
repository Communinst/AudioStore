package service



type EntityService[E repository.Entity] struct {
	PostOne(ctx context.Context, data *E) (int64, error)
	GetOneById(ctx context.Context, id uint64) (*E, error)
	GetManyById(ctx context.Context, ids []uint64) ([]*E, error)
	GetAll(ctx context.Context) ([]*E, error)
	DeleteOneById(ctx context.Context, id uint64) error
	DeleteManyById(ctx context.Context, ids []uint64) (uint64, error)
}

type Service struct {

}