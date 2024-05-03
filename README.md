# Direct-Message_backend
## Account stuff
### Register
``` bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"user2@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register
```
### Login
``` bash
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login
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
curl -X GET -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-self-info
```
### Get avatar
``` bash
curl -X GET -F 'imgName=1.jpg' -F 'token=c8ff7675fcac9a4bec35759751f4315a3a79b8126f906ef012cb5bbdff03acaa' localhost:8080/get-avatar --output 1.jpg
```
### Get background
``` bash
// Get background
curl -X GET -F 'imgName=1.jpg' -F 'token=aa5d73f34c3f0f2ad080101bba90f13bdd8cdb1f16ada718ff7c743a3ffb540f' localhost:8080/get-background --output 1.jpg
```
### Delete your account
You will need to resend email and password, afterward you token and account will be mark as delete
``` bash
curl -X DELETE -F 'email=user1@mail.com' -F 'password=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' -F 'token=c8ff7675fcac9a4bec35759751f4315a3a79b8126f906ef012cb5bbdff03acaa' localhost:8080/delete-self
```
## Search friend stuff
### Search by name
This will return users but not return private users
``` bash
curl -X GET -F 'name=user' -F 'page=0' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-by-name
```
### Search by email
This will return a user even if that user are private
``` bash
curl -X GET -F 'email=user2@mail.com' -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-by-email
```
## Friend request stuff
### Get prekey bundle
``` bash
curl -X GET -F 'email=user1@mail.com' -F 'token=e89b6a8f11494ffc94399f6bdacaa30eb5a783327c88c8cbc878c792bd4dca29' localhost:8080/get-prekey-bundle
```
### Send a friend request
In this request opk can be empty
``` bash
curl -X POST -F 'toEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'opkUsed=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4' -F 'token=00f857c72453676829967742fab2a8420542bce4ab14acaf551cd728bab64f12' localhost:8080/add-friend-request
```
### List all friend request you have
``` bash 
curl -X GET -F 'token=1777593ba77e512e72a750a90f7ab85a50d729d7d2fdb30984be02dd361e111d' localhost:8080/get-friend-request
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
curl -X GET -F 'token=01f36eb7afe7a112e019fb7f494ca5219aefb1668115d5e1a1494eb85d6ae36a' localhost:8080/get-friend-list
```
### Get all your message
This shold be run when you have a new device and need all old message
``` bash
curl -X GET -F 'token=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' localhost:8080/get-all-message
```
### Send message friend unencrypt
This api is for messaging with websocket, before sending anymessage you need to add youself to the online list of the server by sending a json file with you token like this
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
curl -X GET -F 'time=2024-05-03 20:57:28' -F 'token=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' localhost:8080/get-all-message-after-time
```