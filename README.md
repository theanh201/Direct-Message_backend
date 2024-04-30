# Direct-Message_backend
## Account stuff
### POST
``` bash
// Register
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register

// Login
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login
```
### PUT
``` bash
// Update avatar
curl -X PUT -F 'avatar=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-avatar

// Update background
curl -X PUT -F 'background=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-background

// Update email this will mark all current token as deleted
curl -X PUT -F 'email=theanh1@mail.com' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-email

// Update password this will mark all current token as deleted
curl -X PUT -F 'password=a9e986de49be77b63571db377f60f76213d9a22471a551a37adcd8a88f26f411' -F 'token=17e2ab217dadbb376170f5f264a76ca93fa39f821033490dc642b054aec51b25' localhost:8080/update-password

// Update name
curl -X PUT -F 'name=the anh' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-name

// Update private status
curl -X PUT -F 'isPrivate=1' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-private-status
```
### GET
``` bash
// Get self info
curl -X GET -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-self-info

// Get avatar
curl -X GET -F 'imgName=1.jpg' -F 'token=c8ff7675fcac9a4bec35759751f4315a3a79b8126f906ef012cb5bbdff03acaa' localhost:8080/get-avatar --output 1.jpg

// Get background
curl -X GET -F 'imgName=1.jpg' -F 'token=aa5d73f34c3f0f2ad080101bba90f13bdd8cdb1f16ada718ff7c743a3ffb540f' localhost:8080/get-background --output 1.jpg
```
### DELETE
``` bash
// Delete self
curl -X DELETE -F 'email=user1@mail.com' -F 'password=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' -F 'token=c8ff7675fcac9a4bec35759751f4315a3a79b8126f906ef012cb5bbdff03acaa' localhost:8080/delete-self
```
## Search friend stuff
### GET
``` bash
// Search by name
curl -X GET -F 'name=user' -F 'page=1' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-by-name
// Search by email
curl -X GET -F 'email=user2@mail.com' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-by-email
```
## Prekey bundle stuff
### PUT
``` bash
// Update prekey bundle
curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-prekey-bundle
```
### POST
``` bash
curl -X POST -F 'toEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'opkUsed=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4' -F 'token=2e4283701cfa56a6f4cbd65ff1c3a9432aa6db7e97955c2b8adabf955b43cf2d' localhost:8080/add-friend-request
curl -X POST -F 'email=user2@mail.com' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/accept-friend-request
curl -X POST -F 'email=user2@mail.com' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/reject-friend-request
```
### GET
``` bash 
curl -X GET -F 'email=user1@mail.com' -F 'token=69f5370d221aab26c9c3bfe4f8afbbf2097b224bdd6a86c8bda79baa10a3182e' localhost:8080/get-prekey-bundle
curl -X GET -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-friend-request
```
## Friend list stuff
### GET
``` bash
curl -X GET -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-friend-list
curl -X GET -F 'token=75d5a92d47a0e97dfc96b78d8708374ce20dd117b8a55c257a190b287b21f738' localhost:8080/get-all-message
curl -X GET -F 'content=1_2024-04-30 22:23:09.txt' -F 'token=75d5a92d47a0e97dfc96b78d8708374ce20dd117b8a55c257a190b287b21f738' localhost:8080/get-message-content --output '1_2024-04-30 17:50:09.txt'
```
### POST
``` bash
curl -X POST -F 'email=user2@mail.com' -F 'content=@/home/admin/Downloads/text.txt' -F 'token=f6a70ddf1a543c89ed0dcc63fa8a09812dd4cac2e6a68aa547add92d58f2ea20' localhost:8080/send-message-friend-unencrypt
```