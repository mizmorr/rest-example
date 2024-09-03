package main

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	// ctx := context.Background()
	// db, err := db.Dial(ctx)

	// if err != nil {
	// 	return err
	// }
	// id := uuid.MustParse("3410daa9-5443-4285-b960-0964ca8b973b")
	// repo_user := repo.NewUserRepo(db)
	// new_user, err := repo_user.Get_User(ctx, id)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(new_user)
	return nil
}
