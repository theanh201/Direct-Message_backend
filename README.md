# Direct-Message_backend
## Account stuff
### Register
``` bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"user2@mail.com", "password":"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"}' http://localhost:8080/register
```
### Login
``` bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"}' http://localhost:8080/login
```
### Update prekey bundle
You need to update prekey bundle, this require ik, spk, and 5 opk
``` bash
curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-prekey-bundle
```
### Update avatar
``` bash
curl -X PUT -F 'avatar=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-avatar
```
### Update background
``` bash
curl -X PUT -F 'background=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-background
```
### Update email
This will mark all current token as deleted
``` bash
curl -X PUT -F 'email=theanh1@mail.com' -F 'token=1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8' localhost:8080/update-email
```
### Update password
This will mark all current token as deleted
``` bash
curl -X PUT -F 'password=a9e986de49be77b63571db377f60f76213d9a22471a551a37adcd8a88f26f411' -F 'token=17e2ab217dadbb376170f5f264a76ca93fa39f821033490dc642b054aec51b25' localhost:8080/update-password
```
### Update name
``` bash
curl -X PUT -F 'name=the anh' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-name
// Update private status
curl -X PUT -F 'isPrivate=1' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/update-private-status
```
### Get self info
``` bash
curl -X GET localhost:8080/get-self-info/a777559c23007158a03c7a55a382d5906b8a758935af6a36c3244883708d5992
```
### Get avatar
``` bash
curl -X GET localhost:8080/get-avatar/a552f4baed1aa0f7d5a181fcd52fc54d3a60444bc4fef56e1b2c5fd8c74349eb/2.jpg --output 1.jpg
```
### Get background
``` bash
// Get background
curl -X GET localhost:8080/get-background/a552f4baed1aa0f7d5a181fcd52fc54d3a60444bc4fef56e1b2c5fd8c74349eb/2.jpg --output 1.jpg
```
### Delete your account
You will need to resend email and password, afterward you token and account will be mark as delete
``` bash
curl -X DELETE localhost:8080/delete-self/12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9/user1@mail.com/6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b
```
## Search friend stuff
### Search by name
This will return users but not return private users
``` bash
curl -X GET localhost:8080/get-by-name/a552f4baed1aa0f7d5a181fcd52fc54d3a60444bc4fef56e1b2c5fd8c74349eb/user/0
```
### Search by email
This will return a user even if that user are private
``` bash
curl -X GET localhost:8080/get-by-email/a552f4baed1aa0f7d5a181fcd52fc54d3a60444bc4fef56e1b2c5fd8c74349eb/user2@mail.com
```
## Friend request stuff
### Get prekey bundle
``` bash
curl -X GET localhost:8080/get-prekey-bundle/1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8/user1@mail.com
```
### Send a friend request
In this request opk can be empty
``` bash
curl -X POST -F 'toEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'opkUsed=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4' -F 'token=00f857c72453676829967742fab2a8420542bce4ab14acaf551cd728bab64f12' localhost:8080/add-friend-request
```
### List all friend request you have
``` bash 
curl -X GET localhost:8080/get-friend-request/1c68980897d4910b24ea1ca2c902d6dbefa7dffb09220833a5c0de0d6f2f28e8
```
### Accept friend request
``` bash
curl -X POST -F 'email=user2@mail.com' -F 'token=f6fdf2ffcf58eac0ec31c97c99efec82f06198c07dde9be5b2c21f66ab5ea81f' localhost:8080/accept-friend-request
```
### Reject friend request
``` bash
curl -X POST -F 'email=user2@mail.com' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/reject-friend-request
```
## Friend stuff
### List your friend
``` bash
curl -X GET -F localhost:8080/get-friend-list/01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a
```
### Get all your message
This shold be run when you have a new device and need all old message
``` bash
curl -X GET localhost:8080/get-all-message/12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9
```
### Send message friend unencrypt
This api is for messaging with websocket, before sending anymessage you need to add youself to the online list of the server by sending a json file with you token like this
localhost:8080/send-message-friend-unencrypt
``` json
{
    "case":0,
    "token":"your_token_is_here"
}
```
after that you can start sending message
``` json
{
    "case":1,
    "token":"your_token_is_here",
    "content":"this_is_my_message",
    "email":"user1@mail.com"
}
```
if other user are online, message are dilivered to the other user
### Get message after time
This is for when you get back from offline
``` bash
curl -X GET localhost:8080/get-all-message-after-time/12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9/2024-05-03 20:57:28
```