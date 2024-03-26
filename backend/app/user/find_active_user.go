package user

// func FindActiveUser(users []UserEntity) (*UserEntity, bool, error) {
// 	if len(users) > 0 {
// 		for _, user := range users {
// 			if user.CreatedDate != "" && user.DeletedDate == "" {
// 				found, err := findSuspendedUser(user)
// 				if err != nil {
// 					return nil, false, err
// 				}
// 				if !found {
// 					lastActivity, err := findLastActivity(user)
// 					if err != nil {
// 						return nil, false, err
// 					}

// 					if time.Now().AddDate(-1, 0, 0).Before(lastActivity) {
// 						return &user, true, nil
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return nil, false, nil
// }

// func FindActiveUser(users []UserEntity) UserEntity {
// 	emptyUser := UserEntity{}
// 	oneYearAgo := time.Now().AddDate(-1, 0, 0)

// LoopAllUser:
// 	for _, user := range users {
// 		if isWrongStampCreatedDate(user) || isDeletedUser(user) {
// 			continue LoopAllUser
// 		}

// 		isSuspened, err := isSuspendedUser(user)
// 		if err != nil {
// 			continue LoopAllUser
// 		}
// 		if isSuspened {
// 			continue LoopAllUser
// 		}

// 		if !isInteractAfter(user, oneYearAgo) {
// 			continue LoopAllUser
// 		}

// 		return user
// 	}

// 	return emptyUser
// }
