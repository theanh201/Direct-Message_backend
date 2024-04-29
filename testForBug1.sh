echo "1. Spam create user and login
2. Spam upate image and background api"
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
esac
