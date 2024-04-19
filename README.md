# Direct-Message_backend
```
// Register
curl -X POST -H "Content-Type: application/json" \
-d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' \
http://localhost:8080/register

// Login
curl -X POST -H "Content-Type: application/json" \
-d '{"username":"user1@mail.com","password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' \
http://localhost:8080/login

// Update avatar
curl -X PUT \
-F 'avatar=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' \
-F 'token=2429560b4977ad5373e4415ce456125af9839c6daa8de9a83f5a85d737f7fb3e' \
localhost:8080/update-avatar

// Update background
curl -X PUT \
-F 'background=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' \
-F 'token=c26182220bc8e58097782baee9c805765e04d331140749ee483bf6cac7134c5a' \
localhost:8080/update-background

// Update email this will mark all current token as deleted
curl -X PUT \
-F 'email=theanh1@mail.com' \
-F 'token=61ce5c6547de8b32b2c8895836be19b0a06993076234a13c2f3a4ef5d773b1ed' \
localhost:8080/update-email

// Update password
curl -X PUT \
-F 'password=a9e986de49be77b63571db377f60f76213d9a22471a551a37adcd8a88f26f411' \
-F 'token=61ce5c6547de8b32b2c8895836be19b0a06993076234a13c2f3a4ef5d773b1ed' \
localhost:8080/update-password

// Update name
curl -X PUT \
-F 'name=the anh' \
-F 'token=61ce5c6547de8b32b2c8895836be19b0a06993076234a13c2f3a4ef5d773b1ed' \
localhost:8080/update-name

// Get self info
curl -X GET \
-F 'token=9fdd52a08e2b4b15287216317c584dbc41febbffb297f5e0a385ebc6dfde4031' \
localhost:8080/get-self-info

// Get avatar
curl -X GET -F 'imgName=1.jpg' \
-F 'token=2429560b4977ad5373e4415ce456125af9839c6daa8de9a83f5a85d737f7fb3e' \
localhost:8080/get-avatar \
--output 1.jpg

```