echo "1. Spam create user and login api
2. Spam update image and background api
3. Spam update email
"
echo -n "Select your option: "
read options
case $options in
    "1")
        for i in {0..200..1}
        do
            echo "Create user${i}@mail.com"
            curl_command="curl --no-progress-meter -X POST -H \"Content-Type: application/json\" -d '{\"username\":\"user${i}@mail.com\", \"password\":\"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9\"}' http://localhost:8080/register"
            eval "$curl_command"
            echo "Login user${i}@mail.com"
            curl_command="curl --no-progress-meter -X POST -H \"Content-Type: application/json\" -d '{\"username\":\"user${i}@mail.com\", \"password\":\"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9\"}' http://localhost:8080/login"
            eval "$curl_command"
            echo "--------------------------------------------"
        done
    ;;
    "2")
        echo "Create user1@mail.com"
        curl -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register
        echo "Login user1@mail.com"
        data=$(curl --no-progress-meter -X POST -H "Content-Type: application/json" -d '{"username":"user1@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/login)
        token1=$(echo "$data" | jq -r '.token')
        for i in {0..200..1}
        do
            curl_command="curl -X PUT -F 'avatar=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=${token1}' localhost:8080/update-avatar"
            eval "$curl_command"
            curl_command="curl -X PUT -F 'background=@/home/admin/MEGA/Pictures/Wallpapers/windowChan.jpg' -F 'token=${token1}' localhost:8080/update-background"
            eval "$curl_command"
        done
    ;;
    "3")
        echo "Create user0@mail.com"
        curl -X POST -H "Content-Type: application/json" -d '{"username":"user0@mail.com", "password":"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9"}' http://localhost:8080/register
        for i in {0..200..1}
        do
        user="user${i}@mail.com"
        j=$((i+1))
        new="user${j}@mail.com"
        echo $new
        curl_command="curl --no-progress-meter -X POST -H \"Content-Type: application/json\" -d '{\"username\":\"${user}\", \"password\":\"12a60f274133d470bd1435a8e845d7f501950452440018f110f85480670d20f9\"}' http://localhost:8080/login"
        data=$(eval "$curl_command")
        token=$(echo "$data" | jq -r '.token')
        curl_command="curl -X PUT -F 'email=${new}' -F 'token=${token}' localhost:8080/update-email"
        eval "$curl_command"
        done
    ;;
esac
