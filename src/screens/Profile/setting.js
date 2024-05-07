import { View, Text, TouchableOpacity, StyleSheet } from 'react-native'
import React from 'react'
import Entypo from 'react-native-vector-icons/Entypo'
import FA5 from 'react-native-vector-icons/FontAwesome5'
import Colors from '../../asset/styles/color'
import MI from 'react-native-vector-icons/MaterialIcons'
export default function SettingScreen({navigation}) {
  return (
    <View>
      <View style={{flexDirection:'row', alignItems:'center', backgroundColor:Colors._white,shadowOffset:{width:10, height:10}, elevation:20, shadowColor:Colors._black}}>
        <TouchableOpacity style={styles.btn} onPress={()=>navigation.goBack()}>
            <Entypo name='chevron-left' color={Colors._black} size={24}/>      
        </TouchableOpacity>
        <Text style={styles.txt_header}>Tùy Chỉnh</Text>
      </View>
      <View style={{marginTop:40, paddingHorizontal:15}}>
        <Text style={styles.txt_intro}>Cá nhân hóa</Text>
        <TouchableOpacity style={styles.box_content}>
          <Entypo name="user" size={20} color={Colors._black}/>
          <Text style={styles.txt_content}>Thông tin cá nhân</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.box_content}>
          <FA5 name="paint-brush" size={20} color={Colors._black}/>
          <Text style={styles.txt_content}>Thay đổi ảnh đại điện</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.box_content}>
          <FA5 name="paint-roller" size={20} color={Colors._black}/>
          <Text style={styles.txt_content}>Thay đổi ảnh bìa</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.box_content}>
          <Entypo name="book" size={20} color={Colors._black}/>
          <Text style={styles.txt_content}>Giới thiệu về bản thân</Text>
        </TouchableOpacity>
      </View>
      <View style={{marginTop:40, paddingHorizontal:15}}>
        <Text style={styles.txt_intro}>Cài đặt</Text>
        <TouchableOpacity style={styles.box_content}>
          <MI name="admin-panel-settings" size={22} color={Colors._black}/>
          <Text style={styles.txt_content}>Quản lý tài khoản</Text>
        </TouchableOpacity>
        <TouchableOpacity style={styles.box_content}>
          <MI name="settings-applications" size={22} color={Colors._black}/>
          <Text style={styles.txt_content}>Cài đặt chung</Text>
        </TouchableOpacity>
      </View>
    </View>
  )
}
const styles = StyleSheet.create({
    btn:{
        width:50,
        height:50,
        alignItems:'center',
        justifyContent:'center',
        borderRadius:50,
        
    },
    txt_header:{
        color:Colors._black,
        fontSize:14,
        fontWeight:'bold',
        textAlign:'center',
        width:'75%'
    },
    txt_intro:{
      color:Colors._secondary, fontWeight:'bold', fontSize:14
    },
    txt_content:{
      marginLeft:10,
      color:Colors._black,
      fontWeight:'500'
    },
    box_content:{
      backgroundColor:Colors._white,
      padding:20,
      borderRadius:10,
      width:'100%',
      shadowColor:Colors._black,
      elevation: 5,
      flexDirection:'row',
      alignItems:'center',
      marginVertical:5,
    },

})