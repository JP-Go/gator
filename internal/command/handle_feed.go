func handleListFeeds(s *state, cmd command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("Unexpected arguments. Expected none.")
	}

	feedsWithUserName, err := s.db.GetFeedsWithUserName(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feeds:")
	for _, feedWithUserName := range feedsWithUserName {
		fmt.Printf("- Name: %s\n  URL: %s\n  Added By: %s\n",
			feedWithUserName.Name,
			feedWithUserName.Url,
			feedWithUserName.UserName,
		)
	}

	return nil
}
