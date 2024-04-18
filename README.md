# Direct-Message_backend
```
// Register
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password1"}' http://localhost:8090/register

// Login
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1","password":"password1"}' http://localhost:8090/login

// Update avatar and background
curl -X PUT \
-F 'avatar=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' \
-F 'background=@/home/user/MEGA/Pictures/Wallpapers/windowUpdate.jpg' \
-F 'token=9166c5aa5d24c431edceb983397d0c870550cb18aec6b5fa656912a352f47930' localhost:8090/update-avatar

// Update email this will mark all current token as deleted
curl -X PUT \
-F 'email=theanh@mail.com' \
-F 'token=9166c5aa5d24c431edceb983397d0c870550cb18aec6b5fa656912a352f47930' localhost:8090/update-avatar
// Update phone number this will mark all current token as deleted
curl -X PUT \
-F 'phoneNumb=0123456789' \
-F 'token=9166c5aa5d24c431edceb983397d0c870550cb18aec6b5fa656912a352f47930' localhost:8090/update-avatar
```