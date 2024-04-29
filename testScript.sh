echo "1. Test add friend user2 add friend user1"
echo -n "Select your option: "
read options
case $options in
    "1")
        for i in {0..1000..1}
        do
        echo "Create user1@mail.com"
        curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register
        echo "Create user2@mail.com"
        curl -X POST -H "Content-Type: application/json" -d '{"username":"user2@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register
        echo "--------------------------------------------"

        echo "Login user1@mail.com"
        data=$(curl --no-progress-meter -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login)
        token1=$(echo "$data" | jq -r '.token')
        echo "Login user2@mail.com"
        data=$(curl --no-progress-meter -X POST -H "Content-Type: application/json" -d '{"username":"user2@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login)
        token2=$(echo "$data" | jq -r '.token')
        echo "--------------------------------------------"

        echo "user1 update keybundle"
        curl_command="curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=${token1}' localhost:8080/update-prekey-bundle"
        eval "$curl_command"
        echo "user2 update keybundle"
        curl_command="curl -X PUT -F 'ik=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933881' -F 'spk=d50ffb8450fc139576ff1efe893f23963e2be19d738080ac260d0bd148933880' -F 'opk=1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d9,2782c5f8c4392877f28e1473ae454ae663a05a3ef5ea962c89707f7a99a429d8,1762c5f8c4392877828e1473ae454ae663a05a3ef5ea962c89707f7a99a429d4,d779737d73332c2db9e7c709019a2626970a0f162b3fa4c0fe57b88fed1d9c82' -F 'token=${token2}' localhost:8080/update-prekey-bundle"
        eval "$curl_command"
        echo "--------------------------------------------"

        echo "user2 get user1 keybundle"
        curl_command="curl --no-progress-meter -X GET -F 'email=user1@mail.com' -F 'token=${token2}' localhost:8080/get-prekey-bundle"
        data=$(eval "$curl_command")
        echo $data
        echo "--------------------------------------------"

        echo "user2 send friend request to user1"
        opk_user1=$(echo "$data" | jq -r '.Opk')
        curl_command="curl --no-progress-meter -X POST -F 'toEmail=user1@mail.com' -F'ek=7fb26648cca726f2cce63eda8e92e220684d0200f08d7076a3a4beec121af720' -F 'opkUsed=${opk_user1}' -F 'token=${token2}' localhost:8080/add-friend-request"
        data=$(eval "$curl_command")
        echo $data
        echo $i
        done
    ;;
esac
