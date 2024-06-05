package main

import (
    "fmt"
    "sort"
)

type User struct {
    Username string
    Password string
    Profile  string
    Friends  []*User
}

type Status struct {
    Username string
    Content  string
    Comments []Comment
}

type Comment struct {
    Username string
    Content  string
}

var users []*User
var statuses []Status

func register(username, password, profile string) {
    newUser := &User{Username: username, Password: password, Profile: profile}
    users = append(users, newUser)
}

func login(username, password string) *User {
    for _, user := range users {
        if user.Username == username && user.Password == password {
            return user
        }
    }
    return nil
}

func viewStatuses() {
    for i, status := range statuses {
        fmt.Printf("Status %d by %s: %s\n", i, status.Username, status.Content)
        for j, comment := range status.Comments {
            fmt.Printf("\tComment %d by %s: %s\n", j, comment.Username, comment.Content)
        }
    }
}

func addComment(currentUser *User, profileName, statusContent, commentContent string) {
    for i, status := range statuses {
        if status.Username == profileName && status.Content == statusContent {
            comment := Comment{Username: currentUser.Username, Content: commentContent}
            statuses[i].Comments = append(statuses[i].Comments, comment)
            return
        }
    }
    fmt.Println("Status not found")
}

func addFriend(currentUser, friend *User) {
    if friend != nil {
        currentUser.Friends = append(currentUser.Friends, friend)
    } else {
        fmt.Println("Invalid friend")
    }
}

func removeFriend(currentUser, friend *User) {
    if friend != nil {
        friends := currentUser.Friends
        for i, f := range friends {
            if f == friend {
                currentUser.Friends = append(friends[:i], friends[i+1:]...)
                break
            }
        }
    } else {
        fmt.Println("Invalid friend")
    }
}

func editProfile(currentUser *User, newProfile string) {
    currentUser.Profile = newProfile
}

func sortFriends(currentUser *User, ascending bool) {
    friends := currentUser.Friends
    if ascending {
        sort.Slice(friends, func(i, j int) bool { return friends[i].Username < friends[j].Username })
    } else {
        sort.Slice(friends, func(i, j int) bool { return friends[i].Username > friends[j].Username })
    }
    fmt.Println("Friends List after sorting:")
    for _, friend := range friends {
        fmt.Printf("Username: %s, Profile: %s\n", friend.Username, friend.Profile)
    }
}

func searchUser(username string) *User {
    for _, user := range users {
        if user.Username == username {
            return user
        }
    }
    return nil
}

func searchFriend(currentUser *User, username string) *User {
    for _, friend := range currentUser.Friends {
        if friend.Username == username {
            return friend
        }
    }
    return nil
}

func main() {
    var choice int
    var currentUser *User

    for {
        if currentUser == nil {
            fmt.Println("1. Register")
            fmt.Println("2. Login")
            fmt.Scan(&choice)

            if choice == 1 {
                var username, password, profile string
                fmt.Println("Enter Username:")
                fmt.Scan(&username)
                fmt.Println("Enter Password:")
                fmt.Scan(&password)
                fmt.Println("Enter Profile:")
                fmt.Scan(&profile)
                register(username, password, profile)
            } else if choice == 2 {
                var username, password string
                fmt.Println("Enter Username:")
                fmt.Scan(&username)
                fmt.Println("Enter Password:")
                fmt.Scan(&password)
                currentUser = login(username, password)
                if currentUser == nil {
                    fmt.Println("Login failed!")
                }
            }
        } else {
            fmt.Println("1. View Statuses")
            fmt.Println("2. Add Status")
            fmt.Println("3. Add Comment")
            fmt.Println("4. Add Friend")
            fmt.Println("5. Remove Friend")
            fmt.Println("6. Edit Profile")
            fmt.Println("7. Sort Friends")
            fmt.Println("8. Search User")
            fmt.Println("9. Logout")
            fmt.Scan(&choice)

            switch choice {
            case 1:
                viewStatuses()
            case 2:
                var content string
                fmt.Println("Enter Status Content:")
                fmt.Scan(&content)
                statuses = append(statuses, Status{Username: currentUser.Username, Content: content})
            case 3:
                var profileName, statusContent, commentContent string
                fmt.Println("Enter Profile Name of the Status Owner:")
                fmt.Scan(&profileName)
                fmt.Println("Enter Status Content:")
                fmt.Scan(&statusContent)
                fmt.Println("Enter Comment Content:")
                fmt.Scan(&commentContent)
                addComment(currentUser, profileName, statusContent, commentContent)
            case 4:
                var friendUsername string
                fmt.Println("Enter Friend's Username:")
                fmt.Scan(&friendUsername)
                friend := searchUser(friendUsername)
                addFriend(currentUser, friend)
            case 5:
                var friendUsername string
                fmt.Println("Enter Friend's Username:")
                fmt.Scan(&friendUsername)
                friend := searchUser(friendUsername)
                removeFriend(currentUser, friend)
            case 6:
                var newProfile string
                fmt.Println("Enter New Profile:")
                fmt.Scan(&newProfile)
                editProfile(currentUser, newProfile)
            case 7:
                var order int
                fmt.Println("Enter 1 for Ascending, 2 for Descending:")
                fmt.Scan(&order)
                sortFriends(currentUser, order == 1)
            case 8:
                var username string
                fmt.Println("Enter Username to Search:")
                fmt.Scan(&username)
                user := searchUser(username)
                if user != nil {
                    fmt.Printf("User found: %v\n", user)
                } else {
                    fmt.Println("User not found")
                }
            case 9:
                currentUser = nil
            }
        }
    }
}
