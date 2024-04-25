# Direct-Message_backend
## Account stuff
### POST
``` bash
// Register
curl -X POST -H "Content-Type: application/json" \
-d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register

// Login
curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com","password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login
```
### PUT
``` bash
// Update avatar
curl -X PUT -F 'avatar=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=99e6f3bb24d151e7e7d2b9cf2c13ad47d1334ffabf9b6c89e0765a46123f8cf3' localhost:8080/update-avatar

// Update background
curl -X PUT -F 'background=@/home/user/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=99e6f3bb24d151e7e7d2b9cf2c13ad47d1334ffabf9b6c89e0765a46123f8cf3' localhost:8080/update-background

// Update email this will mark all current token as deleted
curl -X PUT -F 'email=theanh1@mail.com' -F 'token=28fa0704b25e5203d7202086115af975aed348e605c6c501a44cb6477bced0a5' localhost:8080/update-email

// Update password
curl -X PUT -F 'password=a9e986de49be77b63571db377f60f76213d9a22471a551a37adcd8a88f26f411' -F 'token=ea51524535ed17ff95aebbce71024f7e1a9f8d638affc2c12af054989d7fda37' localhost:8080/update-password

// Update name
curl -X PUT -F 'name=the anh' -F 'token=99e6f3bb24d151e7e7d2b9cf2c13ad47d1334ffabf9b6c89e0765a46123f8cf3' localhost:8080/update-name

// Update private status
curl -X PUT -F 'isPrivate=0' -F 'token=bfd4b19c287fd32707544bd43069b45901aca0c22428f5c669108c5d993f5952' localhost:8080/update-private-status
```
### GET
``` bash
// Get self info
curl -X GET -F 'token=bfd4b19c287fd32707544bd43069b45901aca0c22428f5c669108c5d993f5952' localhost:8080/get-self-info

// Get avatar
curl -X GET -F 'imgName=1.jpg' -F 'token=bfd4b19c287fd32707544bd43069b45901aca0c22428f5c669108c5d993f5952' localhost:8080/get-avatar --output 1.jpg

// Get background
curl -X GET -F 'imgName=1.jpg' -F 'token=bfd4b19c287fd32707544bd43069b45901aca0c22428f5c669108c5d993f5952' localhost:8080/get-background --output 1.jpg
```
### DELETE
``` bash
// Delete self
curl -X DELETE -F 'email=user1@mail.com' -F 'password=12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9' -F 'token=1679acb67da0b7c5233aaa615bc38b39678d3dd3dd36b2e5b63409e98dbe5941' localhost:8080/delete-self
```
## Search friend stuff
### GET
``` bash
// Search by name
// This wont return user at position 0
curl -X GET -F 'name=user' -F 'page=0' -F 'token=04ce9447a7cc8d61cd18d063b40bace2681de4cf7658b0ea70412268757a318f' localhost:8080/get-by-name
// Search by email
curl -X GET -F 'email=user1@mail.com' -F 'token=d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' localhost:8080/get-by-email
```
## Prekey bundle stuff
### PUT
``` bash
// Update prekey bundle
curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933886' localhost:8080/update-prekey-bundle
```
### POST
``` bash
curl -X POST -F 'requestToEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'token=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720'
```
### GET
``` bash 
curl -X GET -F 'email=user1@mail.com' -F 'token=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' localhost:8080/get-prekey-bundle
curl -X GET -F 'token=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' localhost:8080/get-friend-request
```