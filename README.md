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
-F 'token=2bd5188bc55a2929eb4996532b5188c2f16b330396f6b97b46c3f065e80fea29' \
localhost:8080/update-avatar

// Update background
curl -X PUT \
-F 'background=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' \
-F 'token=a5e763f83a5b302c3b9b638c8e3d03a28cac2d62a91fb7a810485ab4d51aa189' \
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
-F 'token=a5e763f83a5b302c3b9b638c8e3d03a28cac2d62a91fb7a810485ab4d51aa189' \
localhost:8080/get-self-info

// Get avatar
curl -X GET -F 'imgName=1.jpg' \
-F 'token=1ed94aa91c02a21b773c1146d9b01fb542008f20e646a30a51fba93ff1257655' \
localhost:8080/get-avatar \
--output 1.jpg

// Get background
curl -X GET -F 'imgName=1.jpg' \
-F 'token=1ed94aa91c02a21b773c1146d9b01fb542008f20e646a30a51fba93ff1257655' \
localhost:8080/get-background \
--output 1.jpg

```