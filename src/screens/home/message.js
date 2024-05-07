import { View, Text, FlatList, TouchableOpacity, Image, StyleSheet, ScrollView, ImageBackground } from 'react-native'
import React from 'react'
import Colors from './../../asset/styles/color';
import Entypo from 'react-native-vector-icons/Entypo'
import AntDesign from 'react-native-vector-icons/AntDesign';
import { SafeAreaView } from 'react-native-safe-area-context';
export default function MessageScreen({navigation}) {
// Dummy data for recent chats
const recentChats = [
    { id: '1', name: 'Trường An',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Xác định tạch môn', lastMessageTime: '10:30 AM' },
    { id: '2', name: 'Thế Anh',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Làm App ngu như con lợn', lastMessageTime: 'Yesterday' },
    { id: '3', name: 'Nguyễn Văn A',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Không có kỹ năng làm việc nhóm', lastMessageTime: 'Yesterday' },
    { id: '4', name: 'Trường An',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Xác định tạch môn', lastMessageTime: '10:30 AM' },
    { id: '5', name: 'Thế Anh',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Làm App ngu như con lợn', lastMessageTime: 'Yesterday' },
    { id: '6', name: 'Nguyễn Văn A',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Không có kỹ năng làm việc nhóm', lastMessageTime: 'Yesterday' },
    { id: '7', name: 'Trường An',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Xác định tạch môn', lastMessageTime: '10:30 AM' },
    { id: '8', name: 'Thế Anh',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Làm App ngu như con lợn', lastMessageTime: 'Yesterday' },
    { id: '9', name: 'Nguyễn Văn A',img:require('../../asset/images/design/user.jpg'), lastMessage: 'Không có kỹ năng làm việc nhóm', lastMessageTime: 'Yesterday' },
  ];
const onlineList = [
  {id:1, name:"Minh", img:require('../../asset/images/design/user.jpg')},
  {id:2, name:"Dũng", img:require('../../asset/images/design/user.jpg')},
  {id:3, name:"Minh", img:require('../../asset/images/design/user.jpg')},
  {id:4, name:"Dũng", img:require('../../asset/images/design/user.jpg')},
  {id:5, name:"Minh", img:require('../../asset/images/design/user.jpg')},
  {id:6, name:"Dũng", img:require('../../asset/images/design/user.jpg')}
]
  return (
    <View style={{padding:15}}>
      <View>
        <Text style={{marginBottom:10, color:Colors._black, fontSize:14}}>Đang hoạt động</Text>
        <FlatList 
          data={onlineList}
          horizontal={true}
          showsHorizontalScrollIndicator={false}
          keyExtractor={(item)=> item.id}
          renderItem={({item})=>(
            <View style={styles.onl_user}>
              <Image style={styles.onl_img} source={item.img}/>
              <View style={styles.onl_icon}>
                <Entypo name='dot-single' color='green' size={40}/>
              </View>
              
              <Text style={{textAlign:'center', color:Colors._black, fontSize:12}}>{item.name}</Text>
            </View>
          )}/>
      </View>
      <View>
        <FlatList
          data={recentChats}
          showsVerticalScrollIndicator={false}
          keyExtractor={(item) => item.id}
          renderItem={({ item }) => (
            <TouchableOpacity style={styles.chatItem}>
              <Image style={{width:50, height:50, borderRadius:20}} source={item.img}/>
              <View style={{marginLeft:20, width:'100%', padding:5}}>
                <Text style={styles.contactName}>{item.name}</Text>
                <Text style={styles.lastMessage}>{item.lastMessage}</Text>
                <Text style={styles.lastMessageTime}>{item.lastMessageTime}</Text>
              </View>
            </TouchableOpacity>
          )}/>
      </View>
    </View>
  )
}
const styles = StyleSheet.create({
    chatItem:{
        flexDirection:'row',
        alignItems:'center',
        backgroundColor:Colors._white,
        paddingHorizontal:10,
        paddingVertical:5,
        marginVertical:5,
        borderRadius:10,
        shadowColor:Colors._black,
        elevation:2
    },
    contactName: {
        fontSize: 14,
        fontWeight: 'bold',
        color:Colors._black, 
      
      },
    lastMessage: {
      fontSize: 14,
      color: Colors._black,
      fontWeight:'500',
      marginTop:5
    },
    lastMessageTime: {
      fontSize: 10,
      color:'#666',
      marginTop:10
    },
    onl_user:{
      padding:5,
      width:70,
      height:70,
      position:'relative',
      alignItems:'center'
      
    },
    onl_img:{
      width:45, 
      height:45,
      borderRadius:100
    },
    onl_icon:{
      position:'absolute',
      top:-10,
      right:-10,
      alignItems:'center',
      justifyContent:'center'
    }
})