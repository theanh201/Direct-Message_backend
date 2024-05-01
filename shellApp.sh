function hash(){
    string=$1
    hash_algorithm=sha256sum
    hashed_string=$(echo "$string" | $hash_algorithm)
    hashed_string=$(head -c 64 <<< "$hashed_string")
    echo -n $hashed_string
}
function updateProfile(){
    token=$1

}
function login(){
    token=$1
    clear
    while true; do
        clear
        echo "--Main Menu--"
        echo "1. Edit profile"
        echo "2. Search other user"
        echo "3. Update prekey bunle (this should happen automatic when login)"
        echo "0. Logout"
        echo -n "Select your option: "
        read options
        case $options in 
            "0")
            clear
            return
            ;;
            "1")
                clear
                while true; do
                    echo "--Current profile--"
                    curl_command="curl -X GET -F 'token=$token' localhost:8080/get-self-info"
                    eval $curl_command
                    echo "--Change options--"
                    echo "1. Update avatar"
                    echo "2. Update background"
                    echo "3. Update email (you will be logout)"
                    echo "4. Update password (you will be logout)"
                    echo "5. Update name"
                    echo "6. Update private status"
                    echo "7. Delete account"
                    echo "0. Return"
                    echo -n "Select your options: "
                    read options
                    case $options in 
                    0)
                        clear
                        break
                    ;;
                    1)
                        clear
                        echo -n "Path to your new avatar: "
                        read path
                        curl_command="curl -X PUT -F 'avatar=@$path' -F 'token=$token' localhost:8080/update-avatar"
                        eval $curl_command
                    ;;
                    2)
                        clear
                        echo -n "Path to your new backgroud: "
                        read path
                        curl_command="curl -X PUT -F 'background=@$path' -F 'token=$token' localhost:8080/update-background"
                        eval $curl_command
                    ;;
                    3)
                        clear
                        echo -n "Your new email: "
                        read email
                        curl_command="curl --no-progress-meter -X PUT -F 'email=$email' -F 'token=$token' localhost:8080/update-email"
                        data=$(eval $curl_command)
                        message=$(echo "$data" | jq -r '.message')
                        if [ -n "$message" ];then
                            return
                        fi                            
                    ;;
                    4)
                        clear
                        echo -n "Your new password: "
                        read password
                        hashed_password=$(hash $password)
                        curl_command="curl --no-progress-meter -X PUT -F 'password=$hashed_password' -F 'token=$token' localhost:8080/update-password"
                        data=$(eval $curl_command)
                        message=$(echo "$data" | jq -r '.message')
                        if [ -n "$message" ];then
                            return
                        fi  
                    ;;
                    5)
                        clear
                        echo -n "Your new name: "
                        read name
                        curl_command="curl -X PUT -F 'name=$name' -F 'token=$token' localhost:8080/update-name"
                        eval $curl_command
                    ;;
                    6)
                        clear
                        echo -n "Private status 0 or 1: "
                        read status
                        curl_command="curl -X PUT -F 'isPrivate=$status' -F 'token=$token' localhost:8080/update-private-status"
                        eval $curl_command
                    ;;
                    7)
                        clear
                        echo "Enter email and password for delete confermation"
                        echo -n "Your email: "
                        read email
                        echo -n "Your password: "
                        read password
                        hashed_password=$(hash $password)
                        curl_command="curl --no-progress-meter -X DELETE -F 'email=$email' -F 'password=$hashed_password' -F 'token=$token' localhost:8080/delete-self" 
                        date=$(eval $curl_command)
                        message=$(echo "$data" | jq -r '.message')
                        if [ -n "$message" ];then
                            return
                        fi        
                    ;;
                    esac
                done
            ;;
            2)
                clear
                while true; do
                    echo "--Search option--"
                    echo "1. Search by name"
                    echo "2. Search by email"
                    echo "0. Return"
                    read options
                    case $options in
                    0)
                        clear
                        break
                    ;;
                    1)
                        clear
                        echo -n "Name want to search: "
                        read name
                        echo -n "Page: "
                        read page
                        curl_command="curl -X GET -F 'name=$name' -F 'page=$page' -F 'token=$token' localhost:8080/get-by-name"
                        eval $curl_command
                    ;;
                    2)
                        clear
                        echo -n "Email want to search: "
                        read email
                        curl_command="curl -X GET -F 'email=$email' -F 'token=$token' localhost:8080/get-by-email"
                        eval $curl_command
                    ;;
                    esac
                done
            ;;
            3)
                echo "not yet implement"  
            ;;
        esac
    done
}
while true; do
    # clear
    echo "--Demo App shell--
1. Login
2. Register
0. Exit"
    echo -n "Select your option: "
    read options
    case $options in
    "1")
        clear
        echo "--Login--"
        echo -n "Account: "
        read account
        echo -n "Password: "
        read password
        hashed_password=$(hash $password)
        # echo -n $hashed_password
        curl_command="curl --no-progress-meter -X POST -H \"Content-Type: application/json\" -d '{\"username\":\"$account\", \"password\":\"$hashed_password\"}' http://localhost:8080/login"
        data=$(eval $curl_command)
        # echo $data
        token=$(echo "$data" | jq -r '.token')
        if [ -n "$token" ];then
            login $token
        fi
    ;;
    "2")
        clear
        echo "--Register--"
        echo -n "Account: "
        read account
        echo -n "Password: "
        read password
        hashed_password=$(hash $password)
        # echo -n $hashed_password
        curl_command="curl --no-progress-meter -X POST -H \"Content-Type: application/json\" -d '{\"username\":\"$account\", \"password\":\"$hashed_password\"}' http://localhost:8080/register"
        eval $curl_command
    ;;
    "0")
        clear
        break
    ;;
    esac
done