# Direct-Message_backend
## Account stuff
### POST
``` bash
// Register
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register

// Login
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com","password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login
```
### PUT
``` bash
// Update avatar
curl -X PUT -F 'avatar=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=2dbfe6e67d6cfb6bb1dc7d6332ef60d39731dc93cb6a0e2b20811b2d14208520' localhost:8080/update-avatar

// Update background
curl -X PUT -F 'background=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=2dbfe6e67d6cfb6bb1dc7d6332ef60d39731dc93cb6a0e2b20811b2d14208520' localhost:8080/update-background

// Update email this will mark all current token as deleted
curl -X PUT -F 'email=theanh1@mail.com' -F 'token=f3d16053e5ddcd02b1c21a1543163e603f73c51bfdc5c6ba916e00ea257625fa' localhost:8080/update-email

// Update password
curl -X PUT -F 'password=a9e986de49be77b63571db377f60f76213d9a22471a551a37adcd8a88f26f411' -F 'token=f3d16053e5ddcd02b1c21a1543163e603f73c51bfdc5c6ba916e00ea257625fa' localhost:8080/update-password

// Update name
curl -X PUT -F 'name=the anh' -F 'token=8bceb11f5c2f037b7a2320dd27f1d65625b91f7be057c35388b248bf1bd9b8cc' localhost:8080/update-name

// Update private status
curl -X PUT -F 'isPrivate=0' -F 'token=8bceb11f5c2f037b7a2320dd27f1d65625b91f7be057c35388b248bf1bd9b8cc' localhost:8080/update-private-status
```
### GET
``` bash
// Get self info
curl -X GET -F 'token=8bceb11f5c2f037b7a2320dd27f1d65625b91f7be057c35388b248bf1bd9b8cc' localhost:8080/get-self-info

// Get avatar
curl -X GET -F 'imgName=1.jpg' -F 'token=2dbfe6e67d6cfb6bb1dc7d6332ef60d39731dc93cb6a0e2b20811b2d14208520' localhost:8080/get-avatar --output 1.jpg

// Get background
curl -X GET -F 'imgName=1.jpg' -F 'token=2dbfe6e67d6cfb6bb1dc7d6332ef60d39731dc93cb6a0e2b20811b2d14208520' localhost:8080/get-background --output 1.jpg
```
### DELETE
``` bash
// Delete self
curl -X DELETE -F 'email=user1@mail.com' -F 'password=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' -F 'token=422167be4da19435b73dadece46117040e74f4440e4b6f51d6985563c61ea03e' localhost:8080/delete-self
```
## Search friend stuff
### GET
``` bash
// Search by name
// This wont return user at position 0
curl -X GET -F 'name=user' -F 'page=0' -F 'token=8bceb11f5c2f037b7a2320dd27f1d65625b91f7be057c35388b248bf1bd9b8cc' localhost:8080/get-by-name
// Search by email
curl -X GET -F 'email=user2@mail.com' -F 'token=8bceb11f5c2f037b7a2320dd27f1d65625b91f7be057c35388b248bf1bd9b8cc' localhost:8080/get-by-email
```
## Prekey bundle stuff
### PUT
``` bash
// Update prekey bundle
curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=5c5af6c9131c4d255a8402eb818a6a3625b88f498de87eabdaf2506dd251b483' localhost:8080/update-prekey-bundle
```
### POST
``` bash
curl -X POST -F 'toEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'opkUsed=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4' -F 'token=95b1abaa7f1b1ec063007fed21ed57b9ddb93297dc2936cd5e21c329b886ffda' localhost:8080/add-friend-request
curl -X POST -F 'email=user2@mail.com' -F 'token=95d0678e0e865f55d967d16f5cea6c4d15922594a72a2fb3ee2c0b1f22859647' localhost:8080/accept-friend-request
curl -X POST -F 'email=user2@mail.com' -F 'token=95d0678e0e865f55d967d16f5cea6c4d15922594a72a2fb3ee2c0b1f22859647' localhost:8080/reject-friend-request
```
### GET
``` bash 
curl -X GET -F 'email=user1@mail.com' -F 'token=95b1abaa7f1b1ec063007fed21ed57b9ddb93297dc2936cd5e21c329b886ffda' localhost:8080/get-prekey-bundle
curl -X GET -F 'token=95d0678e0e865f55d967d16f5cea6c4d15922594a72a2fb3ee2c0b1f22859647' localhost:8080/get-friend-request
```
## Friend list stuff
### GET
``` bash
curl -X GET -F 'token=8e0b98788e6115e595486b3361876a144a87c7eb0ae1ac443b33a3082f0dfac0' localhost:8080/get-friend-list
curl -X GET -F 'token=8e0b98788e6115e595486b3361876a144a87c7eb0ae1ac443b33a3082f0dfac0' localhost:8080/get-all-message
curl -X GET -F 'content=1_2024-04-30 17:50:09.txt' -F 'token=8e0b98788e6115e595486b3361876a144a87c7eb0ae1ac443b33a3082f0dfac0' localhost:8080/get-message-content --output '1_2024-04-30 17:50:09.txt'
```
### POST
``` bash
curl -X POST -F 'email=user2@mail.com' -F 'content=@/home/admin/Downloads/text.txt' -F 'token=51cfdddad1ed965974035d54b42a5f2920a467f7bbf5d62c756c6f55dbc0fcb9' localhost:8080/send-message-friend-unencrypt
```